package php

import "strings"

/**
 * @param string $str The string being translated.
 * @param array $replace_pairs The replace_pairs parameter may be used as a substitute for to and from in which case it's an array in the form array('from' => 'to', ...).
 * @return string A copy of str, translating all occurrences of each character in from to the corresponding character in to.
 */
func Strtr(content string, replaces map[string]string) string {
	for old, new := range replaces {
		content = strings.ReplaceAll(content, old, new)
	}
	return content
}
