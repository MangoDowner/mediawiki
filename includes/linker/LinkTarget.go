package linker

/**
 * @since 1.27
 */
type LinkTarget interface {

	/**
	 * Get the namespace index.
	 * @since 1.27
	 *
	 * @return int Namespace index
	 */
	GetNamespace() int

	/**
	 * Convenience function to test if it is in the namespace
	 * @since 1.27
	 *
	 * @param int $ns
	 * @return bool
	 */
	InNamespace(ns int)bool

	/**
	 * Get the link fragment (i.e. the bit after the #) in text form.
	 * @since 1.27
	 *
	 * @return string link fragment
	 */
	GetFragment() string

	/**
	 * Whether the link target has a fragment
	 * @since 1.27
	 *
	 * @return bool
	 */
	HasFragment() bool

	/**
	 * Get the main part with underscores.
	 * @since 1.27
	 *
	 * @return string Main part of the link, with underscores (for use in href attributes)
	 */
	GetDBkey() string

	/**
	 * Returns the link in text form, without namespace prefix or fragment.
	 * This is computed from the DB key by replacing any underscores with spaces.
	 * @since 1.27
	 *
	 * @return string
	 */
	GetText() string

	/**
	 * Creates a new LinkTarget for a different fragment of the same page.
	 * It is expected that the same type of object will be returned, but the
	 * only requirement is that it is a LinkTarget.
	 * @since 1.27
	 *
	 * @param string $fragment The fragment name, or "" for the entire page.
	 *
	 * @return LinkTarget
	 */
	CreateFragmentTarget(fragment string) *LinkTarget

	/**
	 * Whether this LinkTarget has an interwiki component
	 * @since 1.27
	 *
	 * @return bool
	 */
	IsExternal() bool

	/**
	 * The interwiki component of this LinkTarget
	 * @since 1.27
	 *
	 * @return string
	 */
	GetInterwiki() string

	/**
	 * Returns an informative human readable representation of the link target,
	 * for use in logging and debugging. There is no requirement for the return
	 * value to have any relationship with the input of TitleParser.
	 * @since 1.31
	 *
	 * @return string
	 */
	 ToString() string

}