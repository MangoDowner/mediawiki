package registration

/**
 * "requires" key that applies to MediaWiki core/$wgVersion
 */
const MEDIAWIKI_CORE = "MediaWiki"

/**
 * Version of the highest supported manifest version
 * Note: Update MANIFEST_VERSION_MW_VERSION when changing this
 */
const MANIFEST_VERSION = 2

/**
 * MediaWiki version constraint representing what the current
 * highest MANIFEST_VERSION is supported in
 */
const MANIFEST_VERSION_MW_VERSION = ">= 1.29.0"

/**
 * Version of the oldest supported manifest version
 */
const OLDEST_MANIFEST_VERSION = 1

/**
 * Bump whenever the registration cache needs resetting
 */
const CACHE_VERSION = 7

/**
 * Special key that defines the merge strategy
 *
 * @since 1.26
 */
const MERGE_STRATEGY = "_merge_strategy";


/**
 * ExtensionRegistry class
 *
 * The Registry loads JSON files, and uses a Processor
 * to extract information from them. It also registers
 * classes with the autoloader.
 *
 * @since 1.25
 */
type ExtensionRegistry struct {
	/**
	 * Array of loaded things, keyed by name, values are credits information
	 *
	 * @var array
	 */
	loaded []string

	/**
	 * List of paths that should be loaded
	 *
	 * @var array
	 */
	queued []string

	/**
	 * Whether we are done loading things
	 *
	 * @var bool
	 */
	finished bool

	/**
	 * Items in the JSON file that aren't being
	 * set as globals
	 *
	 * @var array
	 */
	attributes map[string][]string

	/**
	 * @var ExtensionRegistry
	 */
	instance *ExtensionRegistry
}

func NewExtensionRegistry() *ExtensionRegistry {
	this := new(ExtensionRegistry)
	return this
}

/**
 * @param string $name
 * @return array
 */
func (e *ExtensionRegistry) GetAttribute(name string) (ret []string) {
	if v, ok := e.attributes[name]; ok {
		return v
	}
	return ret
}
