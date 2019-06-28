/**
 * Implements the User class for the %MediaWiki software.
 */
package user

import (
	"github.com/MangoDowner/mediawiki/includes/dao"
	"github.com/MangoDowner/mediawiki/includes/session"
)

/**
 * @const int Number of characters in user_token field.
 */
const TOKEN_LENGTH = 32

/**
 * @const string An invalid value for user_token
 */
const INVALID_TOKEN = "*** INVALID ***"

/**
 * Global constant made accessible as class constants so that autoloader
 * magic can be used.
 * @deprecated since 1.27, use \MediaWiki\Session\Token::SUFFIX
 */
const EDIT_TOKEN_SUFFIX = session.TOKEN_SUFFIX

/**
 * @const int Serialized record version.
 */
const VERSION = 12

/**
 * Exclude user options that are set to their default value.
 * @since 1.25
 */
const GETOPTIONS_EXCLUDE_DEFAULTS = 1

/**
 * @since 1.27
 */
const CHECK_USER_RIGHTS = true

/**
 * @since 1.27
 */
const IGNORE_USER_RIGHTS = false

/**
 * Array of Strings List of member variables which are saved to the
 * shared cache (memcached). Any operation which changes the
 * corresponding database fields must call a cache-clearing function.
 * @showinitializer
 */
var mCacheVars = []string{
	// user table
	"mId",
	"mName",
	"mRealName",
	"mEmail",
	"mTouched",
	"mToken",
	"mEmailAuthenticated",
	"mEmailToken",
	"mEmailTokenExpires",
	"mRegistration",
	"mEditCount",
	// user_groups table
	"mGroupMemberships",
	// user_properties table
	"mOptionOverrides",
	// actor table
	"mActorId",
}

/**
 * Array of Strings Core rights.
 * Each of these should have a corresponding message of the form
 * "right-$right".
 * @showinitializer
 */
var mCoreRights = []string{
	"apihighlimits",
	"applychangetags",
	"autoconfirmed",
	"autocreateaccount",
	"autopatrol",
	"bigdelete",
	"block",
	"blockemail",
	"bot",
	"browsearchive",
	"changetags",
	"createaccount",
	"createpage",
	"createtalk",
	"delete",
	"deletechangetags",
	"deletedhistory",
	"deletedtext",
	"deletelogentry",
	"deleterevision",
	"edit",
	"editcontentmodel",
	"editinterface",
	"editprotected",
	"editmyoptions",
	"editmyprivateinfo",
	"editmyusercss",
	"editmyuserjson",
	"editmyuserjs",
	"editmywatchlist",
	"editsemiprotected",
	"editsitecss",
	"editsitejson",
	"editsitejs",
	"editusercss",
	"edituserjson",
	"edituserjs",
	"hideuser",
	"import",
	"importupload",
	"ipblock-exempt",
	"managechangetags",
	"markbotedits",
	"mergehistory",
	"minoredit",
	"move",
	"movefile",
	"move-categorypages",
	"move-rootuserpages",
	"move-subpages",
	"nominornewtalk",
	"noratelimit",
	"override-export-depth",
	"pagelang",
	"patrol",
	"patrolmarks",
	"protect",
	"purge",
	"read",
	"reupload",
	"reupload-own",
	"reupload-shared",
	"rollback",
	"sendemail",
	"siteadmin",
	"suppressionlog",
	"suppressredirect",
	"suppressrevision",
	"unblockself",
	"undelete",
	"unwatchedpages",
	"upload",
	"upload_by_url",
	"userrights",
	"userrights-interwiki",
	"viewmyprivateinfo",
	"viewmywatchlist",
	"viewsuppressed",
	"writeapi",
}

/**
 * The User object encapsulates all of the user-specific settings (user_id,
 * name, rights, email address, options, last login time). Client
 * classes use the getXXX() functions to access these fields. These functions
 * do all the work of determining whether the user is logged in,
 * whether the requested option can be satisfied from cookies or
 * whether a database query is needed. Most of the settings needed
 * for rendering normal pages are set in the cookie to minimize use
 * of the database.
 */
type User struct {
	/**
	 * String Cached results of getAllRights()
	 */
	mAllRights string

	/** Cache variables */
	// @{
	/** @var int */
	MId int
	/** @var string */
	MName string
	/** @var int|null */
	mActorId int
	/** @var string */
	MRealName string

	/** @var string */
	MEmail string
	/** @var string TS_MW timestamp from the DB */
	MTouched string
	/** @var string TS_MW timestamp from cache */
	mQuickTouched string
	/** @var string */
	mToken string
	/** @var string */
	MEmailAuthenticated string
	/** @var string */
	mEmailToken string
	/** @var string */
	mEmailTokenExpires string
	/** @var string */
	mRegistration string
	/** @var int */
	mEditCount int
	/** @var UserGroupMembership[] Associative array of (group name => UserGroupMembership object) */
	mGroupMemberships []string
	/** @var array */
	mOptionOverrides []string
	// @}

	/**
	 * Bool Whether the cache variables have been loaded.
	 */
	// @{
	MOptionsLoaded bool

	/**
	 * Array with already loaded items or true if all items have been loaded.
	 */
	mLoadedItems []string
	// @}

	/**
	 * String Initialization data source if mLoadedItems!==true. May be one of:
	 *  - 'defaults'   anonymous user initialised from class defaults
	 *  - 'name'       initialise from mName
	 *  - 'id'         initialise from mId
	 *  - 'actor'      initialise from mActorId
	 *  - 'session'    log in from session if possible
	 *
	 * Use the User::newFrom*() family of functions to set this.
	 */
	MFrom string

	/**
	 * Lazy-initialized variables, invalidated with clearInstanceCache
	 */
	mNewtalk interface{}
	/** @var string */
	mDatePreference string
	/** @var string */
	mBlockedby string
	/** @var string */
	mHash string
	/** @var array */
	MRights string
	/** @var string */
	mBlockreason string
	/** @var array */
	mEffectiveGroups []string
	/** @var array */
	mImplicitGroups []string
	/** @var array */
	mFormerGroups []string
	/** @var Block */
	mGlobalBlock interface{}
	/** @var bool */
	mLocked bool
	/** @var bool */
	mHideName bool
	/** @var array */
	MOptions []string

	/** @var WebRequest */
	mRequest interface{}

	/** @var Block */
	MBlock interface{}

	/** @var bool */
	mAllowUsertalk bool

	/** @var Block */
	mBlockedFromCreateAccount interface{}

	/** @var int User::READ_* constant bitfield used to load data */
	queryFlagsUsed int

	IdCacheByName []string
}
/**
 * Lightweight constructor for an anonymous user.
 * Use the User::newFrom* factory functions for other kinds of users.
 *
 * @see newFromName()
 * @see newFromId()
 * @see newFromActorId()
 * @see newFromConfirmationCode()
 * @see newFromSession()
 * @see newFromRow()
 */
func NewUser() *User {
	this := new(User)
	this.queryFlagsUsed = dao.READ_NORMAL
	this.clearInstanceCache("defaults")
	return this
}

func (u *User) clearInstanceCache(reloadFrom string) {
	u.mNewtalk = -1
	u.mDatePreference = ""
	u.mBlockedby = "" // unset
	u.mHash = ""
	u.MRights = ""
	u.mEffectiveGroups = []string{}
	u.mImplicitGroups = []string{}
	u.mGroupMemberships = []string{}
	u.MOptions = []string{}
	u.MOptionsLoaded = false
	u.mEditCount = 0

	if reloadFrom != "" {
		u.mLoadedItems = []string{}
		u.MFrom = reloadFrom
	}
}

