package includes

import "github.com/MangoDowner/mediawiki/includes/consts"

var (
	/**
	 * Definitions of the NS_ constants are in Defines.php
	 * @private
	 */
	WgCanonicalNamespaceNames = map[int]string{
		consts.NS_MEDIA:          "Media",
		consts.NS_SPECIAL:        "Special",
		consts.NS_TALK:           "Talk",
		consts.NS_USER:           "User",
		consts.NS_USER_TALK:      "User_talk",
		consts.NS_PROJECT:        "Project",
		consts.NS_PROJECT_TALK:   "Project_talk",
		consts.NS_FILE:           "File",
		consts.NS_FILE_TALK:      "File_talk",
		consts.NS_MEDIAWIKI:      "MediaWiki",
		consts.NS_MEDIAWIKI_TALK: "MediaWiki_talk",
		consts.NS_TEMPLATE:       "Template",
		consts.NS_TEMPLATE_TALK:  "Template_talk",
		consts.NS_HELP:           "Help",
		consts.NS_HELP_TALK:      "Help_talk",
		consts.NS_CATEGORY:       "Category",
		consts.NS_CATEGORY_TALK:  "Category_talk",
	}
)
