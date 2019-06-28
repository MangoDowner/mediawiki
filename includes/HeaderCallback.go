package includes

import (
	"github.com/MangoDowner/mediawiki/includes/php"
	"github.com/astaxie/beego/logs"
)

type HeaderCallback struct {
	headersSentException error
	messageSent bool
}

func NewHeaderCallback() *HeaderCallback {
	this := new(HeaderCallback)
	return this
}

/**
 * Log a warning message if headers have already been sent. This can be
 * called before flushing the output.
 */
func (h *HeaderCallback) WarnIfHeadersSent() {
	if !php.HeadersSent() || h.messageSent {
		return
	}
	h.messageSent = true
	//TODO
	logs.Warn("Headers already sent, should send headers earlier than " +
		WfGetCaller(3	))
	logs.Error("Warning: headers were already sent from the location below")

}
