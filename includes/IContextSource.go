package includes

import (
	"github.com/MangoDowner/mediawiki/includes/languages"
)

/**
 * Interface for objects which can provide a MediaWiki context on request
 *
 * Context objects contain request-dependent objects that manage the core
 * web request/response logic for essentially all requests to MediaWiki.
 * The contained objects include:
 *   a) Key objects that depend (for construction/loading) on the HTTP request
 *   b) Key objects used for response building and PHP session state control
 *   c) Performance metric deltas accumulated from request execution
 *   d) The site configuration object
 * All of the objects are useful for the vast majority of MediaWiki requests.
 * The site configuration object is included on grounds of extreme
 * utility, even though it should not actually depend on the web request.
 *
 * More specifically, the scope of the context includes:
 *   a) Objects that represent the HTTP request/response and PHP session state
 *   b) Object representing the MediaWiki user (as determined by the HTTP request)
 *   c) Primary MediaWiki output builder objects (OutputPage, user skin object)
 *   d) The language object for the user/request
 *   e) The title and wiki page objects requested via URL (if any)
 *   f) Performance metric deltas accumulated from request execution
 *   g) The site configuration object
 *
 * This class is not intended as a service-locator nor a service singleton.
 * Objects that only depend on site configuration do not belong here (aside
 * from Config itself). Objects that represent persistent data stores do not
 * belong here either. Session state changes should only be propagated on
 * shutdown by separate persistence handler objects, for example.
 */

type IContextSource interface {

	/**
	 * @return WebRequest
	 */
	GetRequest() *WebRequest

	/**
	 * @return Title|null
	 */
	GetTitle() interface{}

	/**
	 * Check whether a WikiPage object can be get with getWikiPage().
	 * Callers should expect that an exception is thrown from getWikiPage()
	 * if this method returns false.
	 *
	 * @since 1.19
	 * @return bool
	 */
	CanUseWikiPage() bool

	/**
	 * Get the WikiPage object.
	 * May throw an exception if there's no Title object set or the Title object
	 * belongs to a special namespace that doesn't have WikiPage, so use first
	 * canUseWikiPage() to check whether this method can be called safely.
	 *
	 * @since 1.19
	 * @return WikiPage
	 */
	GetWikiPage(interface{}) interface{}

	/**
	 * @return OutputPage
	 */
	GetOutput() interface{}

	/**
	 * @return User
	 */
	GetUser() interface{}

	/**
	 * @return Language
	 * @since 1.19
	 */
	GetLanguage() *languages.Language

	/**
	 * @return Skin
	 */
	GetSkin() interface{}

	/**
	 * Get the site configuration
	 *
	 * @since 1.23
	 * @return Config
	 */
	GetConfig() interface{}

	/**
	 * @deprecated since 1.27 use a StatsdDataFactory from MediaWikiServices (preferably injected)
	 *
	 * @since 1.25
	 * @return IBufferingStatsdDataFactory
	 */
	GetStats() interface{}

	/**
	 * @since 1.27
	 * @return Timing
	 */
	GetTiming() interface{}

	/**
	 * Export the resolved user IP, HTTP headers, user ID, and session ID.
	 * The result will be reasonably sized to allow for serialization.
	 *
	 * @return array
	 * @since 1.21
	 */
	ExportSession() interface{}

	languages.MessageLocalizer
}