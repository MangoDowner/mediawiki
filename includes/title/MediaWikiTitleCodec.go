/**
 * A codec for MediaWiki page titles.
 */
package title

import (
	"github.com/MangoDowner/mediawiki/includes/languages"
)

/**
 * A codec for MediaWiki page titles.
 *
 * @note Normalization and validation is applied while parsing, not when formatting.
 * It's possible to construct a TitleValue with an invalid title, and use MediaWikiTitleCodec
 * to generate an (invalid) title string from it. TitleValues should be constructed only
 * via parseTitle() or from a (semi)trusted source, such as the database.
 *
 * @see https://www.mediawiki.org/wiki/Requests_for_comment/TitleValue
 * @since 1.23
 */
type MediaWikiTitleCodec struct {
	/**
	 * @var Language
	 */
	language *languages.Language

	/**
	 * @var GenderCache
	 */
	genderCache interface{}

	/**
	 * @var string[]
	 */
	localInterWikis []string

	/**
	 * @var InterWikiLookup
	 */
	interWikiLookup interface{}
}

/**
 * @param Language $language The language object to use for localizing namespace names.
 * @param GenderCache $genderCache The gender cache for generating gendered namespace names
 * @param string[]|string $localInterwikis
 * @param InterwikiLookup|null $interwikiLookup
 */
func NewMediaWikiTitleCodec(language *languages.Language, genderCache interface{},
	localInterWikis []string, interwikiLookup interface{}) *MediaWikiTitleCodec {
	this := new(MediaWikiTitleCodec)
	this.language = language
	this.genderCache = genderCache
	this.localInterWikis = localInterWikis
	//this.interWikiLookup
	return this
}

/**
 * Clear internal caches
 *
 * For use in unit testing when namespace configuration is changed.
 *
 * @since 1.31
 */
func (m *MediaWikiTitleCodec) ClearCaches() {

}
