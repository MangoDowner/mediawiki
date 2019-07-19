/**
 * Parent class for all special pages.
 */
package includes

import (
	"github.com/MangoDowner/mediawiki/includes/consts"
	"github.com/MangoDowner/mediawiki/includes/title"
)

/**
 * Parent class for all special pages.
 *
 * Includes some static functions for handling the special page list deprecated
 * in favor of SpecialPageFactory.
 *
 * @ingroup SpecialPage
 */
type SpecialPage struct {
	// The canonical name of this special page
	// Also used for the default <h1> heading, @see getDescription()
	mName string

	// The local name of this special page
	mLocalName string

	// Minimum user level required to access this page, or "" for anyone.
	// Also used to categorise the pages in Special:Specialpages
	mRestriction string

	// Listed in Special:Specialpages?
	mListed bool

	// Whether or not this special page is being included from an article
	mIncluding bool

	// Whether the special page can be included in an article
	mIncludable bool

	/**
	 * Current request context
	 * @var IContextSource
	 */
	mContext interface{}

	/**
	 * @var \MediaWiki\Linker\LinkRenderer|null
	 */
	linkRenderer interface{}
}

func NewSpecialPage() *SpecialPage {
	this := new(SpecialPage)
	return this
}

/**
 * Get a localised Title object for a specified special page name
 * If you don't need a full Title object, consider using TitleValue through
 * getTitleValueFor() below.
 *
 * @since 1.9
 * @since 1.21 $fragment parameter added
 *
 * @param string $name
 * @param string|bool $subpage Subpage string, or false to not use a subpage
 * @param string $fragment The link fragment (after the "#")
 * @return Title
 * @throws MWException
 */
func (w *SpecialPage) GetTitleFor(name, subpage, fragment string) (ret *Title) {
	return NewTitle().NewFromTitleValue(w.GetTitleValueFor(name, subpage, fragment))
}

/**
 * Get a localised TitleValue object for a specified special page name
 *
 * @since 1.28
 * @param string $name
 * @param string|bool $subpage Subpage string, or false to not use a subpage
 * @param string $fragment The link fragment (after the "#")
 * @return TitleValue
 */
func (w *SpecialPage) GetTitleValueFor(name, subpage, fragment string) (ret *title.TitleValue) {
	// TODO:
	ret, _ = title.NewTitleValue(consts.NS_SPECIAL, name, fragment, "")
	return ret
}