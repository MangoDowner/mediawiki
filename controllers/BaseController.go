package controllers

import (
	"github.com/MangoDowner/mediawiki/includes"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	includes.RequestContext
}


