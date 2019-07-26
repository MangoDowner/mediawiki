/**
 * Deal with importing all those nasty globals and things
 */
package includes

import (
	"github.com/MangoDowner/mediawiki/includes/php"
	"github.com/astaxie/beego/context"
	"reflect"
	"strings"
	"time"
)

// The point of this class is to be a wrapper around super globals
// phpcs:disable MediaWiki.Usage.SuperGlobalsUsage.SuperGlobals

/**
 * The WebRequest class encapsulates getting at data passed in the
 * URL or via a POSTed form stripping illegal input characters and
 * normalizing Unicode sequences.
 *
 * @ingroup HTTP
 */

/**
 * Flag to make WebRequest::getHeader return an array of values.
 * @since 1.26
 */
const GETHEADER_LIST = 1

type WebRequest struct {
	// FIXME: 新增字段，使用beego的context
	context *context.Context
	data map[string]string
	headers map[string]string

	/**
	 * The unique request ID.
	 * @var string
	 */
	reqId string

	/**
	 * Lazy-init response object
	 * @var WebResponse
	 */
	response interface{}

	/**
	 * Cached client IP address
	 * @var string
	 */
	ip string

	/**
	 * The timestamp of the start of the request, with microsecond precision.
	 * @var float
	 */
	requestTime int64

	/**
	 * Cached URL protocol
	 * @var string
	 */
	protocol string

	/**
	 * @var SessionId|null Session ID to use for this
	 *  request. We can't save the session directly due to reference cycles not
	 *  working too well (slow GC in Zend and never collected in HHVM).
	 */
	sessionId interface{}

	/** @var bool Whether this HTTP request is "safe" (even if it is an HTTP post) */
	markedAsSafe bool
}

/**
 * @codeCoverageIgnore
 */
func NewWebRequest(context *context.Context) *WebRequest {
	this := new(WebRequest)
	this.requestTime = time.Now().UnixNano()
	// POST overrides GET data
	// We don't use $_REQUEST here to avoid interference from cookies...
	// TODO: 此处用beego的params代替php的$_POST + $_GET;
	this.context = context
	this.data = context.Input.Params()
	return this
}

/**
 * Fetch a value from the given array or return $default if it's not set.
 *
 * @param array $arr
 * @param string $name
 * @param mixed $default
 * @return mixed
 */
func (w *WebRequest) GetGPCVal(arr map[string]string, name string, defaultVal interface{}) (result interface{}) {
	// PHP is so nice to not touch input data, except sometimes:
	// https://secure.php.net/variables.external#language.variables.external.dot-in-names
	// Work around PHP *feature* to avoid *bugs* elsewhere.
	// PHP 不会改变传递给脚本中的变量名。然而应该注意到点（句号）不是 PHP 变量名中的合法字符
	name = php.Strtr(name, map[string]string{"." : "_"})
	data, ok := arr[name]
	if !ok {
		return defaultVal
	}
	// TODO： 补充
	nameVal := ""
	//nameVal := beego.Controller{}.GetString(name)
	if nameVal != "" && reflect.ValueOf(data).Kind() != reflect.Slice {
		// Check for alternate/legacy character encoding.
		//TODO: 省略了一些代码
	}
	return data
}


/**
 * Fetch a scalar from the input without normalization, or return $default
 * if it's not set.
 *
 * Unlike self::getVal(), this does not perform any normalization on the
 * input value.
 *
 * @since 1.28
 * @param string $name
 * @param string|null $default
 * @return string|null
 */
func (w *WebRequest) GetRawVal(name string, defaultVal interface{}) (result interface{}) {
	name = strings.Replace(name, ".", "_", -1 )  // See comment in self::getGPCVal()
	// TODO: 还需要判断是否为数组
	if val, ok := w.data[name]; ok {
		result = val
	} else {
		result = defaultVal
	}
	return result
}

/**
 * Fetch a scalar from the input or return $default if it's not set.
 * Returns a string. Arrays are discarded. Useful for
 * non-freeform text inputs (e.g. predefined internal text keys
 * selected by a drop-down menu). For freeform input, see getText().
 *
 * @param string $name
 * @param string|null $default Optional default (or null)
 * @return string|null
 */
func (w *WebRequest) GetVal(name string, defaultVal string) (result string) {
	val := w.GetGPCVal(w.data, name, defaultVal)
	if php.IsArray(val) {
		result = defaultVal
	}
	return result
}

/**
 * Fetch a boolean value from the input or return $default if not set.
 * Guaranteed to return true or false, with normal PHP semantics for
 * boolean interpretation of strings.
 *
 * @param string $name
 * @param bool $default
 * @return bool
 */
func (w *WebRequest) GetBool( name string, defaultVal bool ) bool {
	raw := w.GetRawVal(name, defaultVal)
	return raw.(bool)
}

/**
 * Fetch a boolean value from the input or return $default if not set.
 * Unlike getBool, the string "false" will result in boolean false, which is
 * useful when interpreting information sent from JavaScript.
 *
 * @param string $name
 * @param bool $default
 * @return bool
 */
func (w *WebRequest) GetFuzzyBool( name string, defaultVal bool ) bool {
	b := w.GetBool(name, defaultVal)
	raw := w.GetRawVal(name, nil)
	b1 := strings.EqualFold(raw.(string),"false")
	return b && b1
}

/**
 * Return true if the named value is set in the input, whatever that
 * value is (even "0"). RwasPostedeturn false if the named value is not set.
 * Example use is checking for the presence of check boxes in forms.
 *
 * @param string $name
 * @return bool
 */
func (w *WebRequest) GetCheck( name string) bool {
	// Checkboxes and buttons are only present when clicked
	// Presence connotes truth, absence false
	return w.GetRawVal(name, nil) != nil
}

/**
 * Get the HTTP method used for this request.
 *
 * @return string
 */
func (w *WebRequest) GetMethod() (method string) {
	method = w.context.Input.Method()
	if method == "" {
		method = "GET"
	}
	return method
}

/**
 * Return true if the named value is set in the input, whatever that
 * value is (even "0"). RwasPostedeturn false if the named value is not set.
 * Example use is checking for the presence of check boxes in forms.
 *
 * @param string $name
 * @return bool
 */
func (w *WebRequest) WasPosted() bool {
	return w.GetMethod() == "POST"
}



