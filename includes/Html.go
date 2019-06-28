/**
 * Collection of methods to generate HTML content
 */
package includes

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes/php"
	"html"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/**
 * This class is a collection of static functions that serve two purposes:
 *
 * 1) Implement any algorithms specified by HTML5, or other HTML
 * specifications, in a convenient and self-contained way.
 *
 * 2) Allow HTML elements to be conveniently and safely generated, like the
 * current Xml class but a) less confused (Xml supports HTML-specific things,
 * but only sometimes!) and b) not necessarily confined to XML-compatible
 * output.
 *
 * There are two important configuration options this class uses:
 *
 * $wgMimeType: If this is set to an xml MIME type then output should be
 *     valid XHTML5.
 *
 * This class is meant to be confined to utility functions that are called from
 * trusted code paths.  It does not do enforcement of policy like not allowing
 * <a> elements.
 *
 * @since 1.16
 */

// List of void elements from HTML5, section 8.1.2 as of 2016-09-19
var HtmlVoidElements = []string{
	"area", "base", "br", "col", "embed",
	"hr", "img", "input", "keygen", "link",
	"meta", "param", "source", "track", "wbr",
}

// Boolean attributes, which may have the value omitted entirely.  Manually
// collected from the HTML5 spec as of 2011-08-12.
var HtmlBoolAttribs = []string{
	"async", "autofocus", "autoplay", "checked", "controls",
	"default", "defer", "disabled", "formnovalidate", "hidden",
	"ismap", "itemscope", "loop", "multiple", "muted",
	"novalidate", "open", "pubdate", "readonly", "required",
	"reversed", "scoped", "seamless", "selected", "truespeed", "typemustmatch",
	// HTML5 Microdata
	"itemscope",
}

type Html struct{}

/**
 * Modifies a set of attributes meant for button elements
 * and apply a set of default attributes when $wgUseMediaWikiUIEverywhere enabled.
 * @param array $attrs HTML attributes in an associative array
 * @param string[] $modifiers classes to add to the button
 * @see https://tools.wmflabs.org/styleguide/desktop/index.html for guidance on available modifiers
 * @return array $attrs A modified attribute array
 */

func (h *Html) ButtonAttributes(attrs map[string]interface{}, modifiers []string) map[string]interface{} {
	if !WgUseMediaWikiUIEverywhere {
		return attrs
	}
	// ensure compatibility with Xml
	if attrs["class"] != "" {
		attrs["class"] = fmt.Sprintf("%s %s %s", attrs["class"], "mw-ui-button", strings.Join(modifiers, " "))
	} else {
		attrs["class"] = fmt.Sprintf("%s %s", "mw-ui-button", strings.Join(modifiers, " "))
	}
	return attrs
}

/**
 * Modifies a set of attributes meant for text input elements
 * and apply a set of default attributes.
 * Removes size attribute when $wgUseMediaWikiUIEverywhere enabled.
 * @param array $attrs An attribute array.
 * @return array $attrs A modified attribute array
 */
func (h *Html) GetTextInputAttributes(attrs map[string]interface{}) map[string]interface{} {
	if !WgUseMediaWikiUIEverywhere {
		return attrs
	}
	if attrs["class"] != "" {
		attrs["class"] = fmt.Sprintf("%s %s", attrs["class"], "mw-ui-button")
	} else {
		attrs["class"] = "mw-ui-button"
	}
	return attrs
}

/**
 * Returns an HTML link element in a string styled as a button
 * (when $wgUseMediaWikiUIEverywhere is enabled).
 *
 * @param string $contents The raw HTML contents of the element: *not*
 *   escaped!
 * @param array $attrs Associative array of attributes, e.g., [
 *   "href" => "https://www.mediawiki.org/" ]. See expandAttributes() for
 *   further documentation.
 * @param string[] $modifiers classes to add to the button
 * @see https://tools.wmflabs.org/styleguide/desktop/index.html for guidance on available modifiers
 * @return string Raw HTML
 */

func (h *Html) LinkButton(contents string, attrs map[string]interface{}, modifiers []string) string {
	return h.Element("a", h.ButtonAttributes(attrs, modifiers), contents)
}

/**
 * Returns an HTML element in a string.  The major advantage here over
 * manually typing out the HTML is that it will escape all attribute
 * values.
 *
 * This is quite similar to Xml::tags(), but it implements some useful
 * HTML-specific logic.  For instance, there is no $allowShortTag
 * parameter: the closing tag is magically omitted if $element has an empty
 * content model.
 *
 * @param string $element The element"s name, e.g., "a"
 * @param array $attribs Associative array of attributes, e.g., [
 *   "href" => "https://www.mediawiki.org/" ]. See expandAttributes() for
 *   further documentation.
 * @param string $contents The raw HTML contents of the element: *not*
 *   escaped!
 * @return string Raw HTML
 */
func (h *Html) RawElement(element string, attribs map[string]interface{}, contents string) string {
	start := h.OpenElement(element, attribs)
	for _, v := range HtmlVoidElements {
		if v == element {
			// Silly XML.
			return start[:len(start)-1] + "/>"
		}
	}
	return fmt.Sprintf("%s%s%s", start, contents, h.CloseElement(element))
}

/**
 * Identical to rawElement(), but HTML-escapes $contents (like
 * Xml::element()).
 *
 * @param string $element Name of the element, e.g., "a"
 * @param array $attribs Associative array of attributes, e.g., [
 *   "href" => "https://www.mediawiki.org/" ]. See expandAttributes() for
 *   further documentation.
 * @param string $contents
 *
 * @return string
 */
func (h *Html) Element(element string, attribs map[string]interface{}, contents string) string {
	if attribs == nil {
		attribs = make(map[string]interface{})
	}
	contents = php.Strtr(contents, map[string]string{
		"&": "&amp;",
		"<": "&lt;",
	})
	return h.RawElement(element, attribs, contents)
}

/**
 * Identical to rawElement(), but has no third parameter and omits the end
 * tag (and the self-closing "/" in XML mode for empty elements).
 *
 * @param string $element Name of the element, e.g., "a"
 * @param array $attribs Associative array of attributes, e.g., [
 *   "href" => "https://www.mediawiki.org/" ]. See expandAttributes() for
 *   further documentation.
 *
 * @return string
 */
func (h *Html) OpenElement(element string, attribs map[string]interface{}) (ret string) {

	// This is not required in HTML5, but let"s do it anyway, for
	// consistency and better compression.
	element = strings.ToLower(element)

	// Remove invalid input types
	if element == "input" {
		validTypes := []string{
			"hidden", "text", "password", "checkbox", "radio",
			"file", "submit", "image", "reset", "button",

			// HTML input types
			"datetime", "datetime-local", "date", "month", "time",
			"week", "number", "range", "email", "url",
			"search", "tel", "color",
		}
		if attribs["type"] != "" {
			exist := false
			for _, v := range validTypes {
				if v == attribs["type"] {
					exist = true
					break
				}
			}
			if !exist {
				delete(attribs, "type")
			}
		}
	}
	// According to standard the default type for <button> elements is "submit".
	// Depending on compatibility mode IE might use "button", instead.
	// We enforce the standard "submit".
	if element == "button" && (attribs["type"] == "" || attribs["type"] == nil) {
		attribs["type"] = "submit"
	}
	ret = fmt.Sprintf("<%s%s>", element, h.ExpandAttributes(h.DropDefaults(element, attribs)))
	return ret
}

/**
 * Returns "</$element>"
 *
 * @since 1.17
 * @param string $element Name of the element, e.g., 'a'
 * @return string A closing tag
 */
func (h *Html) CloseElement(element string) string {
	element = strings.ToLower(element)
	return fmt.Sprintf("</%s>", element)
}

/**
 * Given an element name and an associative array of element attributes,
 * return an array that is functionally identical to the input array, but
 * possibly smaller.  In particular, attributes might be stripped if they
 * are given their default values.
 *
 * This method is not guaranteed to remove all redundant attributes, only
 * some common ones and some others selected arbitrarily at random.  It
 * only guarantees that the output array should be functionally identical
 * to the input array (currently per the HTML 5 draft as of 2009-09-06).
 *
 * @param string $element Name of the element, e.g., 'a'
 * @param array $attribs Associative array of attributes, e.g., [
 *   'href' => 'https://www.mediawiki.org/' ].  See expandAttributes() for
 *   further documentation.
 * @return array An array of attributes functionally identical to $attribs
 */
func (h *Html) DropDefaults(element string, attribs map[string]interface{}) map[string]interface{} {
	// Whenever altering this array, please provide a covering test case
	// in HtmlTest::provideElementsWithAttributesHavingDefaultValues
	attribDefaults := map[string]map[string]string{
		"area": {"shape": "rect"},
		"button": {
			"formaction":  "GET",
			"formenctype": "application/x-www-form-urlencoded",
		},
		"canvas": {
			"height": "150",
			"width":  "300",
		},
		"form": {
			"action":       "GET",
			"autocomplete": "on",
			"enctype":      "application/x-www-form-urlencoded",
		},
		"input": {
			"formaction": "GET",
			"type":       "text",
		},
		"keygen": {"keytype": "rsa"},
		"link":   {"media": "all"},
		"menu":   {"type": "list"},
		"script": {"type": "text/javascript"},
		"style": {
			"media": "all",
			"type":  "text/css",
		},
		"textarea": {"wrap": "soft"},
	}

	element = strings.ToLower(element)

	for attrib, value := range attribs {
		// 还有数组的情况
		lcAttrib := strings.ToLower(attrib)
		// Simple checks using $attribDefaults
		if v, ok := attribDefaults[element]; ok {
			if v1, ok := v[lcAttrib]; ok && v1 == fmt.Sprint(value) {
				delete(attribs, attrib)
			}
		}

		if lcAttrib == "class" && (value == "" || value == nil) {
			delete(attribs, attrib)
		}
	}

	// More subtle checks
	if element == "link" && attribs["type"] == "text/css" {
		delete(attribs, "type")
	}

	if element == "input" {
		types := attribs["type"]
		value := attribs["value"]
		if types == "checkbox" || types == "radio" {
			// The default value for checkboxes and radio buttons is 'on'
			// not ''. By stripping value="" we break radio boxes that
			// actually wants empty values.
			if value == "on" {
				delete(attribs, "value")
			}
		} else if types == "submit" {
			// The default value for submit appears to be "Submit" but
			// let's not bother stripping out localized text that matches
			// that.
		} else {
			// The default value for nearly every other field type is ''
			// The 'range' and 'color' types use different defaults but
			// stripping a value="" does not hurt them.
			if value == "" {
				delete(attribs, "value")
			}
		}
	}
	if element == "select" && attribs["size"] != "" {
		if attribs["multiple"] != nil && attribs["multiple"] != false {
			// A multi-select
			if attribs["size"] == "4" || attribs["size"] == 4 {
				delete(attribs, "size")
			}
		} else {
			// Single select
			if attribs["size"] == "1" || attribs["size"] == 1 {
				delete(attribs, "size")
			}
		}
	}
	return attribs
}

/**
 * Given an associative array of element attributes, generate a string
 * to stick after the element name in HTML output.  Like [ 'href' =>
 * 'https://www.mediawiki.org/' ] becomes something like
 * ' href="https://www.mediawiki.org"'.  Again, this is like
 * Xml::expandAttributes(), but it implements some HTML-specific logic.
 *
 * Attributes that can contain space-separated lists ('class', 'accesskey' and 'rel') array
 * values are allowed as well, which will automagically be normalized
 * and converted to a space-separated string. In addition to a numerical
 * array, the attribute value may also be an associative array. See the
 * example below for how that works.
 *
 * @par Numerical array
 * @code
 *     Html::element( 'em', [
 *         'class' => [ 'foo', 'bar' ]
 *     ] );
 *     // gives '<em class="foo bar"></em>'
 * @endcode
 *
 * @par Associative array
 * @code
 *     Html::element( 'em', [
 *         'class' => [ 'foo', 'bar', 'foo' => false, 'quux' => true ]
 *     ] );
 *     // gives '<em class="bar quux"></em>'
 * @endcode
 *
 * @param array $attribs Associative array of attributes, e.g., [
 *   'href' => 'https://www.mediawiki.org/' ].  Values will be HTML-escaped.
 *   A value of false or null means to omit the attribute.  For boolean attributes,
 *   you can omit the key, e.g., [ 'checked' ] instead of
 *   [ 'checked' => 'checked' ] or such.
 *
 * @throws MWException If an attribute that doesn't allow lists is set to an array
 * @return string HTML fragment that goes between element name and '>'
 *   (starting with a space if at least one attribute is output)
 */
func (h *Html) ExpandAttributes(attribs map[string]interface{}) (ret string) {
	for key, value := range attribs {
		// Support intuitive [ 'checked' => true/false ] form
		if value == false || value == nil {
			continue
		}

		// For boolean attributes, support [ 'foo' ] instead of
		// requiring [ 'foo' => 'meaningless' ].
		// TODO: 这里有key为数字情况的判断
		//if ( is_int( $key ) && in_array( strtolower( $value ), self::$boolAttribs ) ) {
		//	$key = $value;
		//}

		// Not technically required in HTML5 but we'd like consistency
		// and better compression anyway.
		key = strings.ToLower(key)

		// https://www.w3.org/TR/html401/index/attributes.html ("space-separated")
		// https://www.w3.org/TR/html5/index.html#attributes-1 ("space-separated")
		spaceSeparatedListAttributes := map[string]string{
			"class":     "class",     // html4, html5
			"accesskey": "accesskey", // as of html5, multiple space-separated values allowed
			// html4-spec doesn't document rel= as space-separated
			// but has been used like that and is now documented as such
			// in the html5-spec.
			"rel": "rel",
		}
		var valueStr string
		// Specific features for attributes that allow a list of space-separated values
		if spaceSeparatedListAttributes[key] != "" {
			// Apply some normalization and remove duplicates
			// Convert into correct array. Array can contain space-separated
			// values. Implode/explode to get those into the main array as well.
			newValueMap := make(map[string]string)
			var newValue []string
			if reflect.ValueOf(value).Kind() == reflect.Slice {
				// If input wasn't an array, we can skip this step
				valueMap := value.([]string)
				for _, v := range valueMap {
					// String values should be normal `array( 'foo' )`
					// Just append them
					newValue = append(newValue, v)
				}
			} else {
				valueStr = value.(string)
				valueArr := strings.Split(valueStr, " ")
				for _, v := range valueArr {
					newValue = append(newValue, v)
				}
			}
			valueStr = strings.Join(newValue, " ")
			newValue = strings.Split(valueStr, " ")
			for _, v := range newValue {
				trim := strings.Trim(v, " ")
				if trim == "" {
					continue
				}
				newValueMap[v] = v
			}
			newValue = []string{}
			for _, v := range newValueMap {
				newValue = append(newValue, v)
			}
			valueStr = strings.Join(newValue, " ")
		} else if reflect.ValueOf(value).Kind() == reflect.Slice {
			panic(fmt.Sprintf("HTML attribute %s can not contain a list of values", key))
		}
		quote := `"`
		exist := false
		for _, v := range HtmlBoolAttribs {
			if v == key {
				exist = true
				break
			}
		}
		if exist {
			ret = fmt.Sprintf("%s %s=\"\"", ret, key)
			continue
		}
		switch value.(type) {
		case int:
			valueStr = fmt.Sprintf("%d", value)
		case float32:
			valueStr = strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32)
		case float64:
			valueStr = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		case string:
			valueStr = value.(string)
		}

		valueStr = html.EscapeString(valueStr)
		ret = fmt.Sprintf("%s %s=%s%s%s", ret, key, quote, php.EncodeAttribute(valueStr), quote)
	}
	return ret
}

/**
 * Output an HTML script tag with the given contents.
 *
 * It is unsupported for the contents to contain the sequence `<script` or `</script`
 * (case-insensitive). This ensures the script can be terminated easily and consistently.
 * It is the responsibility of the caller to avoid such character sequence by escaping
 * or avoiding it. If found at run-time, the contents are replaced with a comment, and
 * a warning is logged server-side.
 *
 * @param string $contents JavaScript
 * @param string|null $nonce Nonce for CSP header, from OutputPage::getCSPNonce()
 * @return string Raw HTML
 */
func (h *Html) InlineScript(contents, nonce string) (ret string) {
	attrs := make(map[string]interface{})
	if nonce != "" {
		attrs["nonce"] = nonce
	} else {
		//TODO: 省略了  RequestContext::getMain()->getConfig()
		if IsNonceRequired() {
			WfWarn("no nonce set on script. CSP will break it")
		}
	}
	// regexp.Match(`^(text|application)/xml$|^.+/.+\+xml$`, []byte(mimetype))
	if b, err := regexp.Match(`(?i).*</?script.*`, []byte(contents)); b && err == nil {
		WfLogWarning("Illegal character sequence found in inline script.")
		contents = `/* ERROR: Invalid script */`
	}
	return h.RawElement("script", attrs, contents)
}

/**
 * Output a "<script>" tag linking to the given URL, e.g.,
 * "<script src=foo.js></script>".
 *
 * @param string $url
 * @param string|null $nonce Nonce for CSP header, from OutputPage::getCSPNonce()
 * @return string Raw HTML
 */
func (h *Html) LinkedScript(url, nonce string) (ret string) {
	attrs := make(map[string]interface{})
	attrs["src"] = url
	if nonce != "" {
		attrs["nonce"] = nonce
	} else {
		//TODO: 省略了  RequestContext::getMain()->getConfig()
		if IsNonceRequired() {
			WfWarn("no nonce set on script. CSP will break it")
		}
	}
	return h.RawElement("script", attrs, "")
}

/**
 * Output a "<style>" tag with the given contents for the given media type
 * (if any).  TODO: do some useful escaping as well, like if $contents
 * contains literal "</style>" (admittedly unlikely).
 *
 * @param string $contents CSS
 * @param string $media A media type string, like 'screen'
 * @param array $attribs (since 1.31) Associative array of attributes, e.g., [
 *   'href' => 'https://www.mediawiki.org/' ]. See expandAttributes() for
 *   further documentation.
 * @return string Raw HTML
 */
func (h *Html) InlineStyle(contents, media string, attribs map[string]interface{}) (ret string) {
	if media == "" {
		media = "all"
	}
	// Don't escape '>' since that is used
	// as direct child selector.
	// Remember, in css, there is no "x" for hexadecimal escapes, and
	// the space immediately after an escape sequence is swallowed.
	contents = php.Strtr(contents, map[string]string{
		"<": `\3C `,
		// CDATA end tag for good measure, but the main security
		// is from escaping the '<'.
		"]]>": `\5D\5D\3E `,
	})

	if b, err := regexp.Match(`.*[<&].*`, []byte(contents)); b && err == nil {
		contents = fmt.Sprintf("/*<![CDATA[*/%s/*]]>*/", contents)
	}
	attribs["media"] = media
	ret = h.RawElement("style", attribs, contents)
	return ret
}

/**
 * Output a "<link rel=stylesheet>" linking to the given URL for the given
 * media type (if any).
 *
 * @param string $url
 * @param string $media A media type string, like 'screen'
 * @return string Raw HTML
 */
func (h *Html) LinkedStyle(url, media string) (ret string) {
	if media == "" {
		media = "all"
	}
	ret = h.Element("link", map[string]interface{}{
		"rel":   "stylesheet",
		"href":  url,
		"media": media,
	}, "")
	return ret
}

/**
 * Convenience function to produce an "<input>" element.  This supports the
 * new HTML5 input types and attributes.
 *
 * @param string $name Name attribute
 * @param string $value Value attribute
 * @param string $type Type attribute
 * @param array $attribs Associative array of miscellaneous extra
 *   attributes, passed to Html::element()
 * @return string Raw HTML
 */
func (h *Html) Input(name, value, types string, attribs map[string]interface{}) (ret string) {
	if types == "" {
		types = "text"
	}
	if attribs == nil {
		attribs = make(map[string]interface{})
	}
	attribs["type"] = types
	attribs["value"] = value
	attribs["name"] = name
	// TODO: use switch structure instead?
	if php.InArray(types, []string{"text", "search", "email", "password", "number"}) {
		attribs = h.GetTextInputAttributes(attribs)
	}
	if php.InArray(types, []string{"button", "reset", "submit"}) {
		attribs = h.ButtonAttributes(attribs, []string{})
	}
	ret = h.Element("input", attribs, "")
	return ret
}

/**
 * Convenience function to produce a checkbox (input element with type=checkbox)
 *
 * @param string $name Name attribute
 * @param bool $checked Whether the checkbox is checked or not
 * @param array $attribs Array of additional attributes
 * @return string Raw HTML
 */
func (h *Html) Check(name string, checked bool, attribs map[string]interface{}) (ret string) {
	// TODO: attribs needs to change map to slice,so attribs can sort
	var value string
	if attribs == nil {
		attribs = make(map[string]interface{})
	}
	if attribs["value"] != nil {
		value = attribs["value"].(string)
	} else {
		value = "1"
	}

	if checked {
		attribs["checked"] = "checked"
	}
	ret = h.Input(name, value, "checkbox", attribs)
	return ret
}

/**
 * Return a warning box.
 * @since 1.31
 * @param string $html of contents of box
 * @return string of HTML representing a warning box.
 */
func (h *Html) MessageBox(html, className, heading string) (ret string) {
	if heading != "" {
		html = h.Element("h2", nil, heading) + html
	}
	ret = h.RawElement("div", map[string]interface{}{
		"class": className,
	}, html)
	return ret
}

/**
 * Return a warning box.
 * @since 1.31
 * @param string $html of contents of box
 * @return string of HTML representing a warning box.
 */
func (h *Html) WarningBox(html string) (ret string) {
	ret = h.MessageBox(html, "warningbox", "")
	return ret
}

/**
 * Return an error box.
 * @since 1.31
 * @param string $html of contents of error box
 * @param string $heading (optional)
 * @return string of HTML representing an error box.
 */
func (h *Html) ErrorBox(html, heading string) (ret string) {
	ret = h.MessageBox(html, "errorbox", heading)
	return ret
}

/**
 * Return a success box.
 * @since 1.31
 * @param string $html of contents of box
 * @return string of HTML representing a success box.
 */
func (h *Html) SuccessBox(html string) (ret string) {
	ret = h.MessageBox(html, "successbox", "")
	return ret
}

/**
 * Convenience function to produce a radio button (input element with type=radio)
 *
 * @param string $name Name attribute
 * @param bool $checked Whether the radio button is checked or not
 * @param array $attribs Array of additional attributes
 * @return string Raw HTML
 */
func (h *Html) Radio(name string, checked bool, attribs map[string]interface{}) (ret string) {
	var value string
	if attribs == nil {
		attribs = make(map[string]interface{})
	}
	if attribs["value"] != nil {
		value = attribs["value"].(string)
	} else {
		value = "1"
	}

	if checked {
		attribs["checked"] = "checked"
	}
	ret = h.Input(name, value, "radio", attribs)
	return ret
}

/**
 * Convenience function for generating a label for inputs.
 *
 * @param string $label Contents of the label
 * @param string $id ID of the element being labeled
 * @param array $attribs Additional attributes
 * @return string Raw HTML
 */
func (h *Html) Label(label, id string, attribs map[string]interface{}) (ret string) {
	if attribs == nil {
		attribs = make(map[string]interface{})
	}
	attribs["for"] = id
	ret = h.Element("label", attribs, label)
	return ret
}

/**
 * Convenience function to produce an input element with type=hidden
 *
 * @param string $name Name attribute
 * @param string $value Value attribute
 * @param array $attribs Associative array of miscellaneous extra
 *   attributes, passed to Html::element()
 * @return string Raw HTML
 */
func (h *Html) Hidden(name, value string, attribs map[string]interface{}) (ret string) {
	if attribs == nil {
		attribs = make(map[string]interface{})
	}
	ret = h.Input(name, value, "hidden", attribs)
	return ret
}

/**
 * Convenience function to produce a <textarea> element.
 *
 * This supports leaving out the cols= and rows= which Xml requires and are
 * required by HTML4/XHTML but not required by HTML5.
 *
 * @param string $name Name attribute
 * @param string $value Value attribute
 * @param array $attribs Associative array of miscellaneous extra
 *   attributes, passed to Html::element()
 * @return string Raw HTML
 */
func (h *Html) Textarea(name, value string, attribs map[string]interface{}) (ret string) {
	if attribs == nil {
		attribs = make(map[string]interface{})
	}
	attribs["name"] = name
	var spacedValue string
	if value[0:1] == "\n" {
		// Workaround for T14130: browsers eat the initial newline
		// assuming that it's just for show, but they do keep the later
		// newlines, which we may want to preserve during editing.
		// Prepending a single newline
		spacedValue = "\n" + value
	} else {
		spacedValue = value
	}
	ret = h.Element("textarea", h.GetTextInputAttributes(attribs), spacedValue)
	return ret
}

/**
 * Helper for Html::namespaceSelector().
 * @param array $params See Html::namespaceSelector()
 * @return array
 */
func (h *Html) NamespaceSelectorOptions(params map[string]interface{}) (ret string) {
	options := make(map[string]string)

	if params["exclude"] == nil || reflect.ValueOf(params["disable"]).Kind() != reflect.Slice {
		params["exclude"] = []string{}
	}

	if params["all"] != nil {
		// add an option that would let the user select all namespaces.
		// Value is provided by user, the name shown is localized for the user.
		//TODO： 完全
		options[params["all"].(string)] = ""

	}
	//// Add all namespaces as options (in the content language)
	//$options +=
	//MediaWikiServices::getInstance()->getContentLanguage()->getFormattedNamespaces();
	//
	//$optionsOut = [];
	//// Filter out namespaces below 0 and massage labels
	//foreach ( $options as $nsId => $nsName ) {
	//if ( $nsId < NS_MAIN || in_array( $nsId, $params['exclude'] ) ) {
	//continue;
	//}
	//if ( $nsId === NS_MAIN ) {
	//// For other namespaces use the namespace prefix as label, but for
	//// main we don't use "" but the user message describing it (e.g. "(Main)" or "(Article)")
	//$nsName = wfMessage( 'blanknamespace' )->text();
	//} elseif ( is_int( $nsId ) ) {
	//$nsName = MediaWikiServices::getInstance()->getContentLanguage()->
	//convertNamespace( $nsId );
	//}
	//$optionsOut[$nsId] = $nsName;
	//}

	return ret
}
/**
 * Build a drop-down box for selecting a namespace
 *
 * @param array $params Params to set.
 * - selected: [optional] Id of namespace which should be pre-selected
 * - all: [optional] Value of item for "all namespaces". If null or unset,
 *   no "<option>" is generated to select all namespaces.
 * - label: text for label to add before the field.
 * - exclude: [optional] Array of namespace ids to exclude.
 * - disable: [optional] Array of namespace ids for which the option should
 *   be disabled in the selector.
 * @param array $selectAttribs HTML attributes for the generated select element.
 * - id:   [optional], default: 'namespace'.
 * - name: [optional], default: 'namespace'.
 * @return string HTML code to select a namespace.
 */
func (h *Html) NamespaceSelector(params map[string]interface{}, selectAttribs map[string]interface{}) (ret string) {
	//selectAttribsArr := php.Ksort(selectAttribs)

	// Is a namespace selected?
	if params["selected"] != nil {
		// If string only contains digits, convert to clean int. Selected could also
		// be "all" or "" etc. which needs to be left untouched.
		// PHP is_numeric() has issues with large strings, PHP ctype_digit has other issues
		// and returns false for already clean ints. Use regex instead..

		// else: leaves it untouched for later processing
	} else {
		params["selected"] = ""
	}

	if params["disable"] == nil || reflect.ValueOf(params["disable"]).Kind() != reflect.Slice {
		params["disable"] = []string{}
	}

	// Associative array between option-values and option-labels
	//options := h.NamespaceSelectorOptions(params)

	//// Convert $options to HTML
	//$optionsHtml = [];
	//foreach ( $options as $nsId => $nsName ) {
	//$optionsHtml[] = self::element(
	//'option', [
	//'disabled' => in_array( $nsId, $params['disable'] ),
	//'value' => $nsId,
	//'selected' => $nsId === $params['selected'],
	//], $nsName
	//);
	//}
	//
	//if ( !array_key_exists( 'id', $selectAttribs ) ) {
	//$selectAttribs['id'] = 'namespace';
	//}
	//
	//if ( !array_key_exists( 'name', $selectAttribs ) ) {
	//$selectAttribs['name'] = 'namespace';
	//}
	//
	//$ret = '';
	//if ( isset( $params['label'] ) ) {
	//$ret .= self::element(
	//'label', [
	//'for' => $selectAttribs['id'] ?? null,
	//], $params['label']
	//) . "\u{00A0}";
	//}
	//
	//// Wrap options in a <select>
	//$ret .= self::openElement( 'select', $selectAttribs )
	//. "\n"
	//. implode( "\n", $optionsHtml )
	//. "\n"
	//. self::closeElement( 'select' );
	//
	//return $ret;
	return ret
}

/**
 * Determines if the given MIME type is xml.
 *
 * @param string $mimetype MIME type
 * @return bool
 */
func (h *Html) IsXmlMimeType(mimetype string) bool {
	// https://html.spec.whatwg.org/multipage/infrastructure.html#xml-mime-type
	// * text/xml
	// * application/xml
	// * Any MIME type with a subtype ending in +xml (this implicitly includes application/xhtml+xml)
	b, err := regexp.Match(`^(text|application)/xml$|^.+/.+\+xml$`, []byte(mimetype))
	if !b || err != nil {
		return false
	}
	return true
}
