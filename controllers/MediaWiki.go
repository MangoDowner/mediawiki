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
	"github.com/astaxie/beego"
)

type MediaWiki struct {
	/**
	 * @var IContextSource
	 */
	beego.Controller
	includes.RequestContext

	/**
	 * @var String Cache what action this request is
	 */
	action string
}

/**
 * Parse the request to get the Title object
 *
 * @throws MalformedTitleException If a title has been provided by the user, but is invalid.
 * @return Title Title object to be $wgTitle
 */
func (b *MediaWiki) parseTitle() (ret *includes.Title, err error) {
	curId := b.GetString("curid")
	title := b.GetString("title")
	action := b.GetString("action")
	//TODO
	if b.RequestContext.GetRequest().GetCheck("search") {
		return ret, nil
	}

	fmt.Println("[parseTitle]", curId, title, action)

	return ret, nil
}


/**
 * Get the Title object that we'll be acting on, as specified in the WebRequest
 * @return Title
 */
func (b *MediaWiki) GetTitle() *includes.Title {
	if !b.HasTitle() {
		title, err := b.parseTitle()
		if err != nil {
			title = includes.NewSpecialPage().GetTitleFor("Badtitle", "", "")
		}
		b.SetTitle(title)
	}
	return b.RequestContext.GetTitle()
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
	//requestTitle := title
	output := b.GetOutput()
	//user := b.GetUser()

	if b.GetString("printable") == "yes" {
		output.SetPrintable()
	}
	var unused bool // To pass it by reference
	includes.NewHooks().Run("BeforeInitialize", []interface{}{&title, &unused}, "")

	// Invalid titles. T23776: The interwikis must redirect even if the page name is empty.
	fmt.Println("INVALID")
	fmt.Println("title:", title)
	fmt.Println("title.GetDBKey():", title.GetDBKey())
	fmt.Println("title.IsExternal():", title.IsExternal())
	fmt.Println(`title.IsSpecial("Badtitle"):`, title.IsSpecial("Badtitle"))

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
		fmt.Println("FAC:", spFactory)
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
func (b *MediaWiki) WasPosted() bool {
	return b.GetMethod() == "POST"
}


















