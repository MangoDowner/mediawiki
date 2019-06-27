/**
 * Localisation messages cache.
 */
package cache

import "sync"

/**
 * MediaWiki message cache structure version.
 * Bump this whenever the message cache format has changed.
 */
const MSG_CACHE_VERSION = 2

var (
	ins  *MessageCache
	once sync.Once
)

type MessageCache struct {
}

func NewMessageCache() *MessageCache {
	this := new(MessageCache)
	return this
}

/**
 * Get the signleton instance of this class
 *
 * @since 1.18
 * @return MessageCache
 */
func SingletonMessageCache() *MessageCache {
	once.Do(func() {
		ins = new(MessageCache)
	})
	return ins
}
