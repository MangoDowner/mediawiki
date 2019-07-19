/**
 * A title formatter service for MediaWiki.
 */
package title

import "github.com/MangoDowner/mediawiki/includes/linker"

/**
 * A title formatter service for MediaWiki.
 *
 * This is designed to encapsulate knowledge about conventions for the title
 * forms to be used in the database, in urls, in wikitext, etc.
 *
 * @see https://www.mediawiki.org/wiki/Requests_for_comment/TitleValue
 * @since 1.23
 */

type TitleFormatter interface {
	/**
	 * Returns the title formatted for display.
	 * Per default, this includes the namespace but not the fragment.
	 *
	 * @note Normalization is applied if $title is not in TitleValue::TITLE_FORM.
	 *
	 * @param int|bool $namespace The namespace ID (or false, if the namespace should be ignored)
	 * @param string $text The page title
	 * @param string $fragment The fragment name (may be empty).
	 * @param string $interwiki The interwiki prefix (may be empty).
	 *
	 * @return string
	 */
	FormatTitle(namespace int, text, fragment, interWiki string)

	/**
	 * Returns the title text formatted for display, without namespace of fragment.
	 *
	 * @note Consider using LinkTarget::getText() directly, it's identical.
	 *
	 * @param LinkTarget $title The title to format
	 *
	 * @return string
	 */
	GetText(title *linker.LinkTarget)

	/**
	 * Returns the title formatted for display, including the namespace name.
	 *
	 * @param LinkTarget $title The title to format
	 *
	 * @return string
	 */
	GetPrefixedText(title *linker.LinkTarget)

	/**
	 * Return the title in prefixed database key form, with interwiki
	 * and namespace.
	 *
	 * @since 1.27
	 *
	 * @param LinkTarget $target
	 *
	 * @return string
	 */
	GetPrefixedDBkey(target  *linker.LinkTarget)

	/**
	 * Returns the title formatted for display, with namespace and fragment.
	 *
	 * @param LinkTarget $title The title to format
	 *
	 * @return string
	 */
	GetFullText(title *linker.LinkTarget)

	/**
	 * Returns the name of the namespace for the given title.
	 *
	 * @note This must take into account gender sensitive namespace names.
	 * @todo Move this to a separate interface
	 *
	 * @param int $namespace
	 * @param string $text
	 *
	 * @throws InvalidArgumentException
	 * @return string
	 */
	GetNamespaceName(namespace int, text string)
}
