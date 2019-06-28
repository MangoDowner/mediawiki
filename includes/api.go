package includes

import "github.com/MangoDowner/mediawiki/includes/consts"

// Set a dummy $wgTitle, because $wgTitle == null breaks various things
// In a perfect world this wouldn't be necessary
var WgTitle *Title

func init() {
	WgTitle = NewTitle().MakeTitle(consts.NS_SPECIAL,
		"Badtitle/dummy title for API calls set in api.php", "", "")
}