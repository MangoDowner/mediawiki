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
 	err error
 }



func NewMWException(msg string) *MWException {
	this := new(MWException)
	this.err = errors.New(msg)
	return this
}