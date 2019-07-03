/**
 * Preparation for the final page rendering.
 */
package includes

import "math"

/**
 * This class should be covered by a general architecture document which does
 * not exist as of January 2011.  This is one of the Core classes and should
 * be read at least once by any new developers.
 *
 * This class is used to prepare the final rendering. A skin is then
 * applied to the output parameters (links, javascript, html, categories ...).
 *
 * @todo FIXME: Another class handles sending the whole page to the client.
 *
 * Some comments comes from a pairing session between Zak Greant and Antoine Musso
 * in November 2010.
 *
 * @todo document
 */
type OutputPage struct {

	/** @var array Should be private. Used with addMeta() which adds "<meta>" */
	mMetatags []string

	/** @var array */
	mLinktags []string

	/** @var bool */
	mCanonicalUrl bool

	/**
	 * @var string The contents of <h1> */
	mPageTitle string

	/**
	 * @var string The displayed title of the page. Different from page title
	 * if overridden by display title magic word or hooks. Can contain safe
	 * HTML. Different from page title which may contain messages such as
	 * "Editing X" which is displayed in h1. This can be used for other places
	 * where the page name is referred on the page.
	 */
	displayTitle string

	/**
	 * @var string Contains all of the "<body>" content. Should be private we
	 *   got set/get accessors and the append() method.
	 */
	mBodytext string

	/** @var string Stores contents of "<title>" tag */
	mHTMLtitle string

	/**
	 * @var bool Is the displayed content related to the source of the
	 *   corresponding wiki article.
	 */
	mIsArticle bool

	/** @var bool Stores "article flag" toggle. */
	mIsArticleRelated bool

	/** @var bool Is the content subject to copyright */
	mHasCopyright bool

	/**
	 * @var bool We have to set isPrintable(). Some pages should
	 * never be printed (ex: redirections).
	 */
	mPrintable bool

	/**
	 * @var array Contains the page subtitle. Special pages usually have some
	 *   links here. Don't confuse with site subtitle added by skins.
	 */
	mSubtitle []string

	/** @var string */
	mRedirect string

	/** @var int */
	mStatusCode int

	/**
	 * @var string Used for sending cache control.
	 *   The whole caching system should probably be moved into its own class.
	 */
	mLastModified string

	/** @var array */
	mCategoryLinks []string

	/** @var array */
	mCategories map[string][]string

	/** @var array */
	mIndicators []string

	/** @var array Array of Interwiki Prefixed (non DB key) Titles (e.g. 'fr:Test page') */
	mLanguageLinks []string

	/**
	 * Used for JavaScript (predates ResourceLoader)
	 * @todo We should split JS / CSS.
	 * mScripts content is inserted as is in "<head>" by Skin. This might
	 * contain either a link to a stylesheet or inline CSS.
	 */
	mScripts string

	/** @var string Inline CSS styles. Use addInlineStyle() sparingly */
	mInlineStyles string

	/**
	 * @var string Used by skin template.
	 * Example: $tpl->set( 'displaytitle', $out->mPageLinkTitle );
	 */
	MPageLinkTitle string

	/** @var array Array of elements in "<head>". Parser might add its own headers! */
	mHeadItems []string

	/** @var array Additional <body> classes; there are also <body> classes from other sources */
	mAdditionalBodyClasses []string

	/** @var array */
	mModules []string

	/** @var array */
	mModuleScripts []string

	/** @var array */
	mModuleStyles []string

	/** @var ResourceLoader */
	mResourceLoader interface{}

	/** @var ResourceLoaderClientHtml */
	rlClient interface{}

	/** @var ResourceLoaderContext */
	rlClientContext interface{}

	/** @var array */
	rlExemptStyleModules []string

	/** @var array */
	mJsConfigVars []string

	/** @var array */
	mTemplateIds []string

	/** @var array */
	mImageTimeKeys []string

	/** @var string */
	mRedirectCode string

	mFeedLinksAppendQuery string

	/** @var array
	 * What level of 'untrustworthiness' is allowed in CSS/JS modules loaded on this page?
	 * @see ResourceLoaderModule::$origin
	 * ResourceLoaderModule::ORIGIN_ALL is assumed unless overridden;
	 */
	mAllowedModules []interface{}

	/** @var bool Whether output is disabled.  If this is true, the 'output' method will do nothing. */
	mDoNothing bool

	// Parser related.

	/** @var int */
	mContainsNewMagic int

	/**
	 * lazy initialised, use parserOptions()
	 * @var ParserOptions
	 */
	mParserOptions interface{}

	/**
	 * Handles the Atom / RSS links.
	 * We probably only support Atom in 2011.
	 * @see $wgAdvertisedFeedTypes
	 */
	mFeedLinks []string

	// Gwicke work on squid caching? Roughly from 2003.
	mEnableClientCache bool

	/** @var bool Flag if output should only contain the body of the article. */
	mArticleBodyOnly bool

	/** @var bool */
	mNewSectionLink bool

	/** @var bool */
	mHideNewSectionLink bool

	/**
	 * @var bool Comes from the parser. This was probably made to load CSS/JS
	 * only if we had "<gallery>". Used directly in CategoryPage.php.
	 * Looks like ResourceLoader can replace this.
	 */
	mNoGallery bool

	/** @var int Cache stuff. Looks like mEnableClientCache */
	mCdnMaxage int
	/** @var int Upper limit on mCdnMaxage */
	mCdnMaxageLimit float64

	/**
	 * @var bool Controls if anti-clickjacking / frame-breaking headers will
	 * be sent. This should be done for pages where edit actions are possible.
	 * Setters: $this->preventClickjacking() and $this->allowClickjacking().
	 */
	mPreventClickjacking bool

	/** @var int To include the variable {{REVISIONID}} */
	mRevisionId int

	/** @var string */
	mRevisionTimestamp string

	/** @var array */
	mFileVersion []string

	/**
	 * @var array An array of stylesheet filenames (relative from skins path),
	 * with options for CSS media, IE conditions, and RTL/LTR direction.
	 * For internal use; add settings in the skin via $this->addStyle()
	 *
	 * Style again! This seems like a code duplication since we already have
	 * mStyles. This is what makes Open Source amazing.
	 */
	styles []string

	mIndexPolicy string
	mFollowPolicy string

	/**
	 * @var array Headers that cause the cache to vary.  Key is header name, value is an array of
	 * options for the Key header.
	 */
	mVaryHeader map[string][]string

	/**
	 * If the current page was reached through a redirect, $mRedirectedFrom contains the Title
	 * of the redirect.
	 *
	 * @var Title
	 */
	mRedirectedFrom interface{}

	/**
	 * Additional key => value data
	 */
	mProperties map[string]string

	/**
	 * @var string|null ResourceLoader target for load.php links. If null, will be omitted
	 */
	mTarget string

	/**
	 * @var bool Whether parser output contains a table of contents
	 */
	mEnableTOC bool

	/**
	 * @var string|null The URL to send in a <link> element with rel=license
	 */
	copyrightUrl string

	/** @var array Profiling data */
	limitReportJSData []string

	/** @var array Map Title to Content */
	contentOverrides map[string]string

	/** @var callable[] */
	contentOverrideCallbacks []interface{}

	/**
	 * Link: header contents
	 */
	mLinkHeader []string

	/**
	 * @var string The nonce for Content-Security-Policy
	 */
	cspNonce string

	/**
	 * @var array A cache of the names of the cookies that will influence the cache
	 */
	cacheVaryCookies string
}

/**
 * Constructor for OutputPage. This should not be called directly.
 * Instead a new RequestContext should be created and it will implicitly create
 * a OutputPage tied to that context.
 * @param IContextSource $context
 */
func NewOutputPage() *OutputPage{
	this := new(OutputPage)
	this.mIsArticleRelated = true
	this.mCategories = map[string][]string{
		"hidden" : []string{},
		"normal": []string{},
	}
	this.mEnableClientCache = true
	this.mCdnMaxageLimit = math.Inf(0)
	this.mPreventClickjacking = true
	this.mIndexPolicy = "index"
	this.mFollowPolicy = "follow"

	this.mVaryHeader =  map[string][]string{
		"Accept-Encoding": {"match=gzip"},
	}

	return this
}

/**
 * Set the page as printable, i.e. it'll be displayed with all
 * print styles included
 */
func (o *OutputPage) SetPrintable() {
	o.mPrintable = true
}

/**
 * Disable output completely, i.e. calling output() will have no effect
 */
func (o *OutputPage) Disable() {
	o.mDoNothing = true
}

/**
 * Return whether the output will be completely disabled
 *
 * @return bool
 */
func (o *OutputPage) IsDisabled() bool {
	return o.mDoNothing
}

/**
 * Show an "add new section" link?
 *
 * @return bool
 */
func (o *OutputPage) ShowNewSectionLink() bool {
	return o.mNewSectionLink
}

/**
 * Forcibly hide the new section link?
 *
 * @return bool
 */
func (o *OutputPage) ForceHideNewSectionLink() bool {
	return o.mHideNewSectionLink
}
