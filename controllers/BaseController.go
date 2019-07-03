/**
 * Helper class for the index.php entry point.
 * 改编自includes/MediaWiki.php
 */
package controllers

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/actions"
	"github.com/MangoDowner/mediawiki/includes/specialpage"
	"github.com/astaxie/beego"
)

type BaseController struct {
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
func (b *BaseController) parseTitle() (ret *includes.Title, err error) {
	curid := b.GetString("curid")
	title := b.GetString("title")
	action := b.GetString("action")
	//TODO
	if b.RequestContext.GetRequest().GetCheck("search") {
		return ret, nil
	}

	fmt.Println("[parseTitle]", curid, title, action)

	return ret, nil
}


/**
 * Get the Title object that we'll be acting on, as specified in the WebRequest
 * @return Title
 */
func (b *BaseController) GetTitle() *includes.Title {
	if !b.HasTitle() {
		title, err := b.parseTitle()
		if err != nil {
			title = specialpage.NewSpecialPage().GetTitleFor("Badtitle", "", "")
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
func (b *BaseController) GetAction() string {
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
func (b *BaseController) performRequest() {
	//request := b.Ctx.Request
	//title := b.GetTitle()
	//requestTitle := title
	output := b.GetOutput()
	//user := b.GetUser()

	if b.GetString("printable") == "yes" {
		output.SetPrintable()
	}
	//unused := nil // To pass it by reference
	includes.NewHooks().Run("BeforeInitialize", []includes.MediaWikiServices{}, "")
}























