/**
 * Per-process memory cache for storing items.
 */
package objectcache

import (
	"github.com/MangoDowner/mediawiki/includes/consts"
)
const KEY_VAL = 0
const KEY_EXP = 1

/**
 * Simple store for keeping values in an associative array for the current process.
 *
 * Data will not persist and is not shared with other processes.
 *
 * @ingroup Cache
 */
type HashBagOStuff struct {
	BagOStuff
	/** @var mixed[] */
	bag map[string]interface{}
	/** @var int Max entries allowed */
	maxCacheKeys int
}

/**
 * @param array $params Additional parameters include:
 *   - maxKeys : only allow this many keys (using oldest-first eviction)
 */
func NewHashBagOStuff(params map[string]interface{}) *HashBagOStuff {
	this := new(HashBagOStuff)
	if value, ok := params["maxKeys"]; ok {
		this.maxCacheKeys = value.(int)
	} else {
		this.maxCacheKeys = consts.INF
	}

	if this.maxCacheKeys <= 0 {
		panic("$maxKeys parameter must be above zero")
	}
	return this
}

func (l *HashBagOStuff) expire(key string) bool {
	et := (l.bag[key]).(map[int]int)[KEY_EXP]
	// TODO
	if et == TTL_INDEFINITE  {
		return false
	}
	return true
}

func (l *HashBagOStuff) Set(key string, value interface{}, exptime, flags int)  {

}

/**
 * @return float UNIX timestamp
 * @codeCoverageIgnore
 */
func (l *HashBagOStuff) getCurrentTime()  {
}