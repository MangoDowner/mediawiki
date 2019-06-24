package skins

/**
 * Fetch the set of available skins.
 * @return array Associative array of strings
 */
func GetSkinNames() {
	return SkinFactory::getDefaultInstance()->getSkinNames();
}