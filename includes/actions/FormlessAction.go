/**
 * Base classes for actions done on pages.
 */
package actions
/**
 * An action which just does something, without showing a form first.
 *
 * @ingroup Actions
 */
type FormlessAction struct {
	Action
}

/**
 * Show something on GET request.
 * @return string|null Will be added to the HTMLForm if present, or just added to the
 *     output if not.  Return null to not add anything
 */
func NewFormlessAction() *FormlessAction {
	this := new(FormlessAction)
	return this
}

func (f *FormlessAction) OnView() string {
	return ""
}

func (f *FormlessAction) Show() {

}


