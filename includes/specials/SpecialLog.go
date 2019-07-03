/**
 * Implements Special:Log
 */
package specials

/**
 * A special page that lists log entries
 *
 * @ingroup SpecialPage
 */
type SpecialLog struct {

}

func NewSpecialLog() *SpecialLog {
	this := new(SpecialLog)
	return this
}

func (s *SpecialLog) Msg() {

}


