package exception

import (
	"errors"
)

/**
 * MediaWiki exception
 *
 * @ingroup Exception
 */
 type MWException struct {
 	error
 }



func NewMWException(msg string) *MWException {
	this := new(MWException)
	this.error = errors.New(msg)
	return this
}