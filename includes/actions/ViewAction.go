/**
 * An action that views article content
 */
package actions

/**
 * An action that views article content
 *
 * This is a wrapper that will call Article::view().
 *
 * @ingroup Actions
 */
type ViewAction struct {
	FormlessAction
}

func NewViewAction() *ViewAction {
	this := new(ViewAction)
	return this
}

func (v *ViewAction) GetName() string {
	return "view"
}

func (v *ViewAction) OnView() string {
	return ""
}

func (v *ViewAction) Show() {

}