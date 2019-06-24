package config

import (
	"github.com/MangoDowner/mediawiki/globals"
	"github.com/MangoDowner/mediawiki/includes/exception"
)

/**
 * Accesses configuration settings from $GLOBALS
 *
 * @since 1.23
 */
 type GlobalVarConfig struct {
	prefix string
 }

/**
* Default builder function
* @return GlobalVarConfig
*/
func NewInstance() *GlobalVarConfig {
	return NewGlobalVarConfig("")
}

 func NewGlobalVarConfig(prefix string) *GlobalVarConfig {
 	this := new(GlobalVarConfig)
 	if prefix == "" {
 		prefix = "wg"
	}
	this.prefix = prefix
	return this
 }


/**
 * @inheritDoc
 */
func (g *GlobalVarConfig) Get(name string) interface{} {
	if  !g.Has(name) {
		//TODO: PHP 为抛出异常
		panic(exception.NewConfigException(name))
	}
	return g.getWithPrefix(g.prefix, name)
}

/**
 * @inheritDoc
 */
func (g *GlobalVarConfig) Has(name string) bool {
	return g.hasWithPrefix(g.prefix, name)
}

/**
 * Get a variable with a given prefix, if not the defaults.
 *
 * @param string $prefix Prefix to use on the variable, if one.
 * @param string $name Variable name without prefix
 * @return mixed
 */
func (g *GlobalVarConfig) getWithPrefix(prefix, name string) interface{} {
	return globals.GLOBALS[prefix + name]
}

/**
 * Check if a variable with a given prefix is set
 *
 * @param string $prefix Prefix to use on the variable
 * @param string $name Variable name without prefix
 * @return bool
 */
func (g *GlobalVarConfig) hasWithPrefix(prefix, name string) bool {
	foo := prefix + name
	_, ok := globals.GLOBALS[foo]
	return ok
}