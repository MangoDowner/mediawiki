package includes

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes/libs"
	"github.com/MangoDowner/mediawiki/includes/php"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"go-common/library/log"
)

/**
 * We want some things to be included as literal characters in our title URLs
 * for prettiness, which urlencode encodes by default.  According to RFC 1738,
 * all of the following should be safe:
 *
 * ;:@&=$-_.+!*"(),
 *
 * RFC 1738 says ~ is unsafe, however RFC 3986 considers it an unreserved
 * character which should not be encoded. More importantly, google chrome
 * always converts %7E back to ~, and converting it in this function can
 * cause a redirect loop (T105265).
 *
 * But + is not safe because it"s used to indicate a space; &= are only safe in
 * paths and not in queries (and we don"t distinguish here); " seems kind of
 * scary; and urlencode() doesn"t touch -_. to begin with.  Plus, although /
 * is reserved, we don"t care.  So the list we unescape is:
 *
 * ;:@$!*(),/~
 *
 * However, IIS7 redirects fail when the url contains a colon (see T24709),
 * so no fancy : for IIS7.
 *
 * %2F in the page titles seems to fatally break for some reason.
 *
 * @param string $s
 * @return string
 */
func WfUrlencode(s string) string {
	var needle []string
	if s == "" {
		needle = []string{}
		return ""
	}

	if len(needle) == 0 {
		needle = []string{"%3B", "%40", "%24", "%21", "%2A",
						  "%28", "%29", "%2C", "%2F", "%7E", }
		// TODO: 判断运行环境是否是Microsoft
	}

	s, _ = php.Urldecode(s)
	s = php.StrIReplace(
		needle,
		[]string{";", "@", "$", "!", "*", "(", ")", ",", "/", "~", ":"},
		s,
	)

	return s
}


/**
 * This is the function for getting translated interface messages.
 *
 * @see Message class for documentation how to use them.
 * @see https://www.mediawiki.org/wiki/Manual:Messages_API
 *
 * This function replaces all old wfMsg* functions.
 *
 * @param string|string[]|MessageSpecifier $key Message key, or array of keys, or a MessageSpecifier
 * @param string|string[] ...$params Normal message parameters
 * @return Message
 *
 * @since 1.17
 *
 * @see Message::__construct
 */
func WfMessage(key string, params ...string) *Message {
	message := NewMessage(key, nil, nil)

	// We call Message::params() to reduce code duplication
	if len(params) != 0 {
		message.Params(params)
	}

	return message
}

/**
 * Throws a warning that $function is deprecated
 *
 * @param string $function
 * @param string|bool $version Version of MediaWiki that the function
 *    was deprecated in (Added in 1.19).
 * @param string|bool $component Added in 1.19.
 * @param int $callerOffset How far up the call stack is the original
 *    caller. 2 = function that called the function that called
 *    wfDeprecated (Added in 1.20)
 *
 * @return null
 */
func WfDeprecated(function, version, component string, callerOffset int) {
	if callerOffset == 0 {
		callerOffset = 2
	}
	logs.Debug("Use of %s was deprecated in %s %s.", function, component, version)
}


/**
 * Send a warning either to the debug log or in a PHP error depending on
 * $wgDevelopmentWarnings. To log warnings in production, use wfLogWarning() instead.
 *
 * @param string $msg Message to send
 * @param int $callerOffset Number of items to go back in the backtrace to
 *        find the correct caller (1 = function calling wfWarn, ...)
 * @param int $level PHP error level; defaults to E_USER_NOTICE;
 *        only used when $wgDevelopmentWarnings is true
 */
func WfWarn(msg string) {
	log.Error(msg)
}

/**
 * Send a warning as a PHP error and the debug log. This is intended for logging
 * warnings in production. For logging development warnings, use WfWarn instead.
 *
 * @param string $msg Message to send
 * @param int $callerOffset Number of items to go back in the backtrace to
 *        find the correct caller (1 = function calling wfLogWarning, ...)
 * @param int $level PHP error level; defaults to E_USER_WARNING
 */
func WfLogWarning(msg string) {
	log.Error(msg)
}

/**
 * Provide a simple HTTP error.
 *
 * @param int|string $code
 * @param string $label
 * @param string $desc
 */
func WfHttpError(ctx *context.Context, code int, label, desc string) error {
	ctx.ResponseWriter.Header().Set("Content-type", "text/html")
	ctx.ResponseWriter.Header().Set("charset", "utf-8")
	// TODO
	libs.NewHttpStatus().Header(ctx, code)
	NewHeaderCallback().WarnIfHeadersSent()
	content := fmt.Sprintf(
		"<!DOCTYPE html><html><head><title>%s</title></head><body><h1>%s</h1><p>%s</p></body></html>\n",
		php.Htmlspecialchars(label),
		php.Htmlspecialchars(label),
		php.Htmlspecialchars(desc),
	)
	 _, err := ctx.ResponseWriter.Write([]byte(content))
	return err
}

/**
 * Get the name of the function which called this function
 * wfGetCaller( 1 ) is the function with the wfGetCaller() call (ie. __FUNCTION__)
 * wfGetCaller( 2 ) [default] is the caller of the function running wfGetCaller()
 * wfGetCaller( 3 ) is the parent of that.
 *
 * @param int $level
 * @return string
 */
func WfGetCaller(level int) string {
	//TODO
	return "unknown"
}