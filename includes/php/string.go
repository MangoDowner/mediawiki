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

/**
 * Case-insensitive version of <function>str_replace</function>.
 * @link http://php.net/manual/en/function.str-ireplace.php
 * @param mixed $search <p>
 * Every replacement with search array is
 * performed on the result of previous replacement.
 * </p>
 * @param mixed $replace <p>
 * </p>
 * @param mixed $subject <p>
 * If subject is an array, then the search and
 * replace is performed with every entry of
 * subject, and the return value is an array as
 * well.
 * </p>
 * @param int $count [optional] <p>
 * The number of matched and replaced needles will
 * be returned in count which is passed by
 * reference.
 * </p>
 * @return mixed a string or an array of replacements.
 * @since 5.0
 */
func StrIReplace(search, replace []string, subject string) string {
	// TODO: 需要无视大小写
	for k, v := range search {
		subject = strings.ReplaceAll(subject, v, replace[k])
	}
	return subject
}

/**
 * Split a string by string
 * @link http://php.net/manual/en/function.explode.php
 * @param string $delimiter <p>
 * The boundary string.
 * </p>
 * @param string $string <p>
 * The input string.
 * </p>
 * @param int $limit [optional] <p>
 * If limit is set and positive, the returned array will contain
 * a maximum of limit elements with the last
 * element containing the rest of string.
 * </p>
 * <p>
 * If the limit parameter is negative, all components
 * except the last -limit are returned.
 * </p>
 * <p>
 * If the limit parameter is zero, then this is treated as 1.
 * </p>
 * @return array If delimiter is an empty string (""),
 * explode will return false.
 * If delimiter contains a value that is not
 * contained in string and a negative
 * limit is used, then an empty array will be
 * returned. For any other limit, an array containing
 * string will be returned.
 * @since 4.0
 * @since 5.0
 */
func Explode(sep, s string, limit int) (ret []string) {
	arr := strings.Split(s, sep)
	if limit == 0 || limit >= len(arr) {
		return arr
	}
	for i := 0; i < limit - 1; i++ {
		ret = append(ret, arr[i])
	}
	ret = append(ret, strings.Join(arr[limit - 1:], sep))
	return ret
}