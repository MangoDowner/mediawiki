package test

import "fmt"

func AssetEqual(answer interface{}, result interface{}, tip string) {
	if result == answer {
		fmt.Print("[PASS] ")
	} else {
		fmt.Print("[FAIL] ")
	}
	fmt.Println(tip)
	if result != answer {
		fmt.Println("> answer: ", answer)
		fmt.Println("> result: ", result)
	}
}

func AssetSame(answer interface{}, result interface{}, tip string) {
	fmt.Print("[TEST] ")
	fmt.Println(tip)
	fmt.Println("> result: ", result)
}

func AssetTrue(result bool, tip string) {
	if result {
		fmt.Print("[PASS] ")
	} else {
		fmt.Print("[FAIL] ")
	}
	fmt.Println(tip)
}

