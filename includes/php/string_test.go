package php

import (
	"fmt"
	"testing"
)

func TestExplode(t *testing.T) {
	str := "one,two,three,four"
	result := Explode(",", str, 0)
	fmt.Println(result)

	result1 := Explode(",", str, 1)
	fmt.Println(result1)

	result2 := Explode(",", str, 4)
	fmt.Println(result2)

	result3 := Explode(",", str, 10)
	fmt.Println(result3)
}
