package resourceloader

import (
	"fmt"
	"testing"
)

func TestExpandModuleNames(t *testing.T) {
	result := expandModuleNames("jquery.foo,bar|jquery.ui.baz,quux")
	fmt.Println(result)
}
