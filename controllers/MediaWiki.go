/**
 * Helper class for the index.php entry point.
 * 改编自includes/MediaWiki.php
 */
package controllers

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/actions"
	"github.com/MangoDowner/mediawiki/includes/config"
	"github.com/MangoDowner/mediawiki/includes/consts"
	"github.com/astaxie/beego"
)

type MediaWiki struct {
	beego.Controller

	context *includes.RequestContext
	/**
	   * @var IContextSource
	*/

	/**
	 * @var String Cache what action this request is
	 */
	action string
}

// Prepare 这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行.
// 用户可以重写这个函数实现类似用户验证之类。
func (c *MediaWiki) Prepare() {
	context := includes.NewRequestContext(c.Controller.Ctx)
	c.context = context
}

func (c *MediaWiki) Entry() {
	// Get Send Ajax requests to the Ajax dispatcher.
	if c.GetString("action") == "ajax" {
		// Set a dummy title, because $wgTitle == null might break things
		title := includes.NewTitle().MakeTitle(consts.NS_SPECIAL,
			fmt.Sprintf("Badtitle/performing an AJAX call in __METHOD__"), "", "")
		c.context.SetTitle(title)
		dispatcher := NewAjaxDispatcher(&c.Controller, nil)
		dispatcher.performAction(nil)
		return
	}

	// Get title from request parameters,
	// is set on the fly by parseTitle the first time.
	title := c.GetTitle()
	//action := c.GetAction()
	//TODO:

	// If the user has forceHTTPS set to true, or if the user
	// is in a group requiring HTTPS, or if they have the HTTPS
	// preference set, redirect them to HTTPS.
	// Note: Do this after $wgTitle is setup, otherwise the hooks run from
	// isLoggedIn() will do all sorts of weird stuff.
	if false {
		// TODO 转向HTTPS
		return
	}

	// TODO 缓存
	if title.CanExist() && false {

	}

	// Actually do the work of the request and build up any output
	c.performRequest()

	// Now commit any transactions, so that unreported errors after
	// output() don't roll back the whole DB transaction and so that
	// we avoid having both success and error text in the response
	// TODO:
	//$this->doPreOutputCommit( $outputWork );

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


/**
 * Parse the request to get the Title object
 *
 * @throws MalformedTitleException If a title has been provided by the user, but is invalid.
 * @return Title Title object to be $wgTitle
 */
func (b *MediaWiki) parseTitle() (ret *includes.Title, err error) {
	request := b.context.GetRequest();

	curId := request.GetInt("curid", 0)
	// 获取/url.../:title
	title := request.GetVal(":title", "")
	//action := request.GetVal("action", "")
	//TODO
	if b.context.GetRequest().GetCheck("search") {
		return ret, nil
	} else if curId != 0 {

	} else {
		ret = includes.NewTitle().NewFromURL(title)
	}
	return ret, nil
}


/**
 * Get the Title object that we'll be acting on, as specified in the WebRequest
 * @return Title
 */
func (b *MediaWiki) GetTitle() *includes.Title {
	if !b.context.HasTitle() {
		title, err := b.parseTitle()
		if err != nil {
			title = includes.NewSpecialPage().GetTitleFor("Badtitle", "", "")
		}
		b.context.SetTitle(title)
	}
	return b.context.GetTitle()
}

/**
 * Returns the name of the action that will be executed.
 *
 * @return string Action
 */
func (b *MediaWiki) GetAction() string {
	// TODO
	if b.action == "" {
		b.action = actions.NewAction().GetActionName(b.Controller)
	}
	return b.action
}

/**
 * Performs the request.
 * - bad titles
 * - read restriction
 * - local interwiki redirects
 * - redirect loop
 * - special pages
 * - normal pages
 *
 * @throws MWException|PermissionsError|BadTitleError|HttpError
 * @return void
 */
func (b *MediaWiki) performRequest() {
	//request := b.Ctx.Request
	title := b.GetTitle()
	fmt.Println("performRequest title:", title)

	//requestTitle := title
	output := b.context.GetOutput()
	//user := b.GetUser()

	if b.GetString("printable") == "yes" {
		output.SetPrintable()
	}
	var unused bool // To pass it by reference
	includes.NewHooks().Run("BeforeInitialize", []interface{}{&title, &unused}, "")

	// Invalid titles. T23776: The interwikis must redirect even if the page name is empty.
	//fmt.Println("title.GetDBKey():", title.GetDBKey())
	//fmt.Println("title.IsExternal():", title.IsExternal())
	//fmt.Println(`title.IsSpecial("Badtitle"):`, title.IsSpecial("Badtitle"))

	if title == nil ||
		(title.GetDBKey() == "" && !title.IsExternal()) ||
		title.IsSpecial("Badtitle") {
		// TODO
	}

	// Check user's permissions to read this page.
	// We have to check here to catch special pages etc.
	// We will check again in Article::view().
	// TODO: 权限控制

	// Interwiki redirects
	if title.IsExternal() {
		//TODO
		return

	}
	var spFactory *includes.SpecialPageFactory
	// Handle any other redirects.
	// Redirect loops, titleless URL, $wgUsePathInfo URLs, and URLs with a variant
	if !b.tryNormaliseRedirect(title) {
		// Prevent information leak via Special:MyPage et al (T109724)
		spFactory = includes.NewMediaWikiServices().GetSpecialPageFactory()
	}

	// Special pages ($title may have changed since if statement above)
	if title.IsSpecialPage() {
		// Actions that need to be made when we have a special pages
		spFactory.ExecutePath(title, b.Ctx, false, nil)
		return
	}
	// ...otherwise treat it as an article view. The article
	// may still be a wikipage redirect to another article or URL.
	// TODO

}

/**
 * Handle redirects for uncanonical title requests.
 *
 * Handles:
 * - Redirect loops.
 * - No title in URL.
 * - $wgUsePathInfo URLs.
 * - URLs with a variant.
 * - Other non-standard URLs (as long as they have no extra query parameters).
 *
 * Behaviour:
 * - Normalise title values:
 *   /wiki/Foo%20Bar -> /wiki/Foo_Bar
 * - Normalise empty title:
 *   /wiki/ -> /wiki/Main
 *   /w/index.php?title= -> /wiki/Main
 * - Don't redirect anything with query parameters other than 'title' or 'action=view'.
 *
 * @param Title $title
 * @return bool True if a redirect was set.
 * @throws HttpError
 */
func (b *MediaWiki) tryNormaliseRedirect(title *includes.Title) bool {

	if b.GetString("action", "view") != "view" ||
		b.WasPosted() ||
		b.GetString("title") != "" && title.GetPrefixedDBkey() == b.GetString("title") ||
		includes.NewHooks().Run("TestCanonicalRedirect", []interface{}{title}, "") {
		return false
	}

	if title.IsSpecialPage() {
		//TODO
	}

	// Redirect to canonical url, make it a 301 to allow caching
	//targetUrl := title.getFullUrl()
	// TODO
	var message string
	if true {
		fmt.Println("jinlaile")
		message = "Redirect loop detected!\n\n" +
			"This means the wiki got confused about what page was " +
			"requested; this sometimes happens when moving a wiki " +
			"to a new server or changing the server configuration.\n\n"

		if config.Configs.GetBool("UsePathInfo") {
			message = message + "The wiki is trying to interpret the page " +
				"title from the URL path portion (PATH_INFO), which " +
				"sometimes fails depending on the web server. Try " +
				"setting \"$wgUsePathInfo = false;\" in your " +
				"LocalSettings.php, or check that $wgArticlePath " +
				"is correct."
		} else {
			message = message + "Your web server was detected as possibly not " +
				"supporting URL path components (PATH_INFO) correctly; " +
				"check your LocalSettings.php for a customized " +
				"$wgArticlePath setting and/or toggle $wgUsePathInfo " +
				"to true."
		}
		// throw http error
	}
	return true
}


/**
 * Get the HTTP method used for this request.
 *
 * @return string
 */
func (b *MediaWiki) GetMethod() (method string) {
	method = b.Ctx.Input.Method()
	if method == "" {
		method = "GET"
	}
	return method
}

/**
 * Return true if the named value is set in the input, whatever that
 * value is (even "0"). RwasPostedeturn false if the named value is not set.
 * Example use is checking for the presence of check boxes in forms.
 *
 * @param string $name
 * @return bool
 */
func (m *MediaWiki) WasPosted() bool {
	return m.GetMethod() == "POST"
}


















