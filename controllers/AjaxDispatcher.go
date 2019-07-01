/**
 * Handle ajax requests and send them to the proper handler.
 */
package controllers

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/config"
	"github.com/MangoDowner/mediawiki/includes/php"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"reflect"
)

// Use superglobals, but since it's deprecated, it's not worth fixing
// phpcs:disable MediaWiki.Usage.SuperGlobalsUsage.SuperGlobals

/**
 * @defgroup Ajax Ajax
 */

/**
 * Object-Oriented Ajax functions.
 * @ingroup Ajax
 */
type AjaxDispatcher struct {
	/**
	 * The way the request was made, either a 'get' or a 'post'
	 * @var string $mode
	 */
	mode string

	/**
	 * Name of the requested handler
	 * @var string $func_name
	 */
	funcName string

	/** Arguments passed
	 * @var array $args
	 */
	args []string

	/**
	 * @var Config
	 */
	config config.IConfig

	controller *beego.Controller
}


/**
 * Load up our object with user supplied data
 * @param Config $config
 */
func NewAjaxDispatcher(c *beego.Controller, config config.IConfig) *AjaxDispatcher {
	this := new(AjaxDispatcher)
	this.controller = c
	this.config = config
	this.mode = c.Ctx.Input.Method()
	switch this.mode {
	case "GET", "POST":
		this.funcName = c.GetString("rs")
		if len(c.GetStrings("rsargs")) != 0 {
			this.args = c.GetStrings("rsargs")
		}
	default:
		// Or we could throw an exception:
		// throw new MWException( __METHOD__ . ' called without any data (mode empty).' );
	}
	return this
}

/**
 * Pass the request to our internal function.
 * BEWARE! Data are passed as they have been supplied by the user,
 * they should be carefully handled in the function processing the
 * request.
 *
 * phan-taint-check triggers as it is not smart enough to understand
 * the early return if func_name not in AjaxExportList.
 * @suppress SecurityCheck-XSS
 * @param User $user
 */
func (a *AjaxDispatcher) performAction(user interface{}) {
	if a.mode == "" {
		return
	}
	list := config.Configs.GetList("AjaxExportList")
	if !php.InArray(a.funcName, list)  {
		logs.Debug(" Bad Request for unknown function %s", a.funcName)
		err := includes.WfHttpError(
			a.controller.Ctx,
			400,
			"Bad Request",
			"unknown function " + a.funcName,
		)
		if err != nil {
			logs.Error("fail to return Bad Request page")
		}
		return
	}
	// TODO: 权限判断
	if false {
	//if true {
		err := includes.WfHttpError(
			a.controller.Ctx,
			403,
			"Forbidden",
			"You are not allowed to view pages.",
		)
		if err != nil {
			logs.Error("fail to return Forbidden page")
		}
		return
	}
	logs.Debug("__METHOD__ dispatching %s", a.funcName)
	result := a.callUserFuncArray(a.funcName, a.args)

	if result[0].String() == "" {
		logs.Debug("__METHOD__ ERROR while dispatching %s (%s): no data returned", a.funcName, php.VarExport(a.args))
		err := includes.WfHttpError(
			a.controller.Ctx,
			500,
			"Internal Error",
			fmt.Sprintf("%s returned no data", a.funcName),
		)
		a.performActionErrorHandler(err)
		return
	}
	if result[0].Kind() != reflect.String {
		return
	}
	r := includes.NewAjaxResponse(a.controller.Ctx, result[0].String())
	// Make sure DB commit succeeds before sending a response
	r.SendHeaders()
	r.PrintText()
	logs.Debug("__METHOD__ dispatch complete for  %s", a.funcName)

}

func (a *AjaxDispatcher) performActionErrorHandler(err error) {
	logs.Debug("__METHOD__ ERROR while dispatching %s (%s): %s", a.funcName, php.VarExport(a.args), err)
	if !php.HeadersSent() {
		err := includes.WfHttpError(
			a.controller.Ctx,
			500,
			"Internal Error",
			err.Error(),
		)
		if err != nil {
			logs.Error("fail to return Internal Error page")
		}
		return
	}
	logs.Error(err)
}

/**
 * Call a user function given with an array of parameters
 * @link http://php.net/manual/en/function.call-user-func-array.php
 * @param callback $function <p>
 * The function to be called.
 * </p>
 * @param array $param_arr <p>
 * The parameters to be passed to the function, as an indexed array.
 * </p>
 * @return mixed the function result, or false on error.
 * @since 4.0.4
 * @since 5.0
 */
func (a *AjaxDispatcher) callUserFuncArray(funcName string, paramArr []string) []reflect.Value {
	do := reflect.ValueOf(a).MethodByName(funcName)
	paramLen := len(paramArr)
	args := make([]reflect.Value, paramLen)
	for k, v := range paramArr {
		args[k] = reflect.ValueOf(v)
	}
	return do.Call(args)
}

func (a *AjaxDispatcher) DescIt(name, what string) (string, error) {
	return fmt.Sprintf("%s is %s", name, what), nil
}





















