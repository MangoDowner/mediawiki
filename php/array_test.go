package php

import (
	"fmt"
	"testing"
)

func TestArrayMerge(t *testing.T) {
	arr1 := []string {"A"}
	arr2 := []string {"B"}
	arr3 := []string {"C"}
	result := ArrayMerge(arr1, arr2, arr3)
	fmt.Println(result)
}