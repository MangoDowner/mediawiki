/**
 * Representation of a title within %MediaWiki.
 */
package includes

import (
	"github.com/MangoDowner/mediawiki/includes/consts"
	"github.com/MangoDowner/mediawiki/includes/linker"
	"github.com/MangoDowner/mediawiki/includes/php"
	"github.com/MangoDowner/mediawiki/includes/title"
)

/**
 * Represents a title within MediaWiki.
 * Optionally may contain an interwiki designation or namespace.
 * @note This class can fetch various kinds of data from the database;
 *       however, it does so inefficiently.
 * @note Consider using a TitleValue object instead. TitleValue is more lightweight
 *       and does not rely on global state or the database.
 */

/**
 * Title::newFromText maintains a cache to avoid expensive re-normalization of
 * commonly used titles. On a batch operation this can become a memory leak
 * if not bounded. After hitting this many titles reset the cache.
 */
const CACHE_CACHE_MAX = 1000
/**
 * Used to be GAID_FOR_UPDATE define. Used with getArticleID() and friends
 * to use the master DB
 */
const CACHE_GAID_FOR_UPDATE = 1

type Title struct {
	/** @var MapCacheLRU */
	titleCache interface{}

	/**
	 * @name Private member variables
	 * Please use the accessor functions instead.
	 * @private
	 */
	// @{

	/** @var string Text form (spaces not underscores) of the main part */
	MTextform string

	/** @var string URL-encoded form of the main part */
	MUrlform string

	/** @var string Main part with underscores */
	MDbkeyform string

	/** @var string Database key with the initial letter in the case specified by the user */
	mUserCaseDBKey string

	/** @var int Namespace index, i.e. one of the NS_xxxx constants */
	MNamespace int

	/** @var string Interwiki prefix */
	MInterwiki string

	/** @var bool Was this Title created from a string with a local interwiki prefix? */
	mLocalInterwiki bool

	/** @var string Title fragment (i.e. the bit after the #) */
	MFragment string

	/** @var int Article ID, fetched from the link cache on demand */
	MArticleID int

	/** @var bool|int ID of most recent revision */
	mLatestID int

	/**
	 * @var bool|string ID of the page's content model, i.e. one of the
	 *   CONTENT_MODEL_XXX constants
	 */
	mContentModel string

	/**
	 * @var bool If a content model was forced via setContentModel()
	 *   this will be true to avoid having other code paths reset it
	 */
	mForcedContentModel bool

	/** @var int Estimated number of revisions; null of not loaded */
	mEstimateRevisions int

	/** @var array Array of groups allowed to edit this article */
	MRestrictions []string

	/**
	 * @var string|bool Comma-separated set of permission keys
	 * indicating who can move or edit the page from the page table, (pre 1.10) rows.
	 * Edit and move sections are separated by a colon
	 * Example: "edit=autoconfirmed,sysop:move=sysop"
	 */
	mOldRestrictions bool

	/** @var bool Cascade restrictions on this page to included templates and images? */
	MCascadeRestriction bool

	/** Caching the results of getCascadeProtectionSources */
	MCascadingRestrictions interface{}

	/** @var array When do the restrictions on this page expire? */
	mRestrictionsExpiry []string

	/** @var bool Are cascading restrictions in effect on this page? */
	mHasCascadingRestrictions bool

	/** @var array Where are the cascading restrictions coming from on this page? */
	MCascadeSources []string

	/** @var bool Boolean for initialisation on demand */
	MRestrictionsLoaded bool

	/**
	 * Text form including namespace/interwiki, initialised on demand
	 *
	 * Only public to share cache with TitleFormatter
	 *
	 * @private
	 * @var string
	 */
	PrefixedText interface{}

	/** @var mixed Cached value for getTitleProtection (create protection) */
	MTitleProtection interface{}

	/**
	 * @var int Namespace index when there is no namespace. Don't change the
	 *   following default, NS_MAIN is hardcoded in several places. See T2696.
	 *   Zero except in {{transclusion}} tags.
	 */
	MDefaultNamespace interface{}

	/** @var int The page length, 0 for special pages */
	mLength int

	/** @var null Is the article at this title a redirect? */
	MRedirect interface{}

	/** @var array Associative array of user ID -> timestamp/false */
	mNotificationTimestamp []string

	/** @var bool Whether a page has any subpages */
	mHasSubpages bool

	/** @var bool The (string) language code of the page's language and content code. */
	mPageLanguage bool

	/** @var string|bool|null The page language code from the database, null if not saved in
	 * the database or false if not loaded, yet. */
	mDbPageLanguage string

	/** @var TitleValue A corresponding TitleValue object */
	mTitleValue interface{}

	/** @var bool Would deleting this page be a big deletion? */
	mIsBigDeletion bool
}

func NewTitle() *Title {
	this := new(Title)
	return this
}

/**
 * Create a new Title from a TitleValue
 *
 * @param TitleValue $titleValue Assumed to be safe.
 *
 * @return Title
 */
func (t *Title) NewFromTitleValue(titleValue *title.TitleValue) *Title {
	return t.NewFromLinkTarget(titleValue)
}

/**
 * Create a new Title from a TitleValue
 *
 * @param TitleValue $titleValue Assumed to be safe.
 *
 * @return Title
 */
func (t *Title) NewFromLinkTarget(linkTarget linker.LinkTarget) *Title {
	// Special case i
	// f it's already a Title object
	// TODO
	return t.MakeTitle(
		linkTarget.GetNamespace(),
		linkTarget.GetText(),
		linkTarget.GetFragment(),
		linkTarget.GetInterwiki(),
	)
}

/**
 * Create a new Title from a namespace index and a DB key.
 *
 * It's assumed that $ns and $title are safe, for instance when
 * they came directly from the database or a special page name,
 * not from user input.
 *
 * No validation is applied. For convenience, spaces are normalized
 * to underscores, so that e.g. user_text fields can be used directly.
 *
 * @note This method may return Title objects that are "invalid"
 * according to the isValid() method. This is usually caused by
 * configuration changes: e.g. a namespace that was once defined is
 * no longer configured, or a character that was once allowed in
 * titles is now forbidden.
 *
 * @param int $ns The namespace of the article
 * @param string $title The unprefixed database key form
 * @param string $fragment The link fragment (after the "#")
 * @param string $interwiki The interwiki prefix
 * @return Title The new object
 */
func (t *Title) MakeTitle(ns int, title, fragment, interwiki string) *Title {
	t.MInterwiki = interwiki
	t.MFragment = fragment
	t.MNamespace = ns
	t.MDbkeyform = php.Strtr(title,  map[string]string{" " : "_"})
	if ns >= 0 {
		t.MArticleID = -1
	} else {
		t.MArticleID = 0
	}
	t.MUrlform = WfUrlencode(t.MDbkeyform)
	t.MTextform = php.Strtr(title, map[string]string{"_" : " "})
	t.mContentModel = "" // initialized lazily in getContentModel()
	return t
}

/**
 * Is this in a namespace that allows actual pages?
 *
 * @return bool
 */
func (t *Title) CanExist() bool {
	return t.MNamespace >= consts.NS_MAIN
}