package includes

import (
	"fmt"
	test "github.com/MangoDowner/mediawiki/tests"
	"go-common/app/admin/main/macross/model/errors"
	"testing"
)

/**
 * @covers Hooks::getHandlers
 */
func TestGetHandlers(t *testing.T) {
	h := new(Hooks)

	test.AssetSame(
		[]func(){},
		h.GetHandlers("MediaWikiHooksTest001"),
		"No hooks registered",
	)

	//a := NewNothingClass()
	//h.register("MediaWikiHooksTest001", a)
	//WgHooks["MediaWikiHooksTest001"][0] = a
	test.AssetSame(
		[]func(){},
		h.GetHandlers("MediaWikiHooksTest001"),
		"Hook registered by $wgHooks",
	)

}

/**
 * @covers Hooks::runWithoutAbort
 * @covers Hooks::callHook
 */
func TestRunWithoutAbort(t *testing.T) {
	h := NewHooks()
	var list []int

	h.register("MediaWikiHooksTest001", func() error{
		list = append(list, 1)
		fmt.Println("1 TIME")
		return nil  // Explicit true
	})
	h.register("MediaWikiHooksTest001", func() error{
		list = append(list, 2)
		fmt.Println("2 TIME")
		return nil  // Implicit null
	})
	h.register("MediaWikiHooksTest001", func() error{
		list = append(list, 3)
		fmt.Println("3 TIME")
		return nil  // No return
	})
	h.RunWithoutAbort("MediaWikiHooksTest001", []interface{}{&list}, "")
	test.AssetSame(
		[]int{1, 2, 3},
		list,
		"All hooks ran.",
	)
}

type NothingClass struct {
	Calls int
}

func NewNothingClass() *NothingClass {
	this := new(NothingClass)
	return this
}

func (n *NothingClass) SomeNonStaticWithData(data, foo *string) error {
	n.Calls++
	if data == nil || foo == nil {
		return errors.New("data/foo cant be nil", nil)
	}
	*data = "changed-nonstatic"
	*foo = "changed-nonstatic"
	return nil
}
