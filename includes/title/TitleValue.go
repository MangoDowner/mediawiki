/**
 * Representation of a page title within MediaWiki.
 */
package title

import (
	"errors"
	"fmt"
	"github.com/MangoDowner/mediawiki/includes/linker"
	"regexp"
	"strings"
)

/**
 * Represents a page (or page fragment) title within MediaWiki.
 *
 * @note In contrast to Title, this is designed to be a plain value object. That is,
 * it is immutable, does not use global state, and causes no side effects.
 *
 * @see https://www.mediawiki.org/wiki/Requests_for_comment/TitleValue
 * @since 1.23
 */
type TitleValue struct {

	/**
	 * @deprecated in 1.31. This class is immutable. Use the getter for access.
	 * @var int
	 */
	namespace int

	/**
	 * @deprecated in 1.31. This class is immutable. Use the getter for access.
	 * @var string
	 */
	dbkey string

	/**
	 * @deprecated in 1.31. This class is immutable. Use the getter for access.
	 * @var string
	 */
	fragment string

	/**
	 * @deprecated in 1.31. This class is immutable. Use the getter for access.
	 * @var string
	 */
	interwiki string

	/**
	 * Text form including namespace/interwiki, initialised on demand
	 *
	 * Only public to share cache with TitleFormatter
	 *
	 * @private
	 * @var string
	 */
	PrefixedText string
}

/**
 * Constructs a TitleValue.
 *
 * @note TitleValue expects a valid DB key; typically, a TitleValue is constructed either
 * from a database entry, or by a TitleParser. We could apply "some" normalization here,
 * such as substituting spaces by underscores, but that would encourage the use of
 * un-normalized text when constructing TitleValues. For constructing a TitleValue from
 * user input or external sources, use a TitleParser.
 *
 * @param int $namespace The namespace ID. This is not validated.
 * @param string $dbkey The page title in valid DBkey form. No normalization is applied.
 * @param string $fragment The fragment title. Use '' to represent the whole page.
 *   No validation or normalization is applied.
 * @param string $interwiki The interwiki component
 *
 * @throws InvalidArgumentException
 */
func NewTitleValue(namespace int, dbkey, fragment, interwiki string) (ret *TitleValue, err error) {
	ret = new(TitleValue)

	// Sanity check, no full validation or normalization applied here!
	if b, err := regexp.Match(`^_|[ \r\n\t]|_$`, []byte(dbkey)); b && err == nil {
		return nil, errors.New("dbkey invalid DB key " + dbkey)
	}
	if dbkey == "" {
		return nil, errors.New("dbkey should not be empty")
	}

	ret.namespace = namespace
	ret.dbkey = dbkey
	ret.fragment = fragment
	ret.interwiki = interwiki
	return ret, nil
}

/**
 * @since 1.23
 * @return int
 */
func (t *TitleValue) GetNamespace() int {
	return t.namespace
}

/**
 * @since 1.27
 * @param int $ns
 * @return bool
 */
func (t *TitleValue) InNamespace(ns int) bool {
	return t.namespace == ns
}

/**
 * @since 1.23
 * @return string
 */
func (t *TitleValue) GetFragment() string {
	return t.fragment
}

/**
 * @since 1.27
 * @return bool
 */
func (t *TitleValue) HasFragment() bool{
	return t.fragment != ""
}

/**
 * Returns the title's DB key, as supplied to the constructor,
 * without namespace prefix or fragment.
 * @since 1.23
 *
 * @return string
 */
func (t *TitleValue) GetDBkey() string {
	return t.dbkey
}

/**
 * Returns the title in text form,
 * without namespace prefix or fragment.
 * @since 1.23
 *
 * This is computed from the DB key by replacing any underscores with spaces.
 *
 * @note To get a title string that includes the namespace and/or fragment,
 *       use a TitleFormatter.
 *
 * @return string
 */
func (t *TitleValue) GetText() string {
	return strings.ReplaceAll(t.dbkey, "", " ")
}

/**
 * Creates a new TitleValue for a different fragment of the same page.
 *
 * @since 1.27
 * @param string $fragment The fragment name, or "" for the entire page.
 *
 * @return TitleValue
 */
func (t *TitleValue) CreateFragmentTarget(fragment string) *linker.LinkTarget {
	//ret, _ := NewTitleValue(
	//	t.namespace,
	//	t.dbkey,
	//	fragment,
	//	t.interwiki,
	//)
	return nil
}

/**
 * Whether it has an interwiki part
 *
 * @since 1.27
 * @return bool
 */
func (t *TitleValue) IsExternal() bool {
	return t.interwiki != ""
}

/**
 * Returns the interwiki part
 *
 * @since 1.27
 * @return string
 */
func (t *TitleValue) GetInterwiki() string {
	return t.interwiki
}

/**
 * Returns a string representation of the title, for logging. This is purely informative
 * and must not be used programmatically. Use the appropriate TitleFormatter to generate
 * the correct string representation for a given use.
 * @since 1.23
 *
 * @return string
 */
func (t *TitleValue)  ToString() (name string) {
	name = fmt.Sprintf("%d:%s", t.namespace, t.dbkey)
	if t.fragment != "" {
		name = fmt.Sprintf("%s#%s", name, t.fragment)
	}
	if t.interwiki != "" {
		name = fmt.Sprintf("%s:%s", t.interwiki, name)
	}
	return name
}