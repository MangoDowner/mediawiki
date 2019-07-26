package objectcache
/**
 * Default connection timeout in seconds. The kernel retransmits the SYN
 * packet after 1 second, so 1.2 seconds allows for 1 retransmit without
 * permanent failure.
 */
const DEFAULT_CONN_TIMEOUT = 1.2

/**
 * Default request timeout
 */
const DEFAULT_REQ_TIMEOUT = 3.0

/**
 * Interface to key-value storage behind an HTTP server.
 *
 * Uses URL of the form "baseURL/{KEY}" to store, fetch, and delete values.
 *
 * E.g., when base URL is `/v1/sessions/`, then the store would do:
 *
 * `PUT /v1/sessions/12345758`
 *
 * and fetch would do:
 *
 * `GET /v1/sessions/12345758`
 *
 * delete would do:
 *
 * `DELETE /v1/sessions/12345758`
 *
 * Configure with:
 *
 * @code
 * $wgObjectCaches['sessions'] = array(
 *	'class' => 'RESTBagOStuff',
 *	'url' => 'http://localhost:7231/wikimedia.org/v1/sessions/'
 * );
 * @endcode
 */
type RESTBagOStuff struct {
	BagOStuff


	/**
	 * @var MultiHttpClient
	 */
	client interface{}

	/**
	 * REST URL to use for storage.
	 * @var string
	 */
	url string
}

func NewRESTBagOStuff(params map[string]interface{}) *RESTBagOStuff {
	this := new(RESTBagOStuff)
	return this
}

/**
 * @param string $key
 * @param int $flags Bitfield of BagOStuff::READ_* constants [optional]
 * @return mixed Returns false on failure and if the item does not exist
 */
func (h *RESTBagOStuff) doGet(key string, flags int) bool {
	return true
}