/**
 * Helper class for the index.php entry point.
 * 改编自includes/MediaWiki.php
 */
package controllers

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/consts"
)

type MainController struct {
	MediaWiki
}

func (c *MainController) Main() {
	// Get Send Ajax requests to the Ajax dispatcher.
	if c.GetString("action") == "ajax" {
		// Set a dummy title, because $wgTitle == null might break things
		title := includes.NewTitle().MakeTitle(consts.NS_SPECIAL,
			fmt.Sprintf("Badtitle/performing an AJAX call in __METHOD__"), "", "")
		c.SetTitle(title)
		dispatcher := NewAjaxDispatcher(&c.Controller, nil)
		dispatcher.performAction(nil)
		return
	}

	// Get title from request parameters,
	// is set on the fly by parseTitle the first time.
	title := c.GetTitle()
	action := c.GetAction()
	fmt.Println("TITLE:", title)
	fmt.Println("ACTION:", action)
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
