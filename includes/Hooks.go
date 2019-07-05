package includes

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/globals"
	"github.com/MangoDowner/mediawiki/includes/exception"
	"github.com/astaxie/beego/logs"
	"reflect"
)

/**
 * A tool for running hook functions.
 */

// Hook function to run
type HookFunc interface{}

/**
* Hooks class.
*
* Used to supersede $wgHooks, because globals are EVIL.
*
* @since 1.18
*/
type Hooks struct {
	/**
	 * Array of events mapped to an array of callbacks to be run
	 * when that event is triggered.
	 */
	 handlers map[string][]HookFunc
}


func NewHooks() *Hooks {
	this := new(Hooks)
	this.handlers = make(map[string][]HookFunc)
	return this
}


/**
 * Attach an event handler to a given hook.
 *
 * @param string $name Name of hook
 * @param callable $callback Callback function to attach
 *
 * @since 1.18
 */
 func (h *Hooks) register(name string, callback interface{}) {
 	h.handlers[name] = append(h.handlers[name], callback)
 }

/**
* Clears hooks registered via Hooks::register(). Does not touch $wgHooks.
* This is intended for use while testing and will fail if MW_PHPUNIT_TEST is not defined.
*
* @param string $name The name of the hook to clear.
*
* @since 1.21
* @throws MWException If not in testing mode.
* @codeCoverageIgnore
*/
func (h *Hooks) clear(name string) {
	if globals.GLOBALS["HW_PHPUNIT_TEST"] == "" && globals.GLOBALS["W_PARSER_TEST"] == "" {
		panic(exception.NewMWException("Cannot reset hooks in operation."))
	}
	h.handlers[name] = nil
}

/**
 * Returns true if a hook has a function registered to it.
 * The function may have been registered either via Hooks::register or in $wgHooks.
 *
 * @since 1.18
 *
 * @param string $name Name of hook
 * @return bool True if the hook has a function registered to it
 */
func (h *Hooks) IsRegistered(name string) bool {
	_, ok := WgHooks[name]
	_, ok1 := h.handlers[name]
	return ok || ok1
}

/**
 * Returns an array of all the event functions attached to a hook
 * This combines functions registered via Hooks::register and with $wgHooks.
 *
 * @since 1.18
 *
 * @param string $name Name of the hook
 * @return array
 */
func (h *Hooks) GetHandlers(name string) (result []HookFunc) {
	if !h.IsRegistered(name) {
		return result
	}
	if _, ok := h.handlers[name]; !ok {
		return WgHooks[name]
	}
	if _, ok := WgHooks[name]; !ok {
		return h.handlers[name]
	}
	for _, v := range h.handlers[name] {
		result = append(result, v)
	}
	for _, v := range WgHooks[name] {
		result = append(result, v)
	}
	return result
}

/**
 * @param string $event Event name
 * @param array|callable $hook
 * @param array $args Array of parameters passed to hook functions
 * @param string|null $deprecatedVersion [optional]
 * @param string &$fname [optional] Readable name of hook [returned]
 * @return null|string|bool
 */
func (h *Hooks) callHook(event string, hook interface{}, args []interface{},
	deprecatedVersion string, funcName *string) (ret interface{}) {
	if funcName == nil {
		funcName = new(string)
	}

	//var hookArr []HookFunc
	//// Turn non-array values into an array. (Can't use casting because of objects.)
	//hookArr = append(hookArr, hook)

	/**
	 * $hook can be: a function, an object, an array of $function and
	 * $data, an array of just a function, an array of object and
	 * method, or an array of object, method, and data.
	 */
	if reflect.ValueOf(hook).Kind() == reflect.Func { // 如果hook是个函数
		*funcName = fmt.Sprintf("hook-%s-closure", event)
		var params []reflect.Value
		for _, v := range args {
			params = append(params, reflect.ValueOf(v))
		}
		result := reflect.ValueOf(hook).Call(params)
		// 考虑到有0返回的情况
		if len(result) > 0 {
			ret = result[0].Interface()
		}
	} else if reflect.ValueOf(hook).Kind() == reflect.String { // 如果hook是个字符串
		panic(fmt.Sprintf("String type hook unsupported now: %s", hook))
	} else {
		panic(fmt.Sprintf("Unknown datatype in hooks for %s.", event))
	}

	// mark hook as deprecated, if deprecation version is specified
	if deprecatedVersion != "" {
		WfDeprecated(fmt.Sprintf("%s hook (used in %s)", event, *funcName),
			deprecatedVersion, "", 0)
	}
	return ret
}


/**
 * Call hook functions defined in Hooks::register and $wgHooks.
 *
 * For the given hook event, fetch the array of hook events and
 * process them. Determine the proper callback for each hook and
 * then call the actual hook using the appropriate arguments.
 * Finally, process the return value and return/throw accordingly.
 *
 * For hook event that are not abortable through a handler's return value,
 * use runWithoutAbort() instead.
 *
 * @param string $event Event name
 * @param array $args Array of parameters passed to hook functions
 * @param string|null $deprecatedVersion [optional] Mark hook as deprecated with version number
 * @return bool True if no handler aborted the hook
 *
 * @throws Exception
 * @throws FatalError
 * @throws MWException
 * @since 1.22 A hook function is not required to return a value for
 *   processing to continue. Not returning a value (or explicitly
 *   returning null) is equivalent to returning true.
 */
func (h *Hooks) Run(event string, args []interface{}, deprecatedVersion string) bool {
	for _, hook := range h.GetHandlers(event) {
		retVal := h.callHook(event, hook, args, deprecatedVersion, nil)
		if retVal == nil {
			continue
		}
		// Process the return value.
		switch v := retVal.(type) {
		case error:
			if v != nil {
				// error was returned. Stop processing, but no error.
				return false
			}
		case bool:
			if !v {
				// false was returned. Stop processing, but no error.
				return false
			}
		case string:
			// String returned means error.
			panic(exception.NewFatalError(v))
			return false
		default:
			continue
		}
	}
	return true
}

/**
 * Call hook functions defined in Hooks::register and $wgHooks.
 *
 * @param string $event Event name
 * @param array $args Array of parameters passed to hook functions
 * @param string|null $deprecatedVersion [optional] Mark hook as deprecated with version number
 * @return bool Always true
 * @throws MWException If a callback is invalid, unknown
 * @throws UnexpectedValueException If a callback returns an abort value.
 * @since 1.30
 */
func (h *Hooks) RunWithoutAbort(event string, args []interface{}, deprecatedVersion string) bool {
	var funcName string
	for _, hook := range h.GetHandlers(event) {
		funcName = ""
		retVal := h.callHook(event, hook, args, deprecatedVersion, &funcName)
		if errA, ok := retVal.(error); ok && errA != nil {
			// False was returned. Stop processing, but no error.
			logs.Debug("Invalid return from %s for unabortable %s.", funcName, event)
		}
	}

	return true
}