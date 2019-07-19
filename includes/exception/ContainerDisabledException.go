package exception

import (
	"fmt"
	"errors"
)

/**
 * Exception thrown when a service was already defined, but the
 * caller expected it to not exist.
 */
type ContainerDisabledException struct {
	err error
}

/**
 * @param string $serviceName
 * @param Exception|null $previous
 */
func NewContainerDisabledException(previous error) *ContainerDisabledException {
	this := new(ContainerDisabledException)
	this.err = errors.New(fmt.Sprintf("Container disabled! : %s", previous))
	return this
}
