package exception

import (
	"errors"
	"fmt"
)

/**
 * Exceptions for config failures
 *
 * @since 1.23
 */
type ConfigException struct {
	err error
}


func NewConfigException(name string) *ConfigException {
	this := new(ConfigException)
	this.err = errors.New(fmt.Sprintf("undefined option:'%s'", name))
	return this
}