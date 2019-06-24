/**
 * Base class for resource loading system.
 */
package resourceloader

import "github.com/MangoDowner/mediawiki/includes/config"

/**
 * Dynamic JavaScript and CSS resource loading system.
 *
 * Most of the documentation is on the MediaWiki documentation wiki starting at:
 *    https://www.mediawiki.org/wiki/ResourceLoader
 * TODO: 大部分变量类型存疑
 */
/** @var int */
const CACHE_VERSION = 8
type ResourceLoader struct {

	/** @var bool */
	debugMode bool

	/**
	 * Module name/ResourceLoaderModule object pairs
	 * @var array
	 */
	modules []string

	/**
	 * Associative array mapping module name to info associative array
	 * @var array
	 */
	moduleInfos []string

	/** @var Config $config */
	config config.Config


	/**
	 * Associative array mapping framework ids to a list of names of test suite modules
	 * like [ 'qunit' => [ 'mediawiki.tests.qunit.suites', 'ext.foo.tests', ... ], ... ]
	 * @var array
	 */
	testModuleNames []string

	/**
	 * E.g. [ 'source-id' => 'http://.../load.php' ]
	 * @var array
	 */
	sources []string

	/**
	 * Errors accumulated during current respond() call.
	 * @var array
	 */
	errors []string

	/**
	 * List of extra HTTP response headers provided by loaded modules.
	 *
	 * Populated by makeModuleResponse().
	 *
	 * @var array
	 */
	extraHeaders []string

	/**
	 * @var MessageBlobStore
	 */
	blobStore string

	/**
	 * @var LoggerInterface
	 */
	logger string
}

/**
 * @return Config
 */
func (r *ResourceLoader) GetConfig() config.Config {
	return r.config
}