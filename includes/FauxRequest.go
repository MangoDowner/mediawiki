/**
 * Deal with importing all those nasty globals and things
 */
package includes

import "github.com/MangoDowner/mediawiki/includes/php"

/**
 * WebRequest clone which takes values from a provided array.
 *
 * @ingroup HTTP
 */
type FauxRequest struct {
	wasPosted  bool
	requestUrl string
	cookies    map[string]string
	WebRequest
}

/**
 * @param array $data Array of *non*-urlencoded key => value pairs, the
 *   fake GET/POST values
 * @param bool $wasPosted Whether to treat the data as POST
 * @param MediaWiki\Session\Session|array|null $session Session, session
 *  data array, or null
 * @param string $protocol 'http' or 'https'
 * @throws MWException
 */
func NewFauxRequest(data map[string]string, wasPosted bool,
	session interface{}, protocol string) *FauxRequest {
	if protocol == "" {
		protocol = "http"
	}
	this := new(FauxRequest)
	this.requestTime = php.Microtime(true)
	this.data = data
	this.wasPosted = wasPosted

	//TODO: 缺失大量代码
	return this
}
