package main

import (
	"github.com/MangoDowner/mediawiki/includes"
	_ "github.com/MangoDowner/mediawiki/routers"
)

func main() {
	includes.WfEntryPointCheck()
	mediaWiki := includes.NewMediaWiki(nil)
	mediaWiki.Run()
}

