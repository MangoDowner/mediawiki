package exception

import (
	"fmt"
	"errors"
)

/**
 * Exception thrown when a service was already defined, but the
 * caller expected it to not exist.
 */
type ServiceAlreadyDefinedException struct {
	err error
}

/**
 * @param string $serviceName
 * @param Exception|null $previous
 */
func NewServiceAlreadyDefinedException(serviceName string, previous error) *NoSuchServiceException {
	this := new(NoSuchServiceException)
	this.err = errors.New(fmt.Sprintf("Service already defined <%s> : %s", serviceName, previous))
	return this
}
