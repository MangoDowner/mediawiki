package controllers

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/consts"
)

type AjaxController struct {
	BaseController
}

// Get Send Ajax requests to the Ajax dispatcher.
func (c *AjaxController) Ajax() {
	// Set a dummy title, because $wgTitle == null might break things
	title := includes.NewTitle().MakeTitle(consts.NS_SPECIAL,
		fmt.Sprintf("Badtitle/performing an AJAX call in __METHOD__"), "", "")
	c.SetTitle(title)
	dispatcher := NewAjaxDispatcher(c.Ctx, nil)
	dispatcher.performAction(nil)
	c.Data["Website"] = c.GetTitle()
	c.TplName = "index.tpl"
	return
}
