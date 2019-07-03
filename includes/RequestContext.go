package includes

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes/languages"
	"github.com/MangoDowner/mediawiki/includes/setup"
	"github.com/astaxie/beego/logs"
)

/**
 * Group all the pieces relevant to the context of a request into one instance
 */
type RequestContext struct {
	/**
	 * @var WebRequest
	 */
	request *WebRequest

	/**
	 * @var Title
	 */
	title *Title

	/**
	 * @var WikiPage
	 */
	 wikipage interface{}

	/**
	 * @var OutputPage
	 */
	output *OutputPage

	/**
	 * @var User
	 */
	user interface{}

	/**
	 * @var Language
	 */
	lang *languages.Language

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
	instance *RequestContext
}


func NewRequestContext() *RequestContext {
	this := new(RequestContext)
	return this
}

/**
 * @param Config $config
 */
func (m *RequestContext) SetConfig(config interface{}) {
	//TODO
}

/**
 * @return Config
 */
func (m *RequestContext) GetConfig() interface{} {
	//TODO
	return m.config
}

/**
 * @return WebRequest
 */
func (m *RequestContext) GetRequest() *WebRequest {
	if m.lang != nil {
		return m.request
	}
	if setup.WgCommandLineMode {
		// TODO: $this->request = new FauxRequest( [] );
		m.request = new(WebRequest)
	} else {
		m.request = new(WebRequest)
	}
	return m.request
}

/**
 * @deprecated since 1.27 use a StatsdDataFactory from MediaWikiServices (preferably injected)
 *
 * @return IBufferingStatsdDataFactory
 */
func (m *RequestContext) GetStats() interface{} {
	//TODO
	return m.config
}

/**
 * @return Timing
 */
func (m *RequestContext) GetTiming() interface{} {
	//TODO
	return m.config
}

/**
 * @param Title|null $title
 */
func (m *RequestContext) SetTitle(title *Title) {
	m.title = title
	// Erase the WikiPage so a new one with the new title gets created.
	m.wikipage = nil
}

/**
 * @return Title|null
 */
func (m *RequestContext) GetTitle() *Title {
	if m.title != nil {
		return m.title
	}
	m.title = WgTitle
	logs.Debug(fmt.Sprintf("GlobalTitleFail %s called by with no title set", "__METHOD__",))
	return m.title
}

/**
 * Check, if a Title object is set
 *
 * @since 1.25
 * @return bool
 */
func (m *RequestContext) HasTitle() bool {
	return m.title != nil
}

/**
 * Check whether a WikiPage object can be get with getWikiPage().
 * Callers should expect that an exception is thrown from getWikiPage()
 * if this method returns false.
 *
 * @since 1.19
 * @return bool
 */
func (m *RequestContext) CanUseWikiPage() bool {
	//TODO
	return true
}

/**
 * Check whether a WikiPage object can be get with getWikiPage().
 * Callers should expect that an exception is thrown from getWikiPage()
 * if this method returns false.
 *
 * @since 1.19
 * @return bool
 */
func (m *RequestContext) SetWikiPage(wikiPage interface{}) {
	//TODO
}

/**
 * Get the WikiPage object.
 * May throw an exception if there's no Title object set or the Title object
 * belongs to a special namespace that doesn't have WikiPage, so use first
 * canUseWikiPage() to check whether this method can be called safely.
 *
 * @since 1.19
 * @throws MWException
 * @return WikiPage
 */
func (m *RequestContext) GetWikiPage(wikiPage interface{}) interface{} {
	//TODO
	return m.wikipage
}

/**
 * @param OutputPage $output
 */
func (m *RequestContext) SetOutput(output interface{}) {
	//TODO
}

/**
 * @param OutputPage $output
 */
func (r *RequestContext) GetOutput() *OutputPage {
	if r.output == nil {
		//TODO:
		r.output = NewOutputPage()
	}
	return r.output
}

/**
 * @param User $user
 */
func (m *RequestContext) SetUser(user interface{}) {
	//TODO
}

/**
 * @return User
 */
func (m *RequestContext) GetUser() interface{} {
	//TODO
	return m.user
}

/**
 * Get the Language object.
 * Initialization of user or request objects can depend on this.
 * @return Language
 * @throws Exception
 * @since 1.19
 */
func (r *RequestContext) GetLanguage() *languages.Language {
	if r.lang != nil {
		return r.lang
	}
	// TODO: 缺失代码
	return r.lang
}

/**
 * @param Skin $skin
 */
func (m *RequestContext) SetSkin(user interface{}) {
	//TODO
}

/**
 * @return Skin
 */
func (m *RequestContext) GetSkin() interface{} {
	//TODO
	return m.skin
}

/**
 * Get a Message object with context set
 * Parameters are the same as wfMessage()
 *
 * @param string|string[]|MessageSpecifier $key Message key, or array of keys,
 *   or a MessageSpecifier.
 * @param mixed $args,...
 * @return Message
 */
func (m *RequestContext) Msg(key ...string) languages.IMessage {
	// TODO: 改为获取instance?
	return nil
}

/**
 * Get the RequestContext object associated with the main request
 *
 * @return RequestContext
 */
func (m *RequestContext) GetMain() *IContextSource {
	// TODO: 改为获取instance?
	if m.instance == nil {
		m.instance = new(RequestContext)
	}
	//return m.instance
	return nil
}

/**
 * Export the resolved user IP, HTTP headers, user ID, and session ID.
 * The result will be reasonably sized to allow for serialization.
 *
 * @return array
 * @since 1.21
 */
func (m *RequestContext) ExportSession() interface{} {
	//TODO
	return m.skin
}
