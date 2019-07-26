package linker

/**
 * Class that generates HTML <a> links for pages.
 *
 * @see https://www.mediawiki.org/wiki/Manual:LinkRenderer
 * @since 1.28
 */

type LinkRenderer struct {

	/**
	 * Whether to force the pretty article path
	 *
	 * @var bool
	 */
	forceArticlePath bool

	/**
	 * A PROTO_* constant or false
	 *
	 * @var string|bool|int
	 */
	expandUrls string

	/**
	 * @var int
	 */
	stubThreshold int

	/**
	 * @var TitleFormatter
	 */
	titleFormatter string

	/**
	 * @var LinkCache
	 */
	linkCache string

	/**
	 * Whether to run the legacy Linker hooks
	 *
	 * @var bool
	 */
	runLegacyBeginHook bool

}

/**
 * @param TitleFormatter $titleFormatter
 * @param LinkCache $linkCache
 */
func NewLinkRenderer() *LinkRenderer {
	this := new(LinkRenderer)
	this.runLegacyBeginHook = true
	return this
}

/**
 * @param bool $force
 */
func (l *LinkRenderer) SetForceArticlePath(force bool){
	l.forceArticlePath = force
}