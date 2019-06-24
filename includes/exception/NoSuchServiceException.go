package exception

import (
	"errors"
	"fmt"
)

/**
 * Exception thrown when the requested service is not known.
 */
type NoSuchServiceException struct {
	err error
}

/**
 * @param string $serviceName
 * @param Exception|null $previous
 */
func NewNoSuchServiceException(serviceName string, previous error) *NoSuchServiceException {
	this := new(NoSuchServiceException)
	this.err = errors.New(fmt.Sprintf("No such service <%s> : %s", serviceName, previous))
	return this
}