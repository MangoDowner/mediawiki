package context

import (
	"github.com/MangoDowner/mediawiki/includes/languages"
	"github.com/MangoDowner/mediawiki/includes/setup"
)

/**
 * Group all the pieces relevant to the context of a request into one instance
 */
type RequestContext struct {
	/**
	 * @var WebRequest
	 */
	request *IWebRequest

	/**
	 * @var Title
	 */
	title interface{}

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
	instance *IRequestContext
}


func NewRequestContext() *RequestContext {
	this := new(RequestContext)
	return this
}

/**
 * @return WebRequest
 */
func (m *RequestContext) GetRequest() *IWebRequest {
	if m.lang != nil {
		return m.request
	}
	if setup.WgCommandLineMode {
		// TODO: $this->request = new FauxRequest( [] );
		m.request = new(IWebRequest)
	} else {
		m.request = new(IWebRequest)
	}
	return m.request
}

/**
 * Get the Language object.
 * Initialization of user or request objects can depend on this.
 * @return Language
 * @throws Exception
 * @since 1.19
 */
func (m *RequestContext) GetLanguage() *languages.Language {
	if m.lang != nil {
		return m.lang
	}
	// TODO: 缺失代码
	return m.lang
}


/**
 * Get the RequestContext object associated with the main request
 *
 * @return RequestContext
 */
func (m *RequestContext) GetMain() *RequestContext {
	// TODO: 改为获取instance?
	//if m.instance == nil {
	//	m.instance = new(RequestContext)
	//}
	return m
}
