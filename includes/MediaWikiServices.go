/**
 * Service locator for MediaWiki core services.
 *
 *  @since 1.27
 */
package includes

import (
	"github.com/MangoDowner/mediawiki/includes/config"
	"github.com/MangoDowner/mediawiki/includes/services"
)

/**
 * MediaWikiServices is the service locator for the application scope of MediaWiki.
 * Its implemented as a simple configurable DI container.
 * MediaWikiServices acts as a top level factory/registry for top level services, and builds
 * the network of service objects that defines MediaWiki's application logic.
 * It acts as an entry point to MediaWiki's dependency injection mechanism.
 *
 * Services are defined in the "wiring" array passed to the constructor,
 * or by calling defineService().
 *
 * @see docs/injection.txt for an overview of using dependency injection in the
 *      MediaWiki code base.
 */

 type MediaWikiServices struct {
	instance interface{}
	includes.ServiceContainer
}

/**
 * Returns the global default instance of the top level service locator.
 *
 * @since 1.27
 *
 * The default instance is initialized using the service instantiator functions
 * defined in ServiceWiring.php.
 *
 * @note This should only be called by static functions! The instance returned here
 * should not be passed around! Objects that need access to a service should have
 * that service injected into the constructor, never a service locator!
 *
 * @return MediaWikiServices
 */
func (m *MediaWikiServices) getInstance() interface{} {
	if m.instance == nil {
		// NOTE: constructing GlobalVarConfig here is not particularly pretty,
		// but some information from the global scope has to be injected here,
		// even if it's just a file name or database credentials to load
		// configuration from.
		bootstrapConfig := config.NewGlobalVarConfig("")
		m.instance = m.newInstance( bootstrapConfig, "load" )
	}
	return m.instance
}

/**
 * Creates a new MediaWikiServices instance and initializes it according to the
 * given $bootstrapConfig. In particular, all wiring files defined in the
 * ServiceWiringFiles setting are loaded, and the MediaWikiServices hook is called.
 *
 * @param Config|null $bootstrapConfig The Config object to be registered as the
 *        'BootstrapConfig' service.
 *
 * @param string $loadWiring set this to 'load' to load the wiring files specified
 *        in the 'ServiceWiringFiles' setting in $bootstrapConfig.
 *
 * @return MediaWikiServices
 * @throws MWException
 * @throws \FatalError
 */
func (m *MediaWikiServices) newInstance( bootstrapConfig config.Config, loadWiring string) *MediaWikiServices {
	instance := NewMediaWikiServices(bootstrapConfig)
	// Load the default wiring from the specified files.
	if loadWiring == "load"  {
		//wiringFiles := bootstrapConfig.Get("ServiceWiringFiles")
		//instance.LoadWiringFiles(wiringFiles)
	}

	// Provide a traditional hook point to allow extensions to configure services.
	NewHooks().Run("MediaWikiServices", []MediaWikiServices{*instance}, "")
	return instance
}

/**
 * @param Config $config The Config object to be registered as the 'BootstrapConfig' service.
 *        This has to contain at least the information needed to set up the 'ConfigFactory'
 *        service.
 */
func NewMediaWikiServices(config config.Config) *MediaWikiServices {
	this := new(MediaWikiServices)
	// Register the given Config object as the bootstrap config service.
	this.DefineService( "BootstrapConfig",
		func () {
			//TODO: 方法参数不一样咋办?
		},
	)
	return this
}