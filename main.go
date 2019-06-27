package main

import (
	"github.com/MangoDowner/mediawiki/includes"
	_ "github.com/MangoDowner/mediawiki/routers"
	"github.com/astaxie/beego"
)

func main() {
	includes.WfEntryPointCheck()
	beego.Run()
}

