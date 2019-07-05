package exception

import (
	"fmt"
	"errors"
)

/**
 * Exception class which takes an HTML error message, and does not
 * produce a backtrace. Replacement for OutputPage::fatalError().
 *
 * @since 1.7
 * @ingroup Exception
 */
type FatalError struct {
	MWException
}


func NewFatalError(name string) *FatalError {
	this := new(FatalError)
	this.error = errors.New(fmt.Sprintf("undefined option:'%s'", name))
	return this
}


/**
 * @return string
 */
func (f *FatalError) GetHTML() string {
	return f.error.Error()
}

/**
 * @return string
 */
func (f *FatalError) GetText() string {
	return f.error.Error()
}