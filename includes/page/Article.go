/**
 * User interface for page actions.
 */
package page

import (
	"github.com/MangoDowner/mediawiki/includes"
)

/**
 * Class for viewing MediaWiki article and history.
 *
 * This maintains WikiPage functions for backwards compatibility.
 *
 * @todo Move and rewrite code to an Action class
 *
 * See design.txt for an overview.
 * Note: edit user interface and cache support functions have been
 * moved to separate EditPage and HTMLFileCache classes.
 */
type Article struct {
	/**
	 * @var IContextSource|null The context this Article is executed in.
	 * If null, RequestContext::getMain() is used.
	 */
	mContext *includes.IContextSource

	/** @var WikiPage|null The WikiPage object of this instance */
	mPage interface{}

	/**
	 * @var ParserOptions|null ParserOptions object for $wgUser articles.
	 * Initialized by getParserOptions by calling $this->mPage->makeParserOptions().
	 */
	MParserOptions interface{}

	/**
	 * @var Content|null Content of the main slot of $this->mRevision.
	 * @note This variable is read only, setting it has no effect.
	 *       Extensions that wish to override the output of Article::view should use a hook.
	 * @todo MCR: Remove in 1.33
	 * @deprecated since 1.32
	 * @since 1.21
	 */
	MContentObject interface{}

	/**
	 * @var bool Is the target revision loaded? Set by fetchRevisionRecord().
	 *
	 * @deprecated since 1.32. Whether content has been loaded should not be relevant to
	 * code outside this class.
	 */
	MContentLoaded bool

	/**
	 * @var int|null The oldid of the article that was requested to be shown,
	 * 0 for the current revision.
	 * @see $mRevIdFetched
	 */
	MOldId int

	/** @var Title|null Title from which we were redirected here, if any. */
	MRedirectedFrom interface{}

	/** @var string|bool URL to redirect to or false if none */
	MRedirectUrl string

	/**
	 * @var int Revision ID of revision that was loaded.
	 * @see $mOldId
	 * @deprecated since 1.32, use getRevIdFetched() instead.
	 */
	MRevIdFetched int

	/**
	 * @var Status|null represents the outcome of fetchRevisionRecord().
	 * $fetchResult->value is the RevisionRecord object, if the operation was successful.
	 *
	 * The information in $fetchResult is duplicated by the following deprecated public fields:
	 * $mRevIdFetched, $mContentLoaded. $mRevision (and $mContentObject) also typically duplicate
	 * information of the loaded revision, but may be overwritten by extensions or due to errors.
	 */
	fetchResult interface{}

	/**
	 * @var Revision|null Revision to be shown. Initialized by getOldIDFromRequest()
	 * or fetchContentObject(). Normally loaded from the database, but may be replaced
	 * by an extension, or be a fake representing an error message or some such.
	 * While the output of Article::view is typically based on this revision,
	 * it may be overwritten by error messages or replaced by extensions.
	 */
	MRevision interface{}

	/**
	 * @var ParserOutput|null|false The ParserOutput generated for viewing the page,
	 * initialized by view(). If no ParserOutput could be generated, this is set to false.
	 * @deprecated since 1.32
	 */
	MParserOutput interface{}

	/**
	 * @var bool Whether render() was called. With the way subclasses work
	 * here, there doesn't seem to be any other way to stop calling
	 * OutputPage::enableSectionEditLinks() and still have it work as it did before.
	 */
	disableSectionEditForRender bool
}

