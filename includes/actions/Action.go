/**
 * @defgroup Actions Action done on pages
 */
package actions

import (
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/astaxie/beego"
)

/**
 * Actions are things which can be done to pages (edit, delete, rollback, etc).  They
 * are distinct from Special Pages because an action must apply to exactly one page.
 *
 * To add an action in an extension, create a subclass of Action, and add the key to
 * $wgActions.
 *
 * Actions generally fall into two groups: the show-a-form-then-do-something-with-the-input
 * format (protect, delete, move, etc), and the just-do-something format (watch, rollback,
 * patrol, etc). The FormAction and FormlessAction classes represent these two groups.
 */
type Action struct {

}

func NewAction() *Action {
	this := new(Action)
	return this
}

/**
 * Get the action that will be executed, not necessarily the one passed
 * passed through the "action" request parameter. Actions disabled in
 * $wgActions will be replaced by "nosuchaction".
 *
 * @since 1.19
 * @param IContextSource $context
 * @return string Action name
 */
func (a *Action) GetActionName(c beego.Controller) (actionName string) {
	actionName = c.GetString("action", "view")

	// Check for disabled actions
	if _, ok := includes.WgActions[actionName]; ok && includes.WgActions[actionName] == "" {
		actionName = "nosuchaction"
	}

	// Workaround for T22966: inability of IE to provide an action dependent
	// on which submit button is clicked.
	if actionName == "historysubmit" {
		if rd, _ := c.GetBool("revisiondelete"); rd {
			actionName = "revisiondelete"
		} else if rd, _ := c.GetBool("editchangetags"); rd {
			actionName = "editchangetags"
		} else {
			actionName = "view"
		}
	} else if actionName == "editredlink" {
		actionName = "edit"
	}

	// Trying to get a WikiPage for NS_SPECIAL etc. will result
	// in WikiPage::factory throwing "Invalid or virtual namespace -1 given."
	// For SpecialPages et al, default to action=view.
	// TODO:

	return actionName
	return "nosuchaction"
}

/**
 * Set output headers for noindexing etc.  This function will not be called through
 * the execute() entry point, so only put UI-related stuff in here.
 * @since 1.17
 */
func (a *Action) GetOutput(c beego.Controller) string {
	//TODO
	return ""
}

/**
 * Set output headers for noindexing etc.  This function will not be called through
 * the execute() entry point, so only put UI-related stuff in here.
 * @since 1.17
 */
func (a *Action) SetHeaders(c beego.Controller) {
}