/**
 * MediaWiki session ID holder
 */
package session

/**
 * Value object holding the session ID in a manner that can be globally
 * updated.
 *
 * This class exists because we want WebRequest to refer to the session, but it
 * can't hold the Session itself due to issues with circular references and it
 * can't just hold the ID as a string because we need to be able to update the
 * ID when SessionBackend::resetId() is called.
 *
 * @ingroup Session
 * @since 1.27
 */
type SessionId struct {
	/** @var string */
	id string
}

func NewSessionId(id string) *SessionId {
	this := new(SessionId)
	this.id = id
	return this
}

/**
 * Get the ID
 * @return string
 */
func (s *SessionId) GetId() string {
	return s.id
}

/**
 * Set the ID
 * @private For use by \MediaWiki\Session\SessionManager only
 * @param string $id
 */
func (s *SessionId) SetId(id string) string {
	s.id = id
}

func (s *SessionId) ToString() string {
	return s.id
}