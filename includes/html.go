/**
 * Collection of methods to generate HTML content
 */
package includes

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/php"
	"html"
	"reflect"
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
func (h *Html) GetTextInputAttributes(attrs map[string]string) map[string]string {
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
	strings.ReplaceAll(contents, "&", "&amp;")
	strings.ReplaceAll(contents, "<", "&lt;")
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
func (h *Html) OpenElement(element string, attribs map[string]interface{}) string {
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
	if element == "button" && attribs["type"] == "" {
		attribs["type"] = "submit"
	}
	return "<$element" + h.ExpandAttributes(h.DropDefaults(element, attribs)) + ">"
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
		lcattrib := strings.ToLower(attrib)
		// Simple checks using $attribDefaults
		if v, ok := attribDefaults[element]; ok {
			if v1, ok := v[lcattrib]; ok && v1 == value {
				delete(attribs, attrib)
			}
		}

		if lcattrib == "class" && value == "" {
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
		for _, v := range attribs {
			if "multiple" != v {
				continue
			}
			// A multi-select
			if attribs["size"] == "4" {
				delete(attribs, "size")
			}
			return attribs
		}
		if attribs["multiple"] != "" && attribs["multiple"] != "false" {
			// A multi-select
			if attribs["size"] == "4" {
				delete(attribs, "size")
			}
		} else {
			// Single select
			if attribs["size"] == "1" {
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
		for _, v := range HtmlBoolAttribs {
			if v == key {
				ret = fmt.Sprintf("%s %s=\"\"", ret, key)
				return ret
			}
		}
		switch value.(type) {
		case int:
			valueStr = fmt.Sprintf("%d", value)
		case float32:
			valueStr = strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32)
		case float64:
			valueStr = strconv.FormatFloat(value.(float64), 'f', -1, 64)
		}
		valueStr = html.EscapeString(valueStr)
		ret = fmt.Sprintf("%s %s=%s%s%s", ret, key, quote, php.EncodeAttribute(valueStr), quote)
	}
	return ret
}
