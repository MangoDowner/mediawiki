/**
 * Helper class for the index.php entry point.
 * 改编自includes/MediaWiki.php
 */
package controllers

import (
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/MangoDowner/mediawiki/includes/actions"
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
 * Get the Title object that we'll be acting on, as specified in the WebRequest
 * @return Title
 */
func (b *BaseController) GetTitle() string {
	// TODO
	title := b.GetString("title")
	if b.GetString("title") == "" {

	}
	return title
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

