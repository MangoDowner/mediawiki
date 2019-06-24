/**
 * Context for ResourceLoader modules.
 */
package resourceloader

import (
	"github.com/MangoDowner/mediawiki/includes"
	"strings"
)

/**
 * Object passed around to modules which contains information about the state
 * of a specific loader request.
 */
type ResourceLoaderContext struct {
	resourceLoader ResourceLoader
	request        includes.WebRequest
	logger         string

	// Module content vary
	skin     string
	language string
	debug    bool
	user     string

	// Request vary (in addition to cache vary)
	modules []string
	only    string
	version string
	raw     bool
	image   string
	variant string
	format  string

	direction string
	hash      string
	userObj   string
	imageObj  string
}

/**
 * @param ResourceLoader $resourceLoader
 * @param WebRequest $request
 */
func NewResourceLoaderContext(resourceLoader ResourceLoader, request includes.WebRequest) *ResourceLoaderContext {
	this := new(ResourceLoaderContext)
	this.resourceLoader = resourceLoader
	this.request = request
	//$this->logger = $resourceLoader->getLogger();

	// Future developers: Use WebRequest::getRawVal() instead getVal().
	// The getVal() method performs slow Language+UTF logic. (f303bb9360)

	// List of modules
	modulesRaw := request.GetRawVal("modules", "")
	modules := modulesRaw.(string)
	if modules != "" {
		this.modules = expandModuleNames(modules)
	} else {
		this.modules = []string{}
	}

	// Various parameters
	this.user = request.GetRawVal("user", nil).(string)
	this.debug = request.GetFuzzyBool("debug",
		resourceLoader.GetConfig().Get("ResourceLoaderDebug").(bool))
	this.only = request.GetRawVal("only", nil).(string)
	this.version = request.GetRawVal( "version", nil ).(string)
	this.raw = request.GetFuzzyBool( "raw", false)

	// Image requests

	this.image = request.GetRawVal( "image", nil ).(string)
	this.variant = request.GetRawVal( "variant", nil ).(string)
	this.format = request.GetRawVal( "format", nil).(string)

	this.skin = request.GetRawVal( "skin", nil ).(string)
	//skinnames := Skin::getSkinNames();
	var skinnames map[string] string
	// If no skin is specified, or we don't recognize the skin, use the default skin
	if this.skin != "" {
		return this
	}
	if _, ok := skinnames[this.skin]; !ok {
		this.skin = resourceLoader.GetConfig().Get("efaultSkin").(string)
	}
	return this
}

/**
 * Expand a string of the form `jquery.foo,bar|jquery.ui.baz,quux` to
 * an array of module names like `[ 'jquery.foo', 'jquery.bar',
 * 'jquery.ui.baz', 'jquery.ui.quux' ]`.
 *
 * This process is reversed by ResourceLoader::makePackedModulesString().
 *
 * @param string $modules Packed module name list
 * @return array Array of module names
 */
func expandModuleNames(modules string) ( retVal[]string) {
	exploded := strings.Split(modules, "|")
	for _, group := range exploded {
		if !strings.ContainsRune(group, ',') {
			// This is not a set of modules in foo.bar,baz notation
			// but a single module
			retVal = append(retVal, group)
		} else {
			// This is a set of modules in foo.bar,baz notation
			pos := strings.LastIndex(group, ".")
			if pos < 0 {
				// Prefixless modules, i.e. without dots
				groupArr := strings.Split(group, ",")
				for _, v := range groupArr {
					retVal = append(retVal, v)
				}
			} else {
				// We have a prefix and a bunch of suffixes
				prefix := group[0 : pos] // 'foo'
				suffix := strings.Split(group[pos + 1 : ], ",") // [ 'bar', 'baz' ]
				for _, suffixes := range suffix {
					retVal = append(retVal, prefix + "." + suffixes)
				}
			}
		}
	}
	return retVal
}

/**
 * @return bool
 */
func (r *ResourceLoaderContext) GetDebug() bool {
	return r.debug
}

/**
 * @return string|null
 */
func (r *ResourceLoaderContext) GetOnly() string {
	return r.only
}

/**
 * @see ResourceLoaderModule::getVersionHash
 * @see ResourceLoaderClientHtml::makeLoad
 * @return string|null
 */
func (r *ResourceLoaderContext) GetVersion() string {
	return r.version
}

/**
 * @return bool
 */
func (r *ResourceLoaderContext) GetRaw() bool {
	return r.raw
}

/**
 * @return string|null
 */
func (r *ResourceLoaderContext) GetImage() string {
	return r.image
}

/**
 * @return string|null
 */
func (r *ResourceLoaderContext) GetVariant() string {
	return r.variant
}

/**
 * @return string|null
 */
func (r *ResourceLoaderContext) GetFormat() string {
	return r.format
}
