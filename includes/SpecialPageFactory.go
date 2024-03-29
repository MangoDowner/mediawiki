/**
 * Factory for handling the special page list and generating SpecialPage objects.
 */
package includes

import (
	"fmt"
	"github.com/MangoDowner/mediawiki/includes/config"
	"github.com/MangoDowner/mediawiki/includes/languages"
	"github.com/MangoDowner/mediawiki/includes/linker"
	"github.com/MangoDowner/mediawiki/includes/php"
	"github.com/MangoDowner/mediawiki/includes/specials"
	"github.com/astaxie/beego/context"
)

/**
 * Factory for handling the special page list and generating SpecialPage objects.
 *
 * To add a special page in an extension, add to $wgSpecialPages either
 * an object instance or an array containing the name and constructor
 * parameters. The latter is preferred for performance reasons.
 *
 * The object instantiated must be either an instance of SpecialPage or a
 * sub-class thereof. It must have an execute() method, which sends the HTML
 * for the special page to $wgOut. The parent class has an execute() method
 * which distributes the call to the historical global functions. Additionally,
 * execute() also checks if the user has the necessary access privileges
 * and bails out if not.
 *
 * To add a core special page, use the similar static list in
 * SpecialPageFactory::$list. To remove a core static special page at runtime, use
 * a SpecialPage_initList hook.
 *
 * @note There are two classes called SpecialPageFactory.  You should use this first one, in
 * namespace MediaWiki\Special, which is a service.  \SpecialPageFactory is a deprecated collection
 * of static methods that forwards to the global service.
 *
 * @ingroup SpecialPage
 * @since 1.17
 */
type SpecialPageFactory struct {
	/**
	 * List of special page names to the subclass of SpecialPage which handles them.
	 * @todo Make this a const when we drop HHVM support (T192166).  It can still be private in PHP
	 * 7.1.
	 */
	coreList map[string]ISpecialPage

	/** @var array Special page name : class name */
	list map[string]ISpecialPage

	/** @var array */
	aliases map[string]string

	/** @var Config */
	config interface{}

	/** @var Language */
	contLang *languages.Language

	test ISpecialPage
}

/**
 * @param Config $config
 * @param Language $contLang
 */
func NewSpecialPageFactory() *SpecialPageFactory {
	this := new(SpecialPageFactory)
	//TODO: 补全
	this.coreList = map[string]ISpecialPage{
		// Maintenance Reports
		"BrokenRedirects" : &specials.SpecialLog{},
		"Deadendpages" : &specials.SpecialLog{},
		"DoubleRedirects" : &specials.SpecialLog{},
		"Longpages" : &specials.SpecialLog{},
		"Ancientpages" : &specials.SpecialLog{},
		"Lonelypages" : &specials.SpecialLog{},
		"Fewestrevisions" : &specials.SpecialLog{},
		"Withoutinterwiki" : &specials.SpecialLog{},
		"Protectedpages" : &specials.SpecialLog{},
		"Protectedtitles" : &specials.SpecialLog{},
		"Shortpages" : &specials.SpecialLog{},
		"Uncategorizedcategories" : &specials.SpecialLog{},
		"Uncategorizedimages" : &specials.SpecialLog{},
		"Uncategorizedpages" : &specials.SpecialLog{},
		"Uncategorizedtemplates" : &specials.SpecialLog{},
		"Unusedcategories" : &specials.SpecialLog{},
		"Unusedimages" : &specials.SpecialLog{},
		"Unusedtemplates" : &specials.SpecialLog{},
		"Unwatchedpages" : &specials.SpecialLog{},
		"Wantedcategories" : &specials.SpecialLog{},
		"Wantedfiles" : &specials.SpecialLog{},
		"Wantedpages" : &specials.SpecialLog{},
		"Wantedtemplates" : &specials.SpecialLog{},

		// List of pages
		"Allpages" : &specials.SpecialLog{},
		"Prefixindex" : &specials.SpecialLog{},
		"Categories" : &specials.SpecialLog{},
		"Listredirects" : &specials.SpecialLog{},
		"PagesWithProp" : &specials.SpecialLog{},
		"TrackingCategories" : &specials.SpecialLog{},

		// Authentication
		"Userlogin" : &specials.SpecialLog{},
		"Userlogout" : &specials.SpecialLog{},
		"CreateAccount" : &specials.SpecialLog{},
		"LinkAccounts" : &specials.SpecialLog{},
		"UnlinkAccounts" : &specials.SpecialLog{},
		"ChangeCredentials" : &specials.SpecialLog{},
		"RemoveCredentials" : &specials.SpecialLog{},

		// Users and rights
		"Activeusers" : &specials.SpecialLog{},
		"Block" : &specials.SpecialLog{},
		"Unblock" : &specials.SpecialLog{},
		"BlockList" : &specials.SpecialLog{},
		"AutoblockList" : &specials.SpecialLog{},
		"ChangePassword" : &specials.SpecialLog{},
		"BotPasswords" : &specials.SpecialLog{},
		"PasswordReset" : &specials.SpecialLog{},
		"DeletedContributions" : &specials.SpecialLog{},
		"Preferences" : &specials.SpecialLog{},
		"ResetTokens" : &specials.SpecialLog{},
		"Contributions" : &specials.SpecialLog{},
		"Listgrouprights" : &specials.SpecialLog{},
		"Listgrants" : &specials.SpecialLog{},
		"Listusers" : &specials.SpecialLog{},
		"Listadmins" : &specials.SpecialLog{},
		"Listbots" : &specials.SpecialLog{},
		"Userrights" : &specials.SpecialLog{},
		"EditWatchlist" : &specials.SpecialLog{},
		"PasswordPolicies" : &specials.SpecialLog{},

		// Recent changes and logs
		"Newimages" : &specials.SpecialLog{},
		"Log" : &specials.SpecialLog{},
		"Watchlist" : &specials.SpecialLog{},
		"Newpages" : &specials.SpecialLog{},
		"Recentchanges" : &specials.SpecialLog{},
		"Recentchangeslinked" : &specials.SpecialLog{},
		"Tags" : &specials.SpecialLog{},

		// Media reports and uploads
		"Listfiles" : &specials.SpecialLog{},
		"Filepath" : &specials.SpecialLog{},
		"MediaStatistics" : &specials.SpecialLog{},
		"MIMEsearch" : &specials.SpecialLog{},
		"FileDuplicateSearch" : &specials.SpecialLog{},
		"Upload" : &specials.SpecialLog{},
		"UploadStash" : &specials.SpecialLog{},
		"ListDuplicatedFiles" : &specials.SpecialLog{},

		// Data and tools
		"ApiSandbox" : &specials.SpecialLog{},
		"Statistics" : &specials.SpecialLog{},
		"Allmessages" : &specials.SpecialLog{},
		"Version" : &specials.SpecialLog{},
		"Lockdb" : &specials.SpecialLog{},
		"Unlockdb" : &specials.SpecialLog{},

		// Redirecting special pages
		"LinkSearch" : &specials.SpecialLog{},
		"Randompage" : &specials.SpecialLog{},
		"RandomInCategory" : &specials.SpecialLog{},
		"Randomredirect" : &specials.SpecialLog{},
		"Randomrootpage" : &specials.SpecialLog{},
		"GoToInterwiki" : &specials.SpecialLog{},

		// High use pages
		"Mostlinkedcategories" : &specials.SpecialLog{},
		"Mostimages" : &specials.SpecialLog{},
		"Mostinterwikis" : &specials.SpecialLog{},
		"Mostlinked" : &specials.SpecialLog{},
		"Mostlinkedtemplates" : &specials.SpecialLog{},
		"Mostcategories" : &specials.SpecialLog{},
		"Mostrevisions" : &specials.SpecialLog{},

		// Page tools
		"ComparePages" : &specials.SpecialLog{},
		"Export" : &specials.SpecialLog{},
		"Import" : &specials.SpecialLog{},
		"Undelete" : &specials.SpecialLog{},
		"Whatlinkshere" : &specials.SpecialLog{},
		"MergeHistory" : &specials.SpecialLog{},
		"ExpandTemplates" : &specials.SpecialLog{},

		// Other
		"Booksources" : &specials.SpecialLog{},

		// Unlisted / redirects
		"ApiHelp" : &specials.SpecialLog{},
		"Blankpage" : &specials.SpecialLog{},
		"Diff" : &specials.SpecialLog{},
		"EditTags" : &specials.SpecialLog{},
		"Emailuser" : &specials.SpecialLog{},
		"Movepage" : &specials.SpecialLog{},
		"Mycontributions" : &specials.SpecialLog{},
		"MyLanguage" : &specials.SpecialLog{},
		"Mypage" : &specials.SpecialLog{},
		"Mytalk" : &specials.SpecialLog{},
		"Myuploads" : &specials.SpecialLog{},
		"AllMyUploads" : &specials.SpecialLog{},
		"PermanentLink" : &specials.SpecialLog{},
		"Redirect" : &specials.SpecialLog{},
		"Revisiondelete" : &specials.SpecialLog{},
		"RunJobs" : &specials.SpecialLog{},
		"Specialpages" : &specials.SpecialLog{},
		"PageData" : &specials.SpecialLog{},
	}

	return this
}

/**
 * Returns a list of canonical special page names.
 * May be used to iterate over all registered special pages.
 *
 * @return string[]
 */
func (s *SpecialPageFactory) GetNames() (ret []string) {
	list := s.getPageList()
	for k := range list {
		ret = append(ret, k)
	}
	return ret
}

/**
 * Get the special page list as an array
 *
 * @return array
 */
func (s *SpecialPageFactory) getPageList() map[string]ISpecialPage {
	if len(s.list) != 0 {
		return s.list
	}
	s.list = s.coreList

	if !config.Configs.GetBool("DisableInternalSearch") {
		s.list["Search"] = &specials.SpecialLog{}
	}

	// TODO: 陆续补全后面的工厂类
	if config.Configs.GetBool( "EmailAuthentication" ) {
		s.list["Confirmemail"] = &specials.SpecialLog{}
		s.list["Invalidateemail"] = &specials.SpecialLog{}
	}

	if config.Configs.GetBool( "EnableEmail" ) {
		s.list["ChangeEmail"] = &specials.SpecialLog{}
	}

	if config.Configs.GetBool( "EnableJavaScriptTest" ) {
		s.list["JavaScriptTest"] = &specials.SpecialLog{}
	}

	if config.Configs.GetBool( "PageLanguageUseDB" ) {
		s.list["PageLanguage"] = &specials.SpecialLog{}
	}

	if config.Configs.GetBool( "ContentHandlerUseDB" ) {
		s.list["ChangeContentModel"] = &specials.SpecialLog{}
	}

	// Add extension special pages
	//$this->list = array_merge( $this->list, $this->config->get( "SpecialPages" ) );

	// This hook can be used to disable unwanted core special pages
	// or conditionally register special pages.
	//Hooks::run( "SpecialPage_initList", [ &$this->list ] );
	return s.list
}

/**
 * Initialise and return the list of special page aliases. Returns an array where
 * the key is an alias, and the value is the canonical name of the special page.
 * All registered special pages are guaranteed to map to themselves.
 * @return array
 */
func (s *SpecialPageFactory) GetAliasList() (ret map[string]string) {
	if len(s.aliases) != 0 {
		return s.aliases
	}
	//aliases := s.contLang.GetSpecialPageAliases()
	return
}

/**
 * Given a special page name with a possible subpage, return an array
 * where the first element is the special page name and the second is the
 * subpage.
 *
 * @param string $alias
 * @return array Array( String, String|null ), or array( null, null ) if the page is invalid
 */
func (s *SpecialPageFactory) ResolveAlias(alias string) (ret []string) {
	bits := php.Explode("/", alias, 2)

	caseFoldedAlias := s.contLang.CaseFold(bits[0])
	caseFoldedAlias = php.Strtr(caseFoldedAlias, map[string]string{" " : "_"})
	//alias := s.GetAliasList()
	return ret
}

/**
 * Find the object with a given name and return it (or NULL)
 *
 * @param string $name Special page name, may be localised and/or an alias
 * @return SpecialPage|null SpecialPage object or null if the page doesn't exist
 */
func (s *SpecialPageFactory) GetPage(name string) bool {

	return true
}

/**
 * Execute a special page path.
 * The path may contain parameters, e.g. Special:Name/Params
 * Extracts the special page name and call the execute method, passing the parameters
 *
 * Returns a title object if the page is redirected, false if there was no such special
 * page, and true if it was successful.
 *
 * @param Title &$title
 * @param IContextSource &$context
 * @param bool $including Bool output is being captured for use in {{special:whatever}}
 * @param LinkRenderer|null $linkRenderer (since 1.28)
 *
 * @return bool|Title
 */
func (s *SpecialPageFactory) ExecutePath(title *Title, context *context.Context, including bool,
 	linkRenderer *linker.LinkRenderer) (ret *Title) {
	// @todo FIXME: Redirects broken due to this call
	bits := php.Explode("/", title.GetDBKey(), 2)
	name := bits[0]
	var par string
	// T4087
	if len(bits) >= 2 && bits[1] != "" {
		par = bits[1]
	}

	page := s.GetPage(name)
	if !page {

	}
	fmt.Println(par)
	return
}