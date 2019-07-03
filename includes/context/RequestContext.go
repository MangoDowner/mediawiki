/**
 * Group all the pieces relevant to the context of a request into one instance
 */
package context

import (
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/astaxie/beego/logs"
)

/**
 * Group all the pieces relevant to the context of a request into one instance
 */
type RequestContext struct {
	/**
	 * @var WebRequest
	 */
	request includes.WebRequest

	/**
	 * @var Title
	 */
	title *includes.Title

	/**
	 * @var WikiPage
	 */
	wikipage interface{}

	/**
	 * @var OutputPage
	 */
	output interface{}

	/**
	 * @var User
	 */
	user interface{}

	/**
	 * @var Language
	 */
	lang interface{}

	/**
	 * @var Skin
	 */
	skin interface{}

	/**
	 * @var Timing
	 */
	timing interface{}

	/**
	 * @var Config
	 */
	config interface{}

	/**
	 * @var RequestContext
	 */
	instance interface{}

}

func NewRequestContext() *RequestContext {
	this := new(RequestContext)
	return this
}

/**
 * @param Title|null $title
 */
func (r *RequestContext) SetTitle(title *includes.Title) {
	r.title = title
	// Erase the WikiPage so a new one with the new title gets created.
	r.wikipage = nil
}

/**
 * @return Title|null
 */
func (r *RequestContext) GetTitle() *includes.Title {
	if r.title == nil {
		// fallback to $wg till we can improve this
		r.title = includes.WgTitle
		logs.Debug("GlobalTitleFail __METHOD__ called by wfGetAllCallers( 5 ) with no title set.")
	}
	return r.title
}

/**
 * Check, if a Title object is set
 *
 * @since 1.25
 * @return bool
 */
func (r *RequestContext) HasTitle() bool {
	return r.title != nil
}

/**
 * Check whether a WikiPage object can be get with getWikiPage().
 * Callers should expect that an exception is thrown from getWikiPage()
 * if this method returns false.
 *
 * @since 1.19
 * @return bool
 */
func (r *RequestContext) CanUseWikiPage() bool {
	if r.wikipage != nil {
		// If there's a WikiPage object set, we can for sure get it
		return true
	}
	// Only pages with legitimate titles can have WikiPages.
	// That usually means pages in non-virtual namespaces.
	title := r.GetTitle()
	if title == nil {
		return false
	}
	return r.title != nil
}
