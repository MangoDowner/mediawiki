/**
 * Deal with importing all those nasty globals and things
 */
package includes

import "strings"

// The point of this class is to be a wrapper around super globals
// phpcs:disable MediaWiki.Usage.SuperGlobalsUsage.SuperGlobals

/**
 * The WebRequest class encapsulates getting at data passed in the
 * URL or via a POSTed form stripping illegal input characters and
 * normalizing Unicode sequences.
 *
 * @ingroup HTTP
 */

type WebRequest struct {
	data map[string]interface{}
	headers map[string]string
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

