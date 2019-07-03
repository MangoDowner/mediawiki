package includes

import (
	"github.com/MangoDowner/mediawiki/includes/languages"
	"github.com/MangoDowner/mediawiki/includes/libs"
	"reflect"
)

/**
 * The Message class provides methods which fulfil two basic services:
 *  - fetching interface messages
 *  - processing messages into a variety of formats
 *
 * First implemented with MediaWiki 1.17, the Message class is intended to
 * replace the old wfMsg* functions that over time grew unusable.
 * @see https://www.mediawiki.org/wiki/Manual:Messages_API for equivalences
 * between old and new functions.
 *
 * You should use the wfMessage() global function which acts as a wrapper for
 * the Message class. The wrapper let you pass parameters as arguments.
 *
 * The most basic usage cases would be:
 *
 * @code
 *     // Initialize a Message object using the 'some_key' message key
 *     $message = wfMessage( 'some_key' );
 *
 *     // Using two parameters those values are strings 'value1' and 'value2':
 *     $message = wfMessage( 'some_key',
 *          'value1', 'value2'
 *     );
 * @endcode
 *
 * @section message_global_fn Global function wrapper:
 *
 * Since wfMessage() returns a Message instance, you can chain its call with
 * a method. Some of them return a Message instance too so you can chain them.
 * You will find below several examples of wfMessage() usage.
 *
 * Fetching a message text for interface message:
 *
 * @code
 *    $button = Xml::button(
 *         wfMessage( 'submit' )->text()
 *    );
 * @endcode
 *
 * A Message instance can be passed parameters after it has been constructed,
 * use the params() method to do so:
 *
 * @code
 *     wfMessage( 'welcome-to' )
 *         ->params( $wgSitename )
 *         ->text();
 * @endcode
 *
 * {{GRAMMAR}} and friends work correctly:
 *
 * @code
 *    wfMessage( 'are-friends',
 *        $user, $friend
 *    );
 *    wfMessage( 'bad-message' )
 *         ->rawParams( '<script>...</script>' )
 *         ->escaped();
 * @endcode
 *
 * @section message_language Changing language:
 *
 * Messages can be requested in a different language or in whatever current
 * content language is being used. The methods are:
 *     - Message->inContentLanguage()
 *     - Message->inLanguage()
 *
 * Sometimes the message text ends up in the database, so content language is
 * needed:
 *
 * @code
 *    wfMessage( 'file-log',
 *        $user, $filename
 *    )->inContentLanguage()->text();
 * @endcode
 *
 * Checking whether a message exists:
 *
 * @code
 *    wfMessage( 'mysterious-message' )->exists()
 *    // returns a boolean whether the 'mysterious-message' key exist.
 * @endcode
 *
 * If you want to use a different language:
 *
 * @code
 *    $userLanguage = $user->getOption( 'language' );
 *    wfMessage( 'email-header' )
 *         ->inLanguage( $userLanguage )
 *         ->plain();
 * @endcode
 *
 * @note You can parse the text only in the content or interface languages
 *
 * @section message_compare_old Comparison with old wfMsg* functions:
 *
 * Use full parsing:
 *
 * @code
 *     // old style:
 *     wfMsgExt( 'key', [ 'parseinline' ], 'apple' );
 *     // new style:
 *     wfMessage( 'key', 'apple' )->parse();
 * @endcode
 *
 * Parseinline is used because it is more useful when pre-building HTML.
 * In normal use it is better to use OutputPage::(add|wrap)WikiMsg.
 *
 * Places where HTML cannot be used. {{-transformation is done.
 * @code
 *     // old style:
 *     wfMsgExt( 'key', [ 'parsemag' ], 'apple', 'pear' );
 *     // new style:
 *     wfMessage( 'key', 'apple', 'pear' )->text();
 * @endcode
 *
 * Shortcut for escaping the message too, similar to wfMsgHTML(), but
 * parameters are not replaced after escaping by default.
 * @code
 *     $escaped = wfMessage( 'key' )
 *          ->rawParams( 'apple' )
 *          ->escaped();
 * @endcode
 *
 * @section message_appendix Appendix:
 *
 * @todo
 * - test, can we have tests?
 * - this documentation needs to be extended
 *
 * @see https://www.mediawiki.org/wiki/WfMessage()
 * @see https://www.mediawiki.org/wiki/New_messages_API
 * @see https://www.mediawiki.org/wiki/Localisation
 *
 * @since 1.17
 */

/** Use message text as-is */
const MESSAGE_FORMAT_PLAIN = "plain"
/** Use normal wikitext -> HTML parsing (the result will be wrapped in a block-level HTML tag) */
const MESSAGE_FORMAT_BLOCK_PARSE = "block-parse"
/** Use normal wikitext -> HTML parsing but strip the block-level wrapper */
const MESSAGE_FORMAT_PARSE = "parse"
/** Transform {{..}} constructs but don't transform to HTML */
const MESSAGE_FORMAT_TEXT = "text"
/** Transform {{..}} constructs, HTML-escape the result */
const MESSAGE_FORMAT_ESCAPED = "escaped"

var listTypeMap = map[string]string {
	"comma" : "commaList",
	"semicolon" : "semicolonList",
	"pipe" : "pipeList",
	"text" : "listToText",
}

type Message struct {
	/**
	 * In which language to get this message. True, which is the default,
	 * means the current user language, false content language.
	 *
	 * @var bool
	 */
	interfaces bool

	/**
	 * In which language to get this message. Overrides the $interface setting.
	 *
	 * @var Language|bool Explicit language object, or false for user language
	 */
	language *languages.Language

	/**
	 * @var string The message key. If $keysToTry has more than one element,
	 * this may change to one of the keys to try when fetching the message text.
	 */
	key string

	/**
	 * @var string[] List of keys to try when fetching the message.
	 */
	keysToTry []string

	/**
	 * @var array List of parameters which will be substituted into the message.
	 */
	parameters []interface{}

	/**
	 * @var string
	 * @deprecated
	 */
	 format string

	/**
	 * @var bool Whether database can be used.
	 */
	useDatabase bool

	/**
	 * @var Title Title object to use as context.
	 */
	title string

	/**
	 * @var Content Content object representing the message.
	 */
	content string

	/**
	 * @var string
	 */
	message string
}

/**
 * @since 1.17
 * @param string|string[]|MessageSpecifier $key Message key, or array of
 * message keys to try and use the first non-empty message for, or a
 * MessageSpecifier to copy from.
 * @param array $params Message parameters.
 * @param Language|null $language [optional] Language to use (defaults to current user language).
 * @throws InvalidArgumentException
 */
func NewMessage(key interface{}, params []interface{}, language *languages.Language) *Message {
	if params == nil {
		params = make([]interface{}, 0)
	}
	// Whether key implements MessageSpecifier
	if _, ok := key.(libs.MessageSpecifier); ok {
		if len(params) != 0 {
			panic("$params must be empty if $key is a MessageSpecifier")
		}
		keyMS := key.(libs.MessageSpecifier)
		params = keyMS.GetParams()
		key = keyMS.GetKey()
	}
	if reflect.ValueOf(key).Kind() != reflect.String && reflect.ValueOf(key).Kind() != reflect.Slice {
		panic("$key must be a string or an array")
	}
	this := new(Message)
	// default values
	this.interfaces = true
	this.language = nil
	this.format = "parse"
	this.useDatabase = true

	this.keysToTry = key.([]string)
	if len(this.keysToTry) == 0 {
		panic("$key must not be an empty list")
	}
	this.key = this.keysToTry[0]
	this.parameters = params
	// User language is only resolved in getLanguage(). This helps preserve the
	// semantic intent of "user language" across serialize() and unserialize().
	this.language = language
	return this
}

/**
 * Returns the Language of the Message.
 *
 * @since 1.23
 *
 * @return Language
 */
func (m *Message) GetLanguage() *languages.Language {
	if m.language != nil {
		return m.language
	}
	// Defaults to false which means current user language
	//return NewRequestContext().GetMain().GetLanguage()
	return nil
}

/**
 * Adds parameters to the parameter list of this message.
 *
 * @since 1.17
 *
 * @param mixed $args,... Parameters as strings or arrays from
 *  Message::numParam() and the like, or a single array of parameters.
 *
 * @return Message $this
 */
func (m *Message) Params(params ...interface{}) *Message {
	var args []interface{}
	for _, v := range params {
		args = append(args, v)
	}
	// If $args has only one entry and it's an array, then it's either a
	// non-varargs call or it happens to be a call with just a single
	// "special" parameter. Since the "special" parameters don't have any
	// numeric keys, we'll test that to differentiate the cases.
	if len(args) == 1 && args[0] != nil && reflect.ValueOf(args).Kind() == reflect.Map {
		if len(args[0].(map[string]string)) == 0 {
			args = []interface{}{}
		} else {
			//TODO 删除了循环
		}
	}
	for _, v := range args {
		m.parameters = append(m.parameters, v)
	}
	return m
}

/**
 * Returns the message parsed from wikitext to HTML.
 *
 * @since 1.17
 *
 * @param string|null $format One of the FORMAT_* constants. Null means use whatever was used
 *   the last time (this is for B/C and should be avoided).
 *
 * @return string HTML
 * @suppress SecurityCheck-DoubleEscaped phan false positive
 */
func (m *Message) ToString(format string) (ret string) {
	if format == "" {
		//TODO: no warning text
		format = m.format
	}
	//strings := m.FetchMessage()
	return ret
}

/**
 * Fully parse the text from wikitext to HTML.
 *
 * @since 1.17
 *
 * @return string Parsed HTML.
 */
func (m *Message) Parse() (ret string) {
	m.format = MESSAGE_FORMAT_PARSE
	ret = m.ToString(MESSAGE_FORMAT_PARSE)
	return ret
}

/**
 * Returns the message text. {{-transformation is done.
 *
 * @since 1.17
 *
 * @return string Unescaped message text.
 */
func (m *Message) Text() (ret string) {
	m.format = MESSAGE_FORMAT_TEXT
	ret = m.ToString(MESSAGE_FORMAT_TEXT)
	return ret
}

/**
 * Returns the message text as-is, only parameters are substituted.
 *
 * @since 1.17
 *
 * @return string Unescaped untransformed message text.
 */
func (m *Message) Plain() (ret string) {
	m.format = MESSAGE_FORMAT_PLAIN
	ret = m.ToString(MESSAGE_FORMAT_PLAIN)
	return ret
}

/**
 * Returns the parsed message text which is always surrounded by a block element.
 *
 * @since 1.17
 *
 * @return string HTML
 */
func (m *Message) ParseAsBlock() (ret string) {
	m.format = MESSAGE_FORMAT_BLOCK_PARSE
	ret = m.ToString(MESSAGE_FORMAT_BLOCK_PARSE)
	return ret
}

/**
 * Returns the message text. {{-transformation is done and the result
 * is escaped excluding any raw parameters.
 *
 * @since 1.17
 *
 * @return string Escaped message text.
 */
func (m *Message) Escaped() (ret string) {
	m.format = MESSAGE_FORMAT_ESCAPED
	ret = m.ToString(MESSAGE_FORMAT_ESCAPED)
	return ret
}

/**
 * Wrapper for what ever method we use to get message contents.
 *
 * @since 1.17
 *
 * @return string
 * @throws MWException If message key array is empty.
 */
func (m *Message) FetchMessage() string {
	if m.message != "" {
		return m.message
	}
	var (
		key string
		message string
	)
	// TODO: 补充缓存代码
	//cache := cache2.SingletonMessageCache()
	//for _, k := range m.keysToTry {
	//	key = k
	//	message := cache.Get(key, m.useDatabase, m.getLanguage())
	//	if message != "" {
	//		break
	//	}
	//}
	// NOTE: The constructor makes sure keysToTry isn't empty,
	//       so we know that $key and $message are initialized.
	m.key = key
	m.message = message
	return m.message
}


