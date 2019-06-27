/**
 * MediaWiki cookie-based session provider interface
 */
package session

/**
 * A CookieSessionProvider persists sessions using cookies
 *
 * @ingroup Session
 * @since 1.27
 */
type CookieSessionProvider struct {
	params        map[string]interface{}
	cookieOptions map[string]string
	SessionProvider
}

/**
 * @param array $params Keys include:
 *  - priority: (required) Priority of the returned sessions
 *  - callUserSetCookiesHook: Whether to call the deprecated hook
 *  - sessionName: Session cookie name. Doesn't honor 'prefix'. Defaults to
 *    $wgSessionName, or $wgCookiePrefix . '_session' if that is unset.
 *  - cookieOptions: Options to pass to WebRequest::setCookie():
 *    - prefix: Cookie prefix, defaults to $wgCookiePrefix
 *    - path: Cookie path, defaults to $wgCookiePath
 *    - domain: Cookie domain, defaults to $wgCookieDomain
 *    - secure: Cookie secure flag, defaults to $wgCookieSecure
 *    - httpOnly: Cookie httpOnly flag, defaults to $wgCookieHttpOnly
 */
func NewCookieSessionProvider() *CookieSessionProvider {
	this := new(CookieSessionProvider)
	this.SessionProvider = *NewSessionProvider()
	// @codeCoverageIgnoreStart
	this.params["cookieOptions"] = []string{}
	return this
}
