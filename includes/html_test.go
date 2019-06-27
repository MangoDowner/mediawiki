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

/**
 * Test out Html::element drops or enforces default value
 * @covers Html::dropDefaults
 * @dataProvider provideElementsWithAttributesHavingDefaultValues
 */
func TestDropDefaults(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		"<area/>",
		h.Element("area", map[string]interface{}{
			"shape":"rect",
		}, ""),
		"Generic cases, match $attribDefault static array",
	)

	test.AssetEqual(
		`<button type="submit"></button>`,
		h.Element("button", map[string]interface{}{
			"formaction":"GET",
		}, ""),
		"Generic cases, match $attribDefault static array",
	)

	test.AssetEqual(
		`<button type="submit"></button>`,
		h.Element("button", map[string]interface{}{
			"formenctype":"application/x-www-form-urlencoded",
		}, ""),
		"Generic cases, match $attribDefault static array",
	)
	test.AssetEqual(
		`<canvas></canvas>`,
		h.Element("canvas", map[string]interface{}{
			"height":"150",
		}, ""),
		"Generic cases, match $attribDefault static array",
	)

	test.AssetEqual(
		`<canvas></canvas>`,
		h.Element("canvas", map[string]interface{}{
			"width":"300",
		}, ""),
		"Generic cases, match $attribDefault static array",
	)
	// Also check with numeric values
	test.AssetEqual(
		`<canvas></canvas>`,
		h.Element("canvas", map[string]interface{}{
			"width":150,
		}, ""),
		"Also check with numeric values",
	)

	test.AssetEqual(
		`<canvas></canvas>`,
		h.Element("canvas", map[string]interface{}{
			"width":300,
		}, ""),
		"Also check with numeric values",
	)

	test.AssetEqual(
		`<form></form>`,
		h.Element("form", map[string]interface{}{
			"action":"GET",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<form></form>`,
		h.Element("form", map[string]interface{}{
			"autocomplete":"on",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<form></form>`,
		h.Element("form", map[string]interface{}{
			"enctype":"application/x-www-form-urlencoded",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<input/>`,
		h.Element("input", map[string]interface{}{
			"type":"text",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<keygen/>`,
		h.Element("keygen", map[string]interface{}{
			"keytype":"rsa",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<link/>`,
		h.Element("link", map[string]interface{}{
			"media":"all",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<menu></menu>`,
		h.Element("menu", map[string]interface{}{
			"type":"list",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<script></script>`,
		h.Element("script", map[string]interface{}{
			"type":"text/javascript",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<style></style>`,
		h.Element("style", map[string]interface{}{
			"media":"all",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<style></style>`,
		h.Element("style", map[string]interface{}{
			"type":"text/css",
		}, ""),
		"",
	)

	test.AssetEqual(
		`<textarea></textarea>`,
		h.Element("textarea", map[string]interface{}{
			"wrap":"soft",
		}, ""),
		"",
	)

	// SPECIFIC CASES
	// <link type="text/css">
	test.AssetEqual(
		`<link/>`,
		h.Element("link", map[string]interface{}{
			"type":"text/css",
		}, ""),
		`<link type="text/css">`,
	)

	// <input> specific handling
	test.AssetEqual(
		`<input type="checkbox"/>`,
		h.Element("input", map[string]interface{}{
			"type":"checkbox", "value":"on",
		}, ""),
		`Default value "on" is stripped of checkboxes`,
	)
	test.AssetEqual(
		`<input type="radio"/>`,
		h.Element("input", map[string]interface{}{
			"type":"radio", "value":"on",
		}, ""),
		`Default value "on" is stripped of radio buttons`,
	)
	test.AssetEqual(
		`<input type="submit" value="Submit"/>`,
		h.Element("input", map[string]interface{}{
			"type":"submit", "value":"submit",
		}, ""),
		`Default value "Submit" is kept on submit buttons (for possible l10n issues)`,
	)
	test.AssetEqual(
		`<input type="color"/>`,
		h.Element("input", map[string]interface{}{
			"type":"color", "value":"",
		}, ""),
		``,
	)
	test.AssetEqual(
		`<input type="range"/>`,
		h.Element("input", map[string]interface{}{
			"type":"range", "value":"",
		}, ""),
		``,
	)

	// <button> specific handling
	// see remarks on https://msdn.microsoft.com/library/ms535211(v=vs.85).aspx
	test.AssetEqual(
		`<select multiple=""></select>`,
		h.Element("select", map[string]interface{}{
			"size":"4", "multiple":true,
		}, ""),
		`<select> specific handling`,
	)
	test.AssetEqual(
		`<select multiple=""></select>`,
		h.Element("select", map[string]interface{}{
			"size":4, "multiple":true,
		}, ""),
		`.. with numeric value`,
	)
	test.AssetEqual(
		`<select></select>`,
		h.Element("select", map[string]interface{}{
			"size":"1", "multiple":false,
		}, ""),
		``,
	)
	test.AssetEqual(
		`<select></select>`,
		h.Element("select", map[string]interface{}{
			"size":1, "multiple":false,
		}, ""),
		`.. with numeric value`,
	)

	// Passing an array as value
	test.AssetEqual(
		`<a class="css-class-one css-class-two"></a>`,
		h.Element("a", map[string]interface{}{
			"class": []string{"css-class-one", "css-class-two"},
		}, ""),
		`dropDefaults accepts values given as an array`,
	)

	// FIXME: doDropDefault should remove defaults given in an array
	// Expected should be '<a></a>'
	test.AssetEqual(
		`<a class=""></a>`,
		h.Element("a", map[string]interface{}{
			"class": []string{"", ""},
		}, ""),
		`dropDefaults accepts values given as an array`,
	)
}

/**
 * @dataProvider provideInlineScript
 * @covers Html::inlineScript
 */
func TestInlineScript(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		"<script></script>",
		h.InlineScript("", ""),
		"Empty",
	)
	test.AssetEqual(
		`<script>EXAMPLE.label("foo");</script>`,
		h.InlineScript(`EXAMPLE.label("foo");`, ""),
		"Simple",
	)
	test.AssetEqual(
		`<script>EXAMPLE.label("<a>");</script>`,
		h.InlineScript(`EXAMPLE.label("<a>");`, ""),
		"HTML",
	)
	test.AssetEqual(
		`<script>/* ERROR: Invalid script */</script>`,
		h.InlineScript(`EXAMPLE.label("</script>");`, ""),
		"Script closing string (lower)",
	)
	test.AssetEqual(
		`<script>/* ERROR: Invalid script */</script>`,
		h.InlineScript(`EXAMPLE.label("</SCriPT and STyLE>");`, ""),
		"Script closing with non-standard attributes (mixed)",
	)
	// In HTML, <script> contents aren't just plain CDATA until </script>,
	// there are levels of escaping modes, and the below sequence puts an
	// HTML parser in a state where </script> would *not* close the script.
	// https://html.spec.whatwg.org/multipage/parsing.html#script-data-double-escape-end-state
	test.AssetEqual(
		`<script>/* ERROR: Invalid script */</script>`,
		h.InlineScript(`var a = "<!--<script>";`, ""),
		"HTML-comment-open and script-open",
	)
}

/**
 * @covers Html::linkedScript
 */
func TestLinkedScript(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		"<script>//example.com/somescript.js</script>",
		h.InlineScript("//example.com/somescript.js", ""),
		"",
	)

	test.AssetEqual(
		"<script>../resources/lib/jquery/jquery.js</script>",
		h.InlineScript("../resources/lib/jquery/jquery.js", ""),
		"",
	)
}

/**
 * @covers Html::inlineStyle
 */
func TestInlineStyle(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		`<style media="pc">/*<![CDATA[*/pre[\3C &]after/*]]>*/</style>`,
		h.InlineStyle("pre[<&]after", "pc", map[string]interface{}{}),
		"",
	)
}

/**
 * @covers Html::input
 */
func TestWrapperInput(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		`<input type="radio" value="testval" name="testname"/>`,
		h.Input("testname", "testval", "radio", nil),
		"Input wrapper with type and value.",
	)

	test.AssetEqual(
		`<input name="testname"/>`,
		h.Input("testname", "", "", nil),
		"Input wrapper with all default values.",
	)
}

/**
 * @covers Html::check
 */
func TestWrapperCheck(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		`<input type="checkbox" value="1" name="testname"/>`,
		h.Check("testname", false, nil),
		"Checkbox wrapper unchecked.",
	)

	test.AssetEqual(
		`<input type="checkbox" value="1" name="testname"/>`,
		h.Check("testname", false, nil),
		"Checkbox wrapper checked.",
	)

	test.AssetEqual(
		`<input type="checkbox" value="testval" name="testname"/>`,
		h.Check("testname", false, map[string]interface{}{
			"value":"testval",
		}),
		"Checkbox wrapper with a value override.",
	)
}

/**
 * @covers Html::warningBox
 * @covers Html::messageBox
 */
func TestMessageBox(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		`<div class="messageBox"><h2>headingText</h2>message text</div>`,
		h.MessageBox("message text", "messageBox", "headingText"),
		"MessageBox with heading",
	)
}

/**
 * @covers Html::warningBox
 * @covers Html::messageBox
 */
func TestWarnBox(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		`<div class="warningbox">warn</div>`,
		h.WarningBox("warn"),
		"MessageBox with heading",
	)
}

/**
 * @covers Html::errorBox
 * @covers Html::messageBox
 */
func TestErrorBox(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		`<div class="errorbox">err</div>`,
		h.ErrorBox("err", ""),
		"MessageBox with heading",
	)

	test.AssetEqual(
		`<div class="errorbox"><h2>heading</h2>err</div>`,
		h.ErrorBox("err", "heading"),
		"MessageBox with heading",
	)

	test.AssetEqual(
		`<div class="errorbox"><h2>0</h2>err</div>`,
		h.ErrorBox("err", "0"),
		"MessageBox with heading",
	)
}

/**
 * @covers Html::successBox
 * @covers Html::messageBox
 */
func TestSuccessBox(t *testing.T) {
	h := new(Html)

	test.AssetEqual(
		`<div class="successbox">great</div>`,
		h.SuccessBox("great"),
		"",
	)

	test.AssetEqual(
		`<div class="successbox"><script>beware no escaping!</script></div>`,
		h.SuccessBox("<script>beware no escaping!</script>"),
		"",
	)
}