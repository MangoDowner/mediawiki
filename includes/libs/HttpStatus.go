/**
 * List of HTTP status codes.
 */
package libs

import (
	"github.com/MangoDowner/mediawiki/includes/consts"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

/**
 * @todo document
 */

type HttpStatus struct {

}

/**
 * Get the message associated with an HTTP response status code
 *
 * @param int $code Status code
 * @return string|null Message, or null if $code is not known
 */
func NewHttpStatus() *HttpStatus {
	this := new(HttpStatus)
	return this
}

/**
 * Log a warning message if headers have already been sent. This can be
 * called before flushing the output.
 */
func (h *HttpStatus) GetMessage(ctx *context.Context, code int) string {
	statusMessage := map[int]string{
		100 : "Continue",
		101 : "Switching Protocols",
		102 : "Processing",
		200 : "OK",
		201 : "Created",
		202 : "Accepted",
		203 : "Non-Authoritative Information",
		204 : "No Content",
		205 : "Reset Content",
		206 : "Partial Content",
		207 : "Multi-Status",
		300 : "Multiple Choices",
		301 : "Moved Permanently",
		302 : "Found",
		303 : "See Other",
		304 : "Not Modified",
		305 : "Use Proxy",
		307 : "Temporary Redirect",
		400 : "Bad Request",
		401 : "Unauthorized",
		402 : "Payment Required",
		403 : "Forbidden",
		404 : "Not Found",
		405 : "Method Not Allowed",
		406 : "Not Acceptable",
		407 : "Proxy Authentication Required",
		408 : "Request Timeout",
		409 : "Conflict",
		410 : "Gone",
		411 : "Length Required",
		412 : "Precondition Failed",
		413 : "Request Entity Too Large",
		414 : "Request-URI Too Large",
		415 : "Unsupported Media Type",
		416 : "Request Range Not Satisfiable",
		417 : "Expectation Failed",
		422 : "Unprocessable Entity",
		423 : "Locked",
		424 : "Failed Dependency",
		428 : "Precondition Required",
		429 : "Too Many Requests",
		431 : "Request Header Fields Too Large",
		500 : "Internal Server Error",
		501 : "Not Implemented",
		502 : "Bad Gateway",
		503 : "Service Unavailable",
		504 : "Gateway Timeout",
		505 : "HTTP Version Not Supported",
		507 : "Insufficient Storage",
		511 : "Network Authentication Required",
	}
	msg := statusMessage[code]
	return msg
}


/**
 * Log a warning message if headers have already been sent. This can be
 * called before flushing the output.
 */
func (h *HttpStatus) Header(ctx *context.Context, code int) {
	message := h.GetMessage(ctx, code)
	if message == "" {
		logs.Error( "Unknown HTTP status code %d", consts.E_USER_WARNING)
		return
	}
	ctx.ResponseWriter.WriteHeader(code)
}
