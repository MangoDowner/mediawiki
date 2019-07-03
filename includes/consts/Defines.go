package consts
/**@{
 * Virtual namespaces; don't appear in the page database
 */
const NS_MEDIA = -2
const NS_SPECIAL = -1

/**@{
 * Real namespaces
 *
 * Number 100 and beyond are reserved for custom namespaces;
 * DO NOT assign standard namespaces at 100 or beyond.
 * DO NOT Change integer values as they are most probably hardcoded everywhere
 * see T2696 which talked about that.
 */
const NS_MAIN = 0
const NS_TALK = 1
const NS_USER = 2
const NS_USER_TALK = 3
const NS_PROJECT = 4
const NS_PROJECT_TALK = 5
const NS_FILE = 6
const NS_FILE_TALK = 7
const NS_MEDIAWIKI = 8
const NS_MEDIAWIKI_TALK = 9
const NS_TEMPLATE = 10
const NS_TEMPLATE_TALK = 11
const NS_HELP = 12
const NS_HELP_TALK = 13
const NS_CATEGORY = 14
const NS_CATEGORY_TALK = 15

/**
 * NS_IMAGE and NS_IMAGE_TALK are the pre-v1.14 names for NS_FILE and
 * NS_FILE_TALK respectively, and are kept for compatibility.
 *
 * When writing code that should be compatible with older MediaWiki
 * versions, either stick to the old names or define the new constants
 * yourself, if they're not defined already.
 *
 * @deprecated since 1.14
 */
const S_IMAGE = NS_FILE

/**
 * @deprecated since 1.14
 */
const S_IMAGE_TALK = NS_FILE_TALK

