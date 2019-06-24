package skins

/**
 * Factory class to create Skin objects
 *
 * @since 1.24
 */

type SkinFactory struct {

}

/**
 * @deprecated in 1.27
 * @return SkinFactory
 */
func GetDefaultInstance() {
	return MediaWikiServices::getInstance()->getSkinFactory();
}