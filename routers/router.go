package routers

import (
	"github.com/MangoDowner/mediawiki/controllers"
	"github.com/astaxie/beego"
)

func init() {

    beego.Router("/",&controllers.MediaWiki{},
		"get,post:Entry")
    beego.Router("/entry/:title:string", &controllers.MediaWiki{},
		"get,post:Entry")
	// Send Ajax requests to the Ajax dispatcher.
    beego.Router("/test", &controllers.TestController{})
	beego.Router("/ajax", &controllers.AjaxController{},
	"get,post:Ajax")

}
