package includes

import (
	"fmt"
	test "github.com/MangoDowner/mediawiki/tests"
	"testing"
)

func TestButtonAttributes(t *testing.T) {
	WgUseMediaWikiUIEverywhere = true
	h := new(Html)
	attrs := map[string]interface{}{
		"class": "c1 c2 c3",
	}
	attrs1 := map[string]interface{}{}
	modifiers := []string{"m1", "m2"}

	result := h.ButtonAttributes(attrs, modifiers)
	fmt.Println(result)

	result1 := h.ButtonAttributes(attrs1, modifiers)
	fmt.Println(result1)
}

/**
 * @covers Html::element
 * @covers Html::rawElement
 * @covers Html::openElement
 * @covers Html::closeElement
 */
func TestElementBasics(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		"<img/>",
		h.Element("img", map[string]interface{}{}, ""),
		"Self-closing tag for short-tag elements",
	)

	test.AssetEqual(
		"<element></element>",
		h.Element("element", map[string]interface{}{}, ""),
		"Close tag for empty element (null, null)",
	)

	test.AssetEqual(
		"<element></element>",
		h.Element("element", map[string]interface{}{}, ""),
		"Close tag for empty element (array, string)",
	)

	test.AssetEqual(
		`<em class="bar quux"></em>`,
		h.Element("em", map[string]interface{}{
			"class" : []string{"bar", "quux"},
		}, ""),
		"Tag for em with class",
	)
}

func dataXmlMimeType() map[string]bool{
	return map[string]bool {
		// ( $mimetype, $isXmlMimeType )
		// HTML is not an XML MimeType
		"text/html" : false,
		// XML is an XML MimeType
		"text/xml" :  true,
		"application/xml" : true ,
		// XHTML is an XML MimeType
		"application/xhtml+xml" : true ,
		// Make sure other +xml MimeTypes are supported
		// SVG is another random MimeType even though we don't use it
		"image/svg+xml" : true ,
		// Complete random other MimeTypes are not XML
		"text/plain" : false ,
	}
}

/**
 * @covers Html::expandAttributes
 */
func TestXmlMimeType(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		false,
		h.IsXmlMimeType("text/html"),
		"HTML is not an XML MimeType",
	)

	test.AssetEqual(
		true,
		h.IsXmlMimeType("text/xml"),
		"XML is an XML MimeType",
	)

	test.AssetEqual(
		true,
		h.IsXmlMimeType("application/xml"),
		"XML is an XML MimeType",
	)

	test.AssetEqual(
		true,
		h.IsXmlMimeType("application/xhtml+xml"),
		"XHTML is an XML MimeType",
	)

	test.AssetEqual(
		true,
		h.IsXmlMimeType("image/svg+xml"),
		"SVG is another random MimeType even though we don't use it",
	)

	test.AssetEqual(
		false,
		h.IsXmlMimeType("text/plain"),
		"Complete random other MimeTypes are not XML",
	)



}

/**
 * @covers Html::expandAttributes
 */
func TestExpandAttributesForBooleans(t *testing.T) {
	h := new(Html)
	var result string

	// Boolean attributes do not generates output when value is false
	result = h.ExpandAttributes(map[string]interface{}{"selected": false})
	fmt.Println(result)

	// Boolean attributes do not generates output when value is null
	result = h.ExpandAttributes(map[string]interface{}{"selected": nil})
	fmt.Println(result)

	// Boolean attributes do not generates output when value is true
	result = h.ExpandAttributes(map[string]interface{}{"selected": true})
	fmt.Println(result)

	// Boolean attributes have no value when value is true (passed as numerical array)
	result = h.ExpandAttributes(map[string]interface{}{"selected": ""})
	fmt.Println(result)
}

/**
 * @covers Html::expandAttributes
 */
/**
 * @covers Html::expandAttributes
 */
func TestExpandAttributesForNumbers(t *testing.T) {
	h := new(Html)
	var result string

	// Integer value is cast to a string
	result = h.ExpandAttributes(map[string]interface{}{"value": 1})
	fmt.Println(result)

	// Float value is cast to a string
	result = h.ExpandAttributes(map[string]interface{}{"value": 1.1})
	fmt.Println(result)
}

/**
 * Html::expandAttributes has special features for HTML
 * attributes that use space separated lists and also
 * allows arrays to be used as values.
 * @covers Html::expandAttributes
 */

/**
 * Test for Html::expandAttributes()
 * Please note it output a string prefixed with a space!
 * @covers Html::expandAttributes
 */
func TestExpandAttributesVariousExpansions(t *testing.T) {
	h := new(Html)
	var result string
	// Empty string is always quoted
	result = h.ExpandAttributes(map[string]interface{}{"empty_string": ""})
	fmt.Println(result)

	// Simple string value needs no quotes
	result = h.ExpandAttributes(map[string]interface{}{"key": "value"})
	fmt.Println(result)

	// Number 1 value needs no quotes
	result = h.ExpandAttributes(map[string]interface{}{"one": 1})
	fmt.Println(result)

	// Number 0 value needs no quotes
	result = h.ExpandAttributes(map[string]interface{}{"zero": 0})
	fmt.Println(result)
}

func TestExpandAttributesListValueAttributes(t *testing.T) {
	h := new(Html)
	var result string
	// Normalization should strip redundant spaces
	result = h.ExpandAttributes(map[string]interface{}{"class": " redundant  spaces  here  "})
	fmt.Println(result)

	// Normalization should remove duplicates in string-lists
	result = h.ExpandAttributes(map[string]interface{}{"class": "foo bar foo bar bar"})
	fmt.Println(result)

	// Array with null, empty string and spaces
	result = h.ExpandAttributes(map[string]interface{}{
		//TODO: 与PHP不同如果为数组，则必为字符串数组
		"class": []string{"", " ", "  "},
	})
	fmt.Println(result)

	// Normalization should remove duplicates in the array
	result = h.ExpandAttributes(map[string]interface{}{
		"class": []string{"foo", "bar", "foo", "bar", "bar"},
	})
	fmt.Println(result)

	// Normalization should remove duplicates in string-lists in the array
	result = h.ExpandAttributes(map[string]interface{}{
		"class": []string{"foo bar", "bar foo", "foo", "bar bar"},
	})
	fmt.Println(result)
}

/**
 * Test feature added by r96188, let pass attributes values as
 * a PHP array. Restricted to class,rel, accesskey.
 * @covers Html::expandAttributes
 */
func TestExpandAttributesSpaceSeparatedAttributesWithBoolean(t *testing.T) {
	h := new(Html)
	var result string
	// Method use isset() internally, make sure we do discard
	// attributes values which have been assigned well known values
	//TODO: 情况不支持
	result = h.ExpandAttributes(map[string]interface{}{"class": map[string]string{
		//"booltrue": true,
		//"emptystring": "",
		//"boolfalse": false,
		//"boolfalse": 0,
		//"one": 1,
	}})
	fmt.Println(result)
}

/**
 * How do we handle duplicate keys in HTML attributes expansion?
 * We could pass a "class" the values: 'GREEN' and array( 'GREEN' => false )
 * The latter will take precedence.
 *
 * Feature added by r96188
 * @covers Html::expandAttributes
 */
func TestValueIsAuthoritativeInSpaceSeparatedAttributesArrays(t *testing.T) {
	h := new(Html)
	var result string
	result = h.ExpandAttributes(map[string]interface{}{"class": map[string]string{
		// TODO: 情况不支持
		//"GREEN": "",
		//"GREEN": false,
		//"GREEN": "",
	}})
	fmt.Println(result)
}

/**
 * @covers Html::expandAttributes
 * @expectedException MWException
 */
func TestExpandAttributes_ArrayOnNonListValueAttribute_ThrowsException(t *testing.T) {
	h := new(Html)
	var result string
	// Real-life test case found in the Popups extension (see Gerrit cf0fd64),
	// when used with an outdated BetaFeatures extension (see Gerrit deda1e7)
	result = h.ExpandAttributes(map[string]interface{}{"src": map[string]string{
		"ltr": "ltr.svg",
		"rtl": "rtl.svg",
	}})
	fmt.Println(result)
}
