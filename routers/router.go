package routers

import (
	"github.com/MangoDowner/mediawiki/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
