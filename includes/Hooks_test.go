package includes

import (
	"errors"
	"fmt"
	test "github.com/MangoDowner/mediawiki/tests"
	"testing"
)

/**
 * @covers Hooks::getHandlers
 */
func TestGetHandlers(t *testing.T) {
	h := NewHooks()

	a := NewNothingClass()
	p1 := "p1"
	p2 := "p2"
	h.register("MediaWikiHooksTest001", a.SomeNonStaticWithData)
	h.clear("MediaWikiHooksTest001")
	WgHooks["MediaWikiHooksTest001"] = append(WgHooks["MediaWikiHooksTest001"], a.SomeNonStaticWithData)
	h.Run("MediaWikiHooksTest001", []interface{}{&p1, &p2}, "")
	fmt.Println(p1)
	fmt.Println(p2)
}

/**
 * @covers Hooks::run
 * @covers Hooks::callHook
 */
func TestFalseReturn(t *testing.T) {
	h := NewHooks()
	foo := "original"

	h.register("MediaWikiHooksTest001", func(param *string) bool {
		return false
	})

	h.register("MediaWikiHooksTest001", func(param *string) bool {
		*param = "test"
		return true
	})

	h.Run("MediaWikiHooksTest001", []interface{}{&foo}, "")
	test.AssetEqual(
		"original",
		foo,
		"Hooks abort after a false return.",
	)
}

/**
 * @covers Hooks::isRegistered
 * @covers Hooks::register
 * @covers Hooks::run
 * @covers Hooks::callHook
 */
func TestNewStyleHookInteraction(t *testing.T) {
	h := NewHooks()
	a := NewNothingClass()
	b := NewNothingClass()

	WgHooks["MediaWikiHooksTest001"] = append(WgHooks["MediaWikiHooksTest001"], a.SomeNonStaticWithData)
	test.AssetTrue(
		h.IsRegistered("MediaWikiHooksTest001"),
		"Hook registered via $wgHooks should be noticed by Hooks::isRegistered",
	)

	WgHooks["MediaWikiHooksTest001"] = append(WgHooks["MediaWikiHooksTest001"], b.SomeNonStaticWithData)
	test.AssetEqual(
		2,
		len(h.GetHandlers("MediaWikiHooksTest001")),
		"Hooks::getHandlers() should return hooks registered via wgHooks as well as Hooks::register",
	)

	foo := "quux"
	bar := "gaax"

	h.Run("MediaWikiHooksTest001", []interface{}{&foo, &bar}, "")
	test.AssetEqual(
		1,
		a.Calls,
		"Hooks::run() should run hooks registered via wgHooks as well as Hooks::register",
	)
	test.AssetEqual(
		1,
		b.Calls,
		"Hooks::run() should run hooks registered via wgHooks as well as Hooks::register",
	)
}

/**
 * 字符串类型hook暂不支持
 * @expectedException MWException
 * @covers Hooks::run
 * @covers Hooks::callHook
 */
func TestUncallableFunction(t *testing.T) {
	h := NewHooks()
	h.register("MediaWikiHooksTest001", "ThisFunctionDoesntExist")
	h.Run("MediaWikiHooksTest001", []interface{}{}, "")
}

/**
 * @covers Hooks::run
 */
func TestNullReturn(t *testing.T) {
	h := NewHooks()
	foo := "original"

	h.register("MediaWikiHooksTest001", func(param *string) {
		return
	})

	h.register("MediaWikiHooksTest001", func(param *string) bool {
		*param = "test"
		return true
	})

	h.Run("MediaWikiHooksTest001", []interface{}{&foo}, "")
	test.AssetEqual(
		"test",
		"test",
		"Hooks continue after a null return.",
	)
}

/**
 * 无测试价值，不应传入false变量
 * @covers Hooks::callHook
 */
func TestCallHook_FalseHook(t *testing.T) {
	h := NewHooks()
	foo := "original"

	h.register("MediaWikiHooksTest001", false)

	h.register("MediaWikiHooksTest001", func(param *string) error {
		return errors.New("func error")
	})

	h.Run("MediaWikiHooksTest001", []interface{}{&foo}, "1.31")
	test.AssetEqual(
		"test",
		"test",
		"Hooks that are falsey are skipped.",
	)
}

/**
 * @covers Hooks::callHook
 * @expectedException MWException
 */
func TestCallHook_(t *testing.T) {
	h := NewHooks()
	h.register("MediaWikiHooksTest001", 123456)
	// NothingClass::someStatic
	h.Run("MediaWikiHooksTest001", []interface{}{}, "1.31")
	test.AssetSame(
		"",
		nil,
		"FatalError",
	)
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
		return errors.New("LAMBDA func error")
	})

	testFunc := func (param *string) bool {
		fmt.Println("testFunc")
		return true
	}

	h.register("MediaWikiHooksTest001", func(param *string) error {
		return errors.New("func error")
	})
	h.register("MediaWikiHooksTest001", testFunc)

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
	*data = "data changed-nonstatic"
	*foo = "foo changed-nonstatic"
	return nil
}
