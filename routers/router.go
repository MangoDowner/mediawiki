package routers

import (
	"github.com/MangoDowner/mediawiki/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},
		"get,post:Main")
    beego.Router("/", &controllers.MainController{})
	// Send Ajax requests to the Ajax dispatcher.
    beego.Router("/test", &controllers.TestController{})
	beego.Router("/ajax", &controllers.AjaxController{},
	"get,post:Ajax")

}
