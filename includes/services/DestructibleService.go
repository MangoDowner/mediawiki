/**
 * Interface for destructible services.
 *	@since 1.27
 */
package includes

/**
 * DestructibleService defines a standard interface for shutting down a service instance.
 * The intended use is for a service container to be able to shut down services that should
 * no longer be used, and allow such services to release any system resources.
 *
 * @note There is no expectation that services will be destroyed when the process (or web request)
 * terminates.
 */
type DestructibleService interface {
	/**
	 * Notifies the service object that it should expect to no longer be used, and should release
	 * any system resources it may own. The behavior of all service methods becomes undefined after
	 * destroy() has been called. It is recommended that implementing classes should throw an
	 * exception when service methods are accessed after destroy() has been called.
	 */
	Destroy()
}