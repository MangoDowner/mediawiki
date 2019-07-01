package php

import (
	"fmt"
	"testing"
)

func TestVarExport(t *testing.T) {
	v1 := map[string]interface{} {
		"Name": "Jack",
		"Gender": 1,
		"Age": 24,
	}
	result := VarExport(v1)
	fmt.Println(result)
}
