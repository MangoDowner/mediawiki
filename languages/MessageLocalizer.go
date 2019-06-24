package languages

/**
 * Interface for localizing messages in MediaWiki
 *
 * @since 1.30
 * @ingroup Language
 */

type MessageLocalizer interface {
	/**
	 * This is the method for getting translated interface messages.
	 *
	 * @see https://www.mediawiki.org/wiki/Manual:Messages_API
	 * @see Message::__construct
	 *
	 * @param string|string[]|MessageSpecifier $key Message key, or array of keys,
	 *   or a MessageSpecifier.
	 * @param mixed $params,... Normal message parameters
	 * @return Message
	 */
	 msg(key... string)
}