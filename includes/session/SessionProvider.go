/**
 * MediaWiki session provider base class
 */
package session

/**
 * A SessionProvider provides SessionInfo and support for Session
 *
 * A SessionProvider is responsible for taking a WebRequest and determining
 * the authenticated session that it's a part of. It does this by returning an
 * SessionInfo object with basic information about the session it thinks is
 * associated with the request, namely the session ID and possibly the
 * authenticated user the session belongs to.
 *
 * The SessionProvider also provides for updating the WebResponse with
 * information necessary to provide the client with data that the client will
 * send with later requests, and for populating the Vary and Key headers with
 * the data necessary to correctly vary the cache on these client requests.
 *
 * An important part of the latter is indicating whether it even *can* tell the
 * client to include such data in future requests, via the persistsSessionId()
 * and canChangeUser() methods. The cases are (in order of decreasing
 * commonness):
 *  - Cannot persist ID, no changing User: The request identifies and
 *    authenticates a particular local user, and the client cannot be
 *    instructed to include an arbitrary session ID with future requests. For
 *    example, OAuth or SSL certificate auth.
 *  - Can persist ID and can change User: The client can be instructed to
 *    return at least one piece of arbitrary data, that being the session ID.
 *    The user identity might also be given to the client, otherwise it's saved
 *    in the session data. For example, cookie-based sessions.
 *  - Can persist ID but no changing User: The request uniquely identifies and
 *    authenticates a local user, and the client can be instructed to return an
 *    arbitrary session ID with future requests. For example, HTTP Digest
 *    authentication might somehow use the 'opaque' field as a session ID
 *    (although getting MediaWiki to return 401 responses without breaking
 *    other stuff might be a challenge).
 *  - Cannot persist ID but can change User: I can't think of a way this
 *    would make sense.
 *
 * Note that many methods that are technically "cannot persist ID" could be
 * turned into "can persist ID but not change User" using a session cookie,
 * as implemented by ImmutableSessionProviderWithCookie. If doing so, different
 * session cookie names should be used for different providers to avoid
 * collisions.
 *
 * @ingroup Session
 * @since 1.27
 * @see https://www.mediawiki.org/wiki/Manual:SessionManager_and_AuthManager
 */
type SessionProvider struct {
	/** @var int Session priority. Used for the default newSessionInfo(), but
	 * could be used by subclasses too.
	 */
	priority int
}

/**
 * @note To fully initialize a SessionProvider, the setLogger(),
 *  setConfig(), and setManager() methods must be called (and should be
 *  called in that order). Failure to do so is liable to cause things to
 *  fail unexpectedly.
 */
func NewSessionProvider() *SessionProvider {
	this := new(SessionProvider)
	this.priority = MIN_PRIORITY + 10
	return this
}