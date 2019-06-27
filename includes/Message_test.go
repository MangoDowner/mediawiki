package includes

import (
	"fmt"
	"testing"
)

func TestParams(t *testing.T) {
	m := new(Message)
	h := new(Html)
	var result string

	result = m.Params(1).Parse()
	fmt.Println(result)

	result = m.Params("page title").Parse()
	fmt.Println(result)

	result = m.Params("page title", 1).Parse()
	fmt.Println(result)

	result = m.Params([]string{"1", "2", "3"}, 1).Parse()
	fmt.Println(result)

	result = m.Params(map[string]string{"1" : "Tom", "2" : "Jack", "3" : "Helen"}, 1).Parse()
	fmt.Println(result)

	result = m.Params(h.Element("span", map[string]interface{}{"dir":"auto"}, "")).Parse()
	fmt.Println(result)
}