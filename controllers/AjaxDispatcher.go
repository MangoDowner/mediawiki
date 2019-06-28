/**
 * Handle ajax requests and send them to the proper handler.
 */
package controllers

import (
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/config"
	"github.com/MangoDowner/mediawiki/includes/php"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
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
	args string

	/**
	 * @var Config
	 */
	config config.IConfig

	context *context.Context
}


/**
 * Load up our object with user supplied data
 * @param Config $config
 */
func NewAjaxDispatcher(ctx *context.Context, config config.IConfig) *AjaxDispatcher {
	this := new(AjaxDispatcher)
	this.context = ctx
	this.config = config
	this.mode = ctx.Input.Method()
	switch this.mode {
	case "GET", "POST":
		this.funcName = ctx.Input.Query("rs")
		if ctx.Input.Query("rsargs") != "" {
			this.args = ctx.Input.Query("rsargs")
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
			a.context,
			400,
			"Bad Request",
			"unknown function " + a.funcName,
		)
		if err != nil {
			logs.Error("fail to return Bad Request page")
		}
		return
	}
}











































