/**
 * Provide things related to namespaces.
 */
package includes

import (
	"github.com/MangoDowner/mediawiki/includes/consts"
	"github.com/MangoDowner/mediawiki/includes/registration"
)

/**
 * This is a utility class with only static functions
 * for dealing with namespaces that encodes all the
 * "magic" behaviors of them based on index.  The textual
 * names of the namespaces are handled by Language.php.
 *
 * These are synonyms for the names given in the language file
 * Users and translators should not change them
 */
type MWNamespace struct {

	/**
	 * These namespaces should always be first-letter capitalized, now and
	 * forevermore. Historically, they could've probably been lowercased too,
	 * but some things are just too ingrained now. :)
	 */
	alwaysCapitalizedNamespaces []int

	/** @var string[]|null Canonical namespaces cache */
	canonicalNamespaces map[int]string

	/** @var array|false Canonical namespaces index cache */
	namespaceIndexes []int

	/** @var int[]|null Valid namespaces cache */
	validNamespaces []int

}

func NewMWNamespace() *MWNamespace {
	this := new(MWNamespace)
	this.alwaysCapitalizedNamespaces = []int{consts.NS_SPECIAL, consts.NS_USER, consts.NS_MEDIAWIKI}
	return this
}

/**
 * Returns whether the specified namespaces are the same namespace
 *
 * @note It's possible that in the future we may start using something
 * other than just namespace indexes. Under that circumstance making use
 * of this function rather than directly doing comparison will make
 * sure that code will not potentially break.
 *
 * @param int $ns1 The first namespace index
 * @param int $ns2 The second namespace index
 *
 * @return bool
 * @since 1.19
 */
func (m *MWNamespace) Equals(ns1, ns2 int) bool {
	return ns1 == ns2
}

/**
 * Clear internal caches
 *
 * For use in unit testing when namespace configuration is changed.
 *
 * @since 1.31
 */
func (m *MWNamespace) ClearCaches() {
	m.canonicalNamespaces = map[int]string{}
	m.namespaceIndexes = []int{}
	m.validNamespaces = []int{}
}

/**
 * Returns array of all defined namespaces with their canonical
 * (English) names.
 *
 * @param bool $rebuild Rebuild namespace list (default = false). Used for testing.
 *  Deprecated since 1.31, use self::clearCaches() instead.
 *
 * @return array
 * @since 1.17
 */
func (m *MWNamespace) GetCanonicalNamespaces(rebuild bool) map[int]string {
	if rebuild {
		m.ClearCaches()
	}

	if len(m.canonicalNamespaces) == 0 {
		m.canonicalNamespaces = WgCanonicalNamespaceNames
		m.canonicalNamespaces[consts.NS_MAIN] = ""
		// Add extension namespaces
		attrs := registration.NewExtensionRegistry().GetAttribute("ExtensionNamespaces")
		if len(attrs) > 0 {
			for k, v := range attrs {
				m.canonicalNamespaces[k] = v
			}
		}
		if len(WgExtraNamespaces) == 0 {
			for k, v := range WgExtraNamespaces {
				m.canonicalNamespaces[k] = v
			}
		}
		NewHooks().Run("CanonicalNamespaces", []interface{}{&m.canonicalNamespaces}, "")
	}
	return m.canonicalNamespaces
}

/**
 * Returns the canonical (English) name for a given index
 *
 * @param int $index Namespace index
 * @return string|bool If no canonical definition.
 */
func (m *MWNamespace) GetCanonicalName(index int) (ret string) {
	nsList := m.GetCanonicalNamespaces(false)
	if v, ok := nsList[index]; ok {
		ret = v
	}
	return ret
}