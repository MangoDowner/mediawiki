package includes

import (
	"fmt"
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