package consts

import "github.com/MangoDowner/mediawiki/includes/libs/rdbms/database"

/**
 * @defgroup Constants MediaWiki constants
 */

// Obsolete aliases
/**
 * @deprecated since 1.28
 */
const DB_SLAVE = -1


/**@{
 * Virtual namespaces; don't appear in the page database
 */
const NS_MEDIA = -2
const NS_SPECIAL = -1

/**@{
 * Obsolete IDatabase::makeList() constants
 * These are also available as Database class constants
 */
const LIST_COMMA = database.LIST_COMMA
const LIST_AND = database.LIST_AND
const LIST_SET = database.LIST_SET
const LIST_NAMES = database.LIST_NAMES
const LIST_OR = database.LIST_OR

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


/** @{
 * Protocol constants for wfExpandUrl()
 */
const PROTO_HTTP = "http://"
const PROTO_HTTPS = "https://"
const PROTO_RELATIVE = "//"
const PROTO_CURRENT = ""
const PROTO_CANONICAL = 1
const PROTO_INTERNAL = 2

/**@}*/

/**@{
 * Content model ids, used by Content and ContentHandler.
 * These IDs will be exposed in the API and XML dumps.
 *
 * Extensions that define their own content model IDs should take
 * care to avoid conflicts. Using the extension name as a prefix is recommended,
 * for example 'myextension-somecontent'.
 */
const CONTENT_MODEL_WIKITEXT = "wikitext"
const CONTENT_MODEL_JAVASCRIPT = "javascript"
const CONTENT_MODEL_CSS = "css"
const CONTENT_MODEL_TEXT = "text"
const CONTENT_MODEL_JSON = "json"

/**@}*/

/**@{
 * Content formats, used by Content and ContentHandler.
 * These should be MIME types, and will be exposed in the API and XML dumps.
 *
 * Extensions are free to use the below formats, or define their own.
 * It is recommended to stick with the conventions for MIME types.
 */
// wikitext
const CONTENT_FORMAT_WIKITEXT = "text/x-wiki"
// for js pages
const CONTENT_FORMAT_JAVASCRIPT = "text/javascript"
// for css pages
const CONTENT_FORMAT_CSS = "text/css"
// for future use, e.g. with some plain-html messages.
const CONTENT_FORMAT_TEXT = "text/plain"
// for future use, e.g. with some plain-html messages.
const CONTENT_FORMAT_HTML = "text/html"
// for future use with the api and for extensions
const CONTENT_FORMAT_SERIALIZED = "application/vnd.php.serialized"
// for future use with the api, and for use by extensions
const CONTENT_FORMAT_JSON = "application/json"
// for future use with the api, and for use by extensions
const CONTENT_FORMAT_XML = "application/xml"
/**@}*/

/**@{
 * Max string length for shell invocations; based on binfmts.h
 */
const SHELL_MAX_ARG_STRLEN = "100000"

/**@}*/

/**@{
 * Schema compatibility flags.
 *
 * Used as flags in a bit field that indicates whether the old or new schema (or both)
 * are read or written.
 *
 * - SCHEMA_COMPAT_WRITE_OLD: Whether information is written to the old schema.
 * - SCHEMA_COMPAT_READ_OLD: Whether information stored in the old schema is read.
 * - SCHEMA_COMPAT_WRITE_NEW: Whether information is written to the new schema.
 * - SCHEMA_COMPAT_READ_NEW: Whether information stored in the new schema is read.
 */
const SCHEMA_COMPAT_WRITE_OLD = 0x01
const SCHEMA_COMPAT_READ_OLD = 0x02
const SCHEMA_COMPAT_WRITE_NEW = 0x10
const SCHEMA_COMPAT_READ_NEW = 0x20
const SCHEMA_COMPAT_WRITE_BOTH = SCHEMA_COMPAT_WRITE_OLD | SCHEMA_COMPAT_WRITE_NEW
const SCHEMA_COMPAT_READ_BOTH = SCHEMA_COMPAT_READ_OLD | SCHEMA_COMPAT_READ_NEW
const SCHEMA_COMPAT_OLD = SCHEMA_COMPAT_WRITE_OLD | SCHEMA_COMPAT_READ_OLD
const SCHEMA_COMPAT_NEW = SCHEMA_COMPAT_WRITE_NEW | SCHEMA_COMPAT_READ_NEW

/**@}*/

/**@{
 * Schema change migration flags.
 *
 * Used as values of a feature flag for an orderly transition from an old
 * schema to a new schema. The numeric values of these constants are compatible with the
 * SCHEMA_COMPAT_XXX bitfield semantics. High bits are used to ensure that the numeric
 * ordering follows the order in which the migration stages should be used.
 *
 * - MIGRATION_OLD: Only read and write the old schema. The new schema need not
 *   even exist. This is used from when the patch is merged until the schema
 *   change is actually applied to the database.
 * - MIGRATION_WRITE_BOTH: Write both the old and new schema. Read the new
 *   schema preferentially, falling back to the old. This is used while the
 *   change is being tested, allowing easy roll-back to the old schema.
 * - MIGRATION_WRITE_NEW: Write only the new schema. Read the new schema
 *   preferentially, falling back to the old. This is used while running the
 *   maintenance script to migrate existing entries in the old schema to the
 *   new schema.
 * - MIGRATION_NEW: Only read and write the new schema. The old schema (and the
 *   feature flag) may now be removed.
 */
const MIGRATION_OLD = 0x00000000 | SCHEMA_COMPAT_OLD
const MIGRATION_WRITE_BOTH = 0x10000000 | SCHEMA_COMPAT_READ_BOTH | SCHEMA_COMPAT_WRITE_BOTH
const MIGRATION_WRITE_NEW = 0x20000000 | SCHEMA_COMPAT_READ_BOTH | SCHEMA_COMPAT_WRITE_NEW
const MIGRATION_NEW = 0x30000000 | SCHEMA_COMPAT_NEW
/**@}*/
