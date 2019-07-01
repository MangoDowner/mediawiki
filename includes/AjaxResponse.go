/**
 * Response handler for Ajax requests.
 */
package includes

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes/config"
	"github.com/MangoDowner/mediawiki/includes/libs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"time"
)

/**
 * Handle responses for Ajax requests (send headers, print
 * content, that sort of thing)
 *
 * @ingroup Ajax
 */
type AjaxResponse struct {
	/**
	 * Number of seconds to get the response cached by a proxy
	 * @var int $mCacheDuration
	 */
	mCacheDuration int

	/**
	 * HTTP header Content-Type
	 * @var string $mContentType
	 */
	mContentType string

	/**
	 * Disables output. Can be set by calling $AjaxResponse->disable()
	 * @var bool $mDisabled
	 */
	mDisabled bool

	/**
	 * Date for the HTTP header Last-modified
	 * @var string|bool $mLastModified
	 */
	mLastModified string

	/**
	 * HTTP response code
	 * @var string $mResponseCode
	 */
	mResponseCode int

	/**
	 * HTTP Vary header
	 * @var string $mVary
	 */
	mVary string

	/**
	 * Content of our HTTP response
	 * @var string $mText
	 */
	mText string

	/**
	 * @var Config
	 */
	mConfig beego.Config

	context *context.Context
}

/**
 * @param string|null $text
 * @param Config|null $config
 */
func NewAjaxResponse(ctx *context.Context,text string) *AjaxResponse {
	this := new(AjaxResponse)
	this.mResponseCode = 200
	this.mContentType = "application/x-wiki"
	this.context = ctx
	if text != "" {
		this.addText(text)
	}
	return this
}

/**
 * Output text
 */
func (a *AjaxResponse) PrintText() {
	if !a.mDisabled {
		a.context.ResponseWriter.Write([]byte(a.mText))
	}
}

/**
 * Add content to the response
 * @param string $text
 */
func (a *AjaxResponse) addText(text string) {
	if !a.mDisabled && text != "" {
		a.mText = a.mText + text
	}
}

func (a *AjaxResponse) SendHeaders() {
	hs := libs.NewHttpStatus()
	if a.mResponseCode != 0 {
		// For back-compat, it is supported that mResponseCode be a string like " 200 OK"
		// (with leading space and the status message after). Cast response code to an integer
		// to take advantage of PHP's conversion rules which will turn "  200 OK" into 200.
		// https://secure.php.net/manual/en/language.types.string.php#language.types.string.conversion
		hs.Header(a.context, a.mResponseCode)
	}

	a.context.ResponseWriter.Header().Set("Content-type", a.mContentType)

	if a.mLastModified != "" {
		a.context.ResponseWriter.Header().Set("Last-Modified", a.mLastModified)
	} else {
		a.context.ResponseWriter.Header().Set("Last-Modified",
			time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	}

	if a.mCacheDuration != 0 {
		// If CDN caches are configured, tell them to cache the response,
		// and tell the client to always check with the CDN. Otherwise,
		// tell the client to use a cached copy, without a way to purge it.
		if config.Configs.GetBool("UseSquid") {
			// Expect explicit purge of the proxy cache, but require end user agents
			// to revalidate against the proxy on each visit.
			// Surrogate-Control controls our CDN, Cache-Control downstream caches
			if  config.Configs.GetBool("UseESI") {
				a.context.ResponseWriter.Header().Set("Surrogate-Control",
					fmt.Sprintf(`max-age=%d, content=="ESI/1.0"`, a.mCacheDuration))
				a.context.ResponseWriter.Header().Set("Cache-Control",
					"s-maxage=0, must-revalidate, max-age=0")
			} else {
				a.context.ResponseWriter.Header().Set("Cache-Control",
					fmt.Sprintf(`s-maxage=%d, must-revalidate, max-age=0"`, a.mCacheDuration))
			}
		} else {
			// Let the client do the caching. Cache is not purged.
			a.context.ResponseWriter.Header().Set("Expires",
				time.Now().Add(time.Second * time.Duration(a.mCacheDuration)).
				Format("Mon, 02 Jan 2006 15:04:05 GMT"),
			)
			a.context.ResponseWriter.Header().Set("Cache-Control",
				fmt.Sprintf("s-maxage=%d, public,max-age= %d", a.mCacheDuration, a.mCacheDuration),
			)
		}

	} else {
		// always expired, always modified
		a.context.ResponseWriter.Header().Set("Expires",
			"Mon, 26 Jul 1997 05:00:00 GMT") // Date in the past
		a.context.ResponseWriter.Header().Set("Cache-Control",
			"no-cache, must-revalidate") // HTTP/1.1
		a.context.ResponseWriter.Header().Set("Pragma",
			"no-cache") // HTTP/1.0
	}

	if a.mVary != "" {
		a.context.ResponseWriter.Header().Set("Vary",
			a.mVary)
	}
}