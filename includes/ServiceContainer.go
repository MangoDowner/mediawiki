/**
 * Generic service container.
 */
package includes

import (
	"github.com/MangoDowner/mediawiki/includes/exception"
)

/**
* ServiceContainer provides a generic service to manage named services using
* lazy instantiation based on instantiator callback functions.
*
* Services managed by an instance of ServiceContainer may or may not implement
* a common interface.
*
* @note When using ServiceContainer to manage a set of services, consider
* creating a wrapper or a subclass that provides access to the services via
* getter methods with more meaningful names and more specific return type
* declarations.
*
* @see docs/injection.txt for an overview of using dependency injection in the
*      MediaWiki code base.
 */ /**
 * Destroys all contained service instances that implement the DestructibleService
 * interface. This will render all services obtained from this MediaWikiServices
 * instance unusable. In particular, this will disable access to the storage backend
 * via any of these services. Any future call to getService() will throw an exception.
 *
 * @see resetGlobalInstance()
 */
type ServiceContainer struct {
	/**
	 * @var object[]
	 */
	services map[string]interface{}

	/**
	 * @var callable[]
	 */
	serviceInstantiators map[string]func()

	/**
	 * @var callable[][]
	 */
	serviceManipulators [][]string

	/**
	 * @var bool[] disabled status, per service name
	 */
	disabled map[string]bool

	/**
	 * @var array
	 */
	extraInstantiationParams []string

	/**
	 * @var bool
	 */
	destroyed bool
}

/**
 * @param array $extraInstantiationParams Any additional parameters to be passed to the
 * instantiator function when creating a service. This is typically used to provide
 * access to additional ServiceContainers or Config objects.
 */
func NewServiceContainer(extraInstantiationParams []string) *ServiceContainer {
	this := new(ServiceContainer)
	this.extraInstantiationParams = extraInstantiationParams
	return this
}

/**
* Destroys all contained service instances that implement the DestructibleService
* interface. This will render all services obtained from this MediaWikiServices
* instance unusable. In particular, this will disable access to the storage backend
* via any of these services. Any future call to getService() will throw an exception.
*
* @see resetGlobalInstance()
 */
func (s *ServiceContainer) Destroy() {
	for _, name := range s.GetServiceNames() {
		service := s.PeekService(name)
		if service == nil {
			continue
		}
		if i, ok := service.(DestructibleService); ok {
			i.Destroy()
		}
	}
	// Break circular references due to the $this reference in closures, by
	// erasing the instantiator array. This allows the ServiceContainer to
	// be deleted when it goes out of scope.
	s.serviceInstantiators = nil
	s.services = nil
	s.destroyed = true
}

/**
* Returns true if a service is defined for $name, that is, if a call to getService( $name )
* would return a service instance.
*
* @param string $name
*
* @return bool
 */
func (s *ServiceContainer) HasService(name string) bool {
	_, ok := s.serviceInstantiators[name]
	return ok
}

/**
 * Returns the service instance for $name only if that service has already been instantiated.
 * This is intended for situations where services get destroyed/cleaned up, so we can
 * avoid creating a service just to destroy it again.
 *
 * @note This is intended for internal use and for test fixtures.
 * Application logic should use getService() instead.
 *
 * @see getService().
 *
 * @param string $name
 *
 * @return object|null The service instance, or null if the service has not yet been instantiated.
 * @throws RuntimeException if $name does not refer to a known service.
 */
func (s *ServiceContainer) PeekService(name string) interface{} {
	if !s.HasService(name) {
		//TODO: PHP 为抛出异常
		panic(exception.NewNoSuchServiceException(name, nil))
	}
	return s.services[name]
}

/**
* @return string[]
 */
func (s *ServiceContainer) GetServiceNames() (result []string) {
	for k := range s.serviceInstantiators {
		result = append(result, k)
	}
	return result
}

/**
 * Define a new service. The service must not be known already.
 *
 * @see getService().
 * @see redefineService().
 *
 * @param string $name The name of the service to register, for use with getService().
 * @param callable $instantiator Callback that returns a service instance.
 *        Will be called with this MediaWikiServices instance as the only parameter.
 *        Any extra instantiation parameters provided to the constructor will be
 *        passed as subsequent parameters when invoking the instantiator.
 *
 * @throws RuntimeException if there is already a service registered as $name.
 */
func (s *ServiceContainer) DefineService(name string, instantiator func()) {
	// Assert::parameterType( 'string', $name, '$name' );
	if s.HasService(name) {
		panic(exception.NewServiceAlreadyDefinedException(name, nil))
	}
	//ServiceAlreadyDefinedException( $name );
	if s.serviceInstantiators == nil {
		s.serviceInstantiators = make(map[string]func())
	}
	s.serviceInstantiators[name] = instantiator
}

/**
 * Returns a service object of the kind associated with $name.
 * Services instances are instantiated lazily, on demand.
 * This method may or may not return the same service instance
 * when called multiple times with the same $name.
 *
 * @note Rather than calling this method directly, it is recommended to provide
 * getters with more meaningful names and more specific return types, using
 * a subclass or wrapper.
 *
 * @see redefineService().
 *
 * @param string $name The service name
 *
 * @throws NoSuchServiceException if $name is not a known service.
 * @throws ContainerDisabledException if this container has already been destroyed.
 * @throws ServiceDisabledException if the requested service has been disabled.
 *
 * @return object The service instance
 */
func (s *ServiceContainer) GetService(name string) (ret *SpecialPageFactory) {
	//if s.destroyed {
	//	panic(exception.NewNoSuchServiceException(name, nil))
	//}
	//if _, ok :=
	//if ( isset( $this->disabled[$name] ) ) {
	//throw new ServiceDisabledException( $name );
	//}
	//
	//if ( !isset( $this->services[$name] ) ) {
	//$this->services[$name] = $this->createService( $name );
	//}
	//
	//return $this->services[$name];
	return ret
}
