package includes

import "go-common/library/log"

/**
 * This is the function for getting translated interface messages.
 *
 * @see Message class for documentation how to use them.
 * @see https://www.mediawiki.org/wiki/Manual:Messages_API
 *
 * This function replaces all old wfMsg* functions.
 *
 * @param string|string[]|MessageSpecifier $key Message key, or array of keys, or a MessageSpecifier
 * @param string|string[] ...$params Normal message parameters
 * @return Message
 *
 * @since 1.17
 *
 * @see Message::__construct
 */
func WfMessage(key string, params ...string) *Message {
	message := NewMessage(key, nil, nil)

	// We call Message::params() to reduce code duplication
	if len(params) != 0 {
		message.Params(params)
	}

	return message
}

/**
 * Send a warning either to the debug log or in a PHP error depending on
 * $wgDevelopmentWarnings. To log warnings in production, use wfLogWarning() instead.
 *
 * @param string $msg Message to send
 * @param int $callerOffset Number of items to go back in the backtrace to
 *        find the correct caller (1 = function calling wfWarn, ...)
 * @param int $level PHP error level; defaults to E_USER_NOTICE;
 *        only used when $wgDevelopmentWarnings is true
 */
func WfWarn(msg string) {
	log.Error(msg)
}

/**
 * Send a warning as a PHP error and the debug log. This is intended for logging
 * warnings in production. For logging development warnings, use WfWarn instead.
 *
 * @param string $msg Message to send
 * @param int $callerOffset Number of items to go back in the backtrace to
 *        find the correct caller (1 = function calling wfLogWarning, ...)
 * @param int $level PHP error level; defaults to E_USER_WARNING
 */
func WfLogWarning(msg string) {
	log.Error(msg)
}
