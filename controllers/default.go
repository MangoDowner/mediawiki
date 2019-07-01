package controllers

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/consts"
)

type MainController struct {
	BaseController
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
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
