package objectcache
/** Possible values for getLastError() */
const ERR_NONE = 0 // no error
const ERR_NO_RESPONSE = 1 // no response
const ERR_UNREACHABLE = 2 // can't connect
const ERR_UNEXPECTED = 3 // response gave some error

/** Bitfield constants for get()/getMulti() */
const READ_LATEST = 1 // use latest data for replicated stores
const READ_VERIFIED = 2 // promise that caller can tell when keys are stale
/** Bitfield constants for set()/merge() */
const WRITE_SYNC = 1 // synchronously write to all locations for replicated stores
const WRITE_CACHE_ONLY = 2 // Only change state of the in-memory cache

/**
 * Class representing a cache/ephemeral data store
 *
 * This interface is intended to be more or less compatible with the PHP memcached client.
 *
 * Instances of this class should be created with an intended access scope, such as:
 *   - a) A single PHP thread on a server (e.g. stored in a PHP variable)
 *   - b) A single application server (e.g. stored in APC or sqlite)
 *   - c) All application servers in datacenter (e.g. stored in memcached or mysql)
 *   - d) All application servers in all datacenters (e.g. stored via mcrouter or dynomite)
 *
 * Callers should use the proper factory methods that yield BagOStuff instances. Site admins
 * should make sure the configuration for those factory methods matches their access scope.
 * BagOStuff subclasses have widely varying levels of support for replication features.
 *
 * For any given instance, methods like lock(), unlock(), merge(), and set() with WRITE_SYNC
 * should semantically operate over its entire access scope; any nodes/threads in that scope
 * should serialize appropriately when using them. Likewise, a call to get() with READ_LATEST
 * from one node in its access scope should reflect the prior changes of any other node its access
 * scope. Any get() should reflect the changes of any prior set() with WRITE_SYNC.
 *
 * @ingroup Cache
 */
type BagOStuff struct {
	/** @var array[] Lock tracking */
	locks []string
	/** @var int ERR_* class constant */
	lastError int
	/** @var string */
	keyspace string
	/** @var LoggerInterface */
	logger interface{}
	/** @var callback|null */
	asyncHandler func()
	/** @var int Seconds */
	syncTimeout int

	/** @var bool */
	debugMode bool
	/** @var array */
	duplicateKeyLookups map[string]int
	/** @var bool */
	reportDupes bool
	/** @var bool */
	dupeTrackScheduled bool

	/** @var callable[] */
	busyCallbacks []string

	/** @var float|null */
	wallClockOverride float64

	/** @var int[] Map of (ATTR_* class constant => QOS_* class constant) */
	attrMap map[int]int

}

func NewBagOStuff() *BagOStuff {
	this := new(BagOStuff)
	this.lastError = ERR_NONE
	this.keyspace = "local"
	return this
}

/**
 * Get an item with the given key
 *
 * If the key includes a deterministic input hash (e.g. the key can only have
 * the correct value) or complete staleness checks are handled by the caller
 * (e.g. nothing relies on the TTL), then the READ_VERIFIED flag should be set.
 * This lets tiered backends know they can safely upgrade a cached value to
 * higher tiers using standard TTLs.
 *
 * @param string $key
 * @param int $flags Bitfield of BagOStuff::READ_* constants [optional]
 * @param int|null $oldFlags [unused]
 * @return mixed Returns false on failure and if the item does not exist
 */
func (h *HashBagOStuff) Get(key string, flags, oldFlags int) bool {
	// B/C for ( $key, &$casToken = null, $flags = 0 )
	if oldFlags != 0 {
		flags = oldFlags
	}

	h.trackDuplicateKeys(key)
	// TODO
	return h.doGet(key, flags)
}

/**
 * @param string $key
 * @param int $flags Bitfield of BagOStuff::READ_* constants [optional]
 * @return mixed Returns false on failure and if the item does not exist
 */
func (h *HashBagOStuff) doGet(key string, flags int) bool {
	return true
}


/**
 * Track the number of times that a given key has been used.
 * @param string $key
 */
func (h *HashBagOStuff) trackDuplicateKeys(key string) {
	if !h.reportDupes {
		return
	}
	if _, ok := h.duplicateKeyLookups[key]; !ok {
		// Track that we have seen this key. This N-1 counting style allows
		// easy filtering with array_filter() later.
		h.duplicateKeyLookups[key] = 0
		return
	}

 	h.duplicateKeyLookups[key] += 1
 	if h.dupeTrackScheduled {
 		return
	}

 	h.dupeTrackScheduled = true
	// Schedule a callback that logs keys processed more than once by get().
	// TODO 去除重复key的错误打印日志
}