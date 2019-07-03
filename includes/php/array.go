package php

import (
	"reflect"
)

/**
 * Checks if a value exists in an array
 * @link http://php.net/manual/en/function.in-array.php
 * @param mixed $needle <p>
 * The searched value.
 * </p>
 * <p>
 * If needle is a string, the comparison is done
 * in a case-sensitive manner.
 * </p>
 * @param array $haystack <p>
 * The array.
 * </p>
 * @param bool $strict [optional] <p>
 * If the third parameter strict is set to true
 * then the in_array function will also check the
 * types of the
 * needle in the haystack.
 * </p>
 * @return bool true if needle is found in the array,
 * false otherwise.
 * @since 4.0
 * @since 5.0
 */
func InArray(needle string, haystack []string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

/**
 * Finds whether a variable is an array
 * @link http://php.net/manual/en/function.is-array.php
 * @param mixed $var <p>
 * The variable being evaluated.
 * </p>
 * @return bool true if var is an array,
 * false otherwise.
 * @since 4.0
 * @since 5.0
 */
func IsArray(varValue interface{}) bool {
	if reflect.ValueOf(varValue).Kind() != reflect.Slice {
		return false
	}
	return true
}