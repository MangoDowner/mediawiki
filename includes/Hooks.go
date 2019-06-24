package includes

import (
	"github.com/MangoDowner/mediawiki/globals"
	"github.com/MangoDowner/mediawiki/includes/exception"
	"reflect"
)

/**
 * A tool for running hook functions.
 */

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
	 handlers map[string][]func()
}

var wgHooks map[string][]func()

func NewHooks() *Hooks {
	this := new(Hooks)
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
 func (h *Hooks) register(name string, callback []func()) {
 	h.handlers[name] = callback
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
	_, ok := wgHooks[name]
	_, ok1 := h.handlers[name]
	return !ok || !ok1
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
func (h *Hooks) GetHandlers(name string) (result []func()) {
	if h.IsRegistered(name) {
		return result
	}
	if _, ok := h.handlers[name]; !ok {
		return wgHooks[name]
	}
	if _, ok := wgHooks[name]; !ok {
		return h.handlers[name]
	}
	for _, v := range h.handlers[name] {
		result = append(result, v)
	}
	for _, v := range wgHooks[name] {
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
func (h *Hooks) callHook(event string, hook interface{}, args []MediaWikiServices,
	deprecatedVersion string, fname *string) interface{} {
	var hookFunc []interface{}
	// Turn non-array values into an array. (Can't use casting because of objects.)
	if reflect.ValueOf(hook).Kind() != reflect.Array {
		hookFunc = append(hookFunc, hook.(interface{}))
	} else {
		hookFunc = hook.([]interface{})
	}
	// Either array is empty or it's an array filled with null/false/empty.
	if len(hookFunc) == 0 {
		return nil
	}
	empty := true
	for _, v := range hookFunc {
		if v != nil {
			empty = false
			break
		}
	}
	if empty {
		return nil
	}

	if reflect.ValueOf(hookFunc[0]).Kind() == reflect.Array {
		// First element is an array, meaning the developer intended
		// the first element to be a callback. Merge it in so that
		// processing can be uniform.
		first := hookFunc[0].([]func())
		hookFunc = hookFunc[1:]
		for _, v := range first {
			hookFunc = append(hookFunc, v)
		}
	}
	return func() {

	}
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
func (h *Hooks) Run(event string, args []MediaWikiServices, deprecatedVersion string) bool {
	for _, hook := range h.GetHandlers(event) {
		retval := h.callHook(event, hook, args, deprecatedVersion, nil)
		if retval == nil {
			continue
		}
		// Process the return value.
		if s, ok := retval.(string); ok {
			// String returned means error.
			panic(exception.NewFatalError(s))
		} else if b, ok := retval.(bool); ok && b == false {
			// False was returned. Stop processing, but no error.
			return false
		}
	}
	return true
}