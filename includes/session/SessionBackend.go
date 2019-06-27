/**
 * MediaWiki session backend
 */
package session

/**
 * This is the actual workhorse for Session.
 *
 * Most code does not need to use this class, you want \MediaWiki\Session\Session.
 * The exceptions are SessionProviders and SessionMetadata hook functions,
 * which get an instance of this class rather than Session.
 *
 * The reasons for this split are:
 * 1. A session can be attached to multiple requests, but we want the Session
 *    object to have some features that correspond to just one of those
 *    requests.
 * 2. We want reasonable garbage collection behavior, but we also want the
 *    SessionManager to hold a reference to every active session so it can be
 *    saved when the request ends.
 *
 * @ingroup Session
 * @since 1.27
 */
type SessionBackend struct {
	/** @var SessionId */
	id SessionId

	/** @var SessionProvider provider */
	provider SessionProvider
}

/**
 * Returns the session ID.
 * @return string
 */
func (s *SessionBackend) GetId() SessionId {
	return s.id
}

/**
 * Fetch the SessionId object
 * @private For internal use by WebRequest
 * @return SessionId
 */
func (s *SessionBackend) GetSessionId() SessionId {
	return s.id
}

/**
 * Changes the session ID
 * @return string New ID (might be the same as the old)
 */
func (s *SessionBackend) ResetId() SessionId {
	if s.provider.persistsSessionId() {

	}
	return s.id
}
