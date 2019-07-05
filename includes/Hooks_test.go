package includes

import (
	"errors"
	test "github.com/MangoDowner/mediawiki/tests"
	"testing"
)

/**
 * @covers Hooks::getHandlers
 */
func TestGetHandlers(t *testing.T) {
	h := NewHooks()

	a := NewNothingClass()
	h.register("MediaWikiHooksTest001", []interface{}{a, "SomeNonStaticWithData"})
	//WgHooks["MediaWikiHooksTest001"][0] = a
	h.Run("MediaWikiHooksTest001", []interface{}{}, "")

}

/**
 * @covers Hooks::callHook
 * @expectedException PHPUnit_Framework_Error_Deprecated
 */
func TestCallHook_Deprecated(t *testing.T) {
	h := NewHooks()
	h.register("MediaWikiHooksTest001", NewNothingClass().SomeNonStaticWithData)
	// NothingClass::someStatic
	h.Run("MediaWikiHooksTest001", []interface{}{}, "1.31")
	test.AssetSame(
		"",
		nil,
		"FatalError",
	)
}

/**
 * @covers Hooks::runWithoutAbort
 * @covers Hooks::callHook
 */
func TestRunWithoutAbort(t *testing.T) {
	h := NewHooks()
	var list []int

	h.register("MediaWikiHooksTest001", func(...interface{}) interface{} {
		list = append(list, 1)
		return nil  // Explicit true
	})
	h.register("MediaWikiHooksTest001", func(...interface{}) interface{} {
		list = append(list, 2)
		return nil  // Implicit null
	})
	h.register("MediaWikiHooksTest001", func(...interface{}) interface{} {
		list = append(list, 3)
		return nil  // No return
	})
	h.RunWithoutAbort("MediaWikiHooksTest001", []interface{}{&list}, "")
	test.AssetSame(
		[]int{1, 2, 3},
		list,
		"All hooks ran.",
	)
}

/**
 * @covers Hooks::runWithoutAbort
 * @covers Hooks::callHook
 */
func TestRunWithoutAbortWarning(t *testing.T) {
	h := NewHooks()
	foo := "original"

	h.register("MediaWikiHooksTest001", func(param *string) error {
		return errors.New("foo error")
	})
	h.register("MediaWikiHooksTest001", func(param *string) interface{} {
		*param = "test"
		return nil
	})

	h.RunWithoutAbort("MediaWikiHooksTest001", []interface{}{&foo}, "")
	test.AssetSame(
		"test",
		foo,
		"Invalid return from hook-MediaWikiHooksTest001-closure for unabortable MediaWikiHooksTest001",
	)
}

/**
 * @expectedException FatalError
 * @covers Hooks::run
 */
func TestFatalError(t *testing.T) {
	h := NewHooks()

	h.register("MediaWikiHooksTest001", func(...interface{}) interface{} {
		return "test"
	})

	h.Run("MediaWikiHooksTest001", []interface{}{}, "")
	test.AssetSame(
		"",
		nil,
		"FatalError",
	)
}

type NothingClass struct {
	Calls int
}

func NewNothingClass() *NothingClass {
	this := new(NothingClass)
	return this
}

func (n *NothingClass) SomeNonStaticWithData(data, foo *string) interface{} {
	n.Calls++
	if data == nil || foo == nil {
		return errors.New("data/foo cant be nil")
	}
	*data = "changed-nonstatic"
	*foo = "changed-nonstatic"
	return nil
}
