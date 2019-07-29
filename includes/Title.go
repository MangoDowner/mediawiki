/**
 * Representation of a title within %MediaWiki.
 */
package includes

import (
	"fmt"
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
	PrefixedText string

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
 * THIS IS NOT THE FUNCTION YOU WANT. Use Title::newFromText().
 *
 * Example of wrong and broken code:
 * $title = Title::newFromURL( $wgRequest->getVal( 'title' ) );
 *
 * Example of right code:
 * $title = Title::newFromText( $wgRequest->getVal( 'title' ) );
 *
 * Create a new Title from URL-encoded text. Ensures that
 * the given title's length does not exceed the maximum.
 *
 * @param string $url The title, as might be taken from a URL
 * @return Title|null The new object, or null on an error
 */
func (t *Title) NewFromURL(url string) *Title {
	// TODO
	return t
}



/**
 * Get the namespace index, i.e. one of the NS_xxxx constants.
 *
 * @return int Namespace index
 */
func (t *Title) GetNamespace() int {
	return t.MNamespace
}

/**
 * Returns true if the title is inside the specified namespace.
 *
 * Please make use of this instead of comparing to getNamespace()
 * This function is much more resistant to changes we may make
 * to namespaces than code that makes direct comparisons.
 * @param int $ns The namespace
 * @return bool
 * @since 1.19
 */
func (t *Title) InNamespace(ns int) bool {
	return NewMWNamespace().Equals(t.MNamespace, ns)
}

/**
 * Get the Title fragment (i.e.\ the bit after the #) in text form
 *
 * Use Title::hasFragment to check for a fragment
 *
 * @return string Title fragment
 */
func (t *Title) GetFragment() string {
	return t.MFragment
}

/**
 * Check if a Title fragment is set
 *
 * @return bool
 * @since 1.23
 */
func (t *Title) HasFragment() bool {
	return t.MFragment != ""
}

/**
 * Get the main part with underscores
 *
 * @return string Main part of the title, with underscores
 */
func (t *Title) GetDBkey() string {
	return t.MDbkeyform
}

/**
 * Get the text form (spaces not underscores) of the main part
 *
 * @return string Main part of the title
 */
func (t *Title) GetText() string {
	return t.MTextform
}

/**
 * Creates a new Title for a different fragment of the same page.
 *
 * @since 1.27
 * @param string $fragment
 * @return Title
 */
func (t *Title) CreateFragmentTarget(fragment string) *Title {
	result := t.MakeTitle(t.MNamespace, t.GetText(), fragment, t.MInterwiki)
	return result
}

/**
 * Is this Title interwiki?
 *
 * @return bool
 */
func (t *Title) IsExternal() bool {
	return t.MInterwiki != ""
}

/**
 * Get the interwiki prefix
 *
 * Use Title::isExternal to check if a interwiki is set
 *
 * @return string Interwiki prefix
 */
func (t *Title) GetInterwiki() string {
	return t.MInterwiki
}

/**
 * Return a string representation of this title
 *
 * @return string Representation of this title
 */
func (t *Title) ToString() string {
	return t.GetPrefixedText()
}

/**
 * Get the prefixed title with spaces.
 * This is the form usually used for display
 *
 * @return string The prefixed title, with spaces
 */
func (t *Title) GetPrefixedText() string {
	if t.PrefixedText != "" {
		return t.PrefixedText
	}
	s := t.prefix(t.MTextform)
	s = php.Strtr(s, map[string]string{"_" : " "});
	t.PrefixedText = s
	return t.PrefixedText
}

/**
 * B/C kludge: provide a TitleParser for use by Title.
 * Ideally, Title would have no methods that need this.
 * Avoid usage of this singleton by using TitleValue
 * and the associated services when possible.
 *
 * @return TitleFormatter
 */
func (t *Title) GetTitleFormatter() {
	//TODO
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
	tn := NewTitle()
	tn.MInterwiki = interwiki
	tn.MFragment = fragment
	tn.MNamespace = ns
	tn.MDbkeyform = php.Strtr(title,  map[string]string{" " : "_"})
	if ns >= 0 {
		tn.MArticleID = -1
	} else {
		tn.MArticleID = 0
	}
	tn.MUrlform = WfUrlencode(t.MDbkeyform)
	tn.MTextform = php.Strtr(title, map[string]string{"_" : " "})
	tn.mContentModel = "" // initialized lazily in getContentModel()
	return tn
}


/**
 * Get the main part with underscores
 *
 * @return string Main part of the title, with underscores
 */
func (t *Title) GetDBKey() string {
	return t.MDbkeyform
}

/**
 * Is this in a namespace that allows actual pages?
 *
 * @return bool
 */
func (t *Title) CanExist() bool {
	return t.MNamespace >= consts.NS_MAIN
}

/**
 * Get the namespace text
 *
 * @return string|false Namespace text
 */
func (t *Title) GetNsText(name string) string {
	if t.IsExternal() {
		// This probably shouldn't even happen, except for interwiki transclusion.
		// If possible, use the canonical name for the foreign namespace.
		nsText := NewMWNamespace().GetCanonicalName(t.MNamespace)
		if nsText != "" {
			return nsText
		}
	}
	//TODO
	return ""
}


/**
 * Returns true if this is a special page.
 *
 * @return bool
 */

func (t *Title) IsSpecialPage() bool {
	return t.MNamespace == consts.NS_SPECIAL
}

/**
 * Returns true if this title resolves to the named special page
 *
 * @param string $name The special page name
 * @return bool
 */
func (t *Title) IsSpecial(name string) bool {
	if !t.IsSpecialPage() {
		return false
	}
	thisName := ""
	if name != thisName {
		return false
	}
	return true
}

/**
 * Prefix some arbitrary text with the namespace or interwiki prefix
 * of this object
 *
 * @param string $name The text
 * @return string The prefixed text
 */
func (t *Title) prefix(name string) string {
	p := ""
	if t.IsExternal() {
		p = t.MInterwiki + ":"
	}
	if 0 != t.MNamespace {
		nsText := t.GetNsText("")
		if nsText != "" {
			// See T165149. Awkward, but better than erroneously linking to the main namespace.

		}
		p = fmt.Sprintf("%s%s:", p, nsText)
	}
	return p + name
}

/**
 * Get the prefixed database key form
 *
 * @return string The prefixed title, with underscores and
 *  any interwiki and namespace prefixes
 */
func (t *Title) GetPrefixedDBkey() string {
	s := t.prefix(t.MDbkeyform)
	s = php.Strtr(s, map[string]string{" ": "_"})
	return s
}

/**
 * Helper to fix up the get{Canonical,Full,Link,Local,Internal}URL args
 * get{Canonical,Full,Link,Local,Internal}URL methods accepted an optional
 * second argument named variant. This was deprecated in favor
 * of passing an array of option with a "variant" key
 * Once $query2 is removed for good, this helper can be dropped
 * and the wfArrayToCgi moved to getLocalURL();
 *
 * @since 1.19 (r105919)
 * @param array|string $query
 * @param string|string[]|bool $query2
 * @return string
 */
func (t *Title) fixUrlQueryArgs(query, query2 string) (ret string) {
	if query2 != "" {
		WfDeprecated("Title::get{Canonical,Full,Link,Local,Internal}URL " +
		"method called with a second parameter is deprecated. Add your " +
		"parameter to an array passed as the first parameter.",
		"1.19", "", 0)
	}
	//TODO
	return ret
}

/**
 * Get a real URL referring to this title, with interwiki link and
 * fragment
 *
 * @see self::getLocalURL for the arguments.
 * @see wfExpandUrl
 * @param string|string[] $query
 * @param string|string[]|bool $query2
 * @param string|int|null $proto Protocol type to use in URL
 * @return string The URL
 */
func (t *Title) GetFullURL(query, query2 string, proto string) (ret string) {
	//TODO
	query = t.fixUrlQueryArgs(query, query2)

	// Hand off all the decisions on urls to getLocalURL
	//url = t.getLocalURL(query)

	return ret
}