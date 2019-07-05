/**
 * Default values for MediaWiki configuration settings.
 *
 */
package includes

var (

	/**
	 * Temporary variable that applies MediaWiki UI wherever it can be supported.
	 * Temporary variable that should be removed when mediawiki ui is more
	 * stable and change has been communicated.
	 * @since 1.24
	 */
	WgUseMediaWikiUIEverywhere = false

	/**
	 * Global list of hooks.
	 *
	 * The key is one of the events made available by MediaWiki, you can find
	 * a description for most of them in docs/hooks.txt. The array is used
	 * internally by Hook:run().
	 *
	 * The value can be one of:
	 *
	 * - A function name:
	 * @code
	 *     $wgHooks['event_name'][] = $function;
	 * @endcode
	 * - A function with some data:
	 * @code
	 *     $wgHooks['event_name'][] = [ $function, $data ];
	 * @endcode
	 * - A an object method:
	 * @code
	 *     $wgHooks['event_name'][] = [ $object, 'method' ];
	 * @endcode
	 * - A closure:
	 * @code
	 *     $wgHooks['event_name'][] = function ( $hookParam ) {
	 *         // Handler code goes here.
	 *     };
	 * @endcode
	 *
	 * @warning You should always append to an event array or you will end up
	 * deleting a previous registered hook.
	 *
	 * @warning Hook handlers should be registered at file scope. Registering
	 * handlers after file scope can lead to unexpected results due to caching.
	 */
	WgHooks = map[string][]HookFunc{}




	/**
	 * Array of allowed values for the "title=foo&action=<action>" parameter. Syntax is:
	 *     "foo" : "ClassName"    Load the specified class which subclasses Action
	 *     "foo" : true           Load the class FooAction which subclasses Action
	 *                             If something is specified in the getActionOverrides()
	 *                             of the relevant Page object it will be used
	 *                             instead of the default class.
	 *     "foo" : false          The action is disabled; show an error message
	 * Unsetting core actions will probably cause things to complain loudly.
	 */
	WgActions = map[string]interface{}{
		"credits" : true,
		"delete" : true,
		"edit" : true,
		"editchangetags" : nil,
		"history" : true,
		"info" : true,
		"markpatrolled" : true,
		"mcrundo" : nil,
		"mcrrestore" : nil,
		"protect" : true,
		"purge" : true,
		"raw" : true,
		"render" : true,
		"revert" : true,
		"revisiondelete" : nil,
		"rollback" : true,
		"submit" : true,
		"unprotect" : true,
		"unwatch" : true,
		"view" : true,
		"watch" : true,
	}
)