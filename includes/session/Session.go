/**
 * MediaWiki session
 */
package session

import "github.com/astaxie/beego"

/**
 * Manages data for an an authenticated session
 *
 * A Session represents the fact that the current HTTP request is part of a
 * session. There are two broad types of Sessions, based on whether they
 * return true or false from self::canSetUser():
 * * When true (mutable), the Session identifies multiple requests as part of
 *   a session generically, with no tie to a particular user.
 * * When false (immutable), the Session identifies multiple requests as part
 *   of a session by identifying and authenticating the request itself as
 *   belonging to a particular user.
 *
 * The Session object also serves as a replacement for PHP's $_SESSION,
 * managing access to per-session data.
 *
 * @ingroup Session
 * @since 1.27
 */
type Session struct {
	/** @var null|string[] Encryption algorithm to use */
	encryptionAlgorithm []string

	/** @var SessionBackend Session backend */
	backend *SessionBackend

	/** @var int Session index */
	index int

	/** @var LoggerInterface */
	logger interface{}
}

/**
 * @param SessionBackend $backend
 * @param int $index
 * @param LoggerInterface $logger
 */
func NewSession(backend SessionBackend, index int, logger interface{}) *Session {
	this := new(Session)
	this.backend = backend
	this.index = index
	this.logger = logger
	return this
}


/**
 * Returns the session ID
 * @return string
 */
func (s *Session) GetId() string {
	return s.backend.GetId().ToString()
}

/**
 * Returns the SessionId object
 * @private For internal use by WebRequest
 * @return SessionId
 */
func (s *Session) GetSessionId() SessionId {
	return s.backend.GetSessionId()
}

/**
 * Changes the session ID
 * @return string New ID (might be the same as the old)
 */
func (s *Session) ResetId() string {
	return s.backend.ResetId()
}