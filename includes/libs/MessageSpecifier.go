package libs

type  MessageSpecifier interface {
	/**
	 * Returns the message key
	 *
	 * If a list of multiple possible keys was supplied to the constructor, this method may
	 * return any of these keys. After the message has been fetched, this method will return
	 * the key that was actually used to fetch the message.
	 *
	 * @return string
	 */
	GetKey() string

	/**
	 * Returns the message parameters
	 *
	 * @return array
	 */
	GetParams() []interface{}
}
