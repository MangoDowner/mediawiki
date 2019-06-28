/**
 * Helper class for the index.php entry point.
 */
package includes

import (
	"github.com/astaxie/beego"
)

/**
 * The MediaWiki class is the helper class for the index.php entry point.
 */
type MediaWiki struct {
	/**
	 * @var IContextSource
	 */
	context IContextSource

	/**
	 * @var Config
	 */
	config interface{}

	/**
	 * @var String Cache what action this request is
	 */
	action string
}

func NewMediaWiki(cs IContextSource) *MediaWiki {
	this := new(MediaWiki)
	if cs == nil {
		cs = NewRequestContext().GetMain()
	}
	this.context = cs
	// TODO
	this.config = nil
	return this
}


/**
 * Run the current MediaWiki instance; index.php just calls this
 */
func (m *MediaWiki) Run() {
	err := m.main()
	if err != nil {

	}
	beego.Run()
}

func (m *MediaWiki) main() error {
	//TODO: 路由处理统一放在routers
	//requests := m.context.GetRequest()
	//// Send Ajax requests to the Ajax dispatcher.
	//if requests.GetVal("action", "") == "ajax" {
	//
	//}
	return nil
}