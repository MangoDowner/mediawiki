package languages

import (
	"fmt"
	test "github.com/MangoDowner/mediawiki/tests"
	"testing"
)

func TestUc(t *testing.T) {

	str1 := "mediawiki"
	str2 := "百科mediawiki"

	result1 := NewLanguage().Uc(str1, true)
	fmt.Println(result1)

	result2 := NewLanguage().Uc(str1, false)
	fmt.Println(result2)

	result3 := NewLanguage().Uc(str2, true)
	fmt.Println(result3)

	result4 := NewLanguage().Uc(str2, false)
	fmt.Println(result4)
}

/**
 * Test Language::isValidBuiltInCode()
 * @dataProvider provideLanguageCodes
 * @covers Language::isValidBuiltInCode
 */
func TestBuiltInCodeValidation(t *testing.T) {
	l := new(Language)

	test.AssetEqual(
		true,
		l.IsValidBuiltInCode("fr"),
		`Two letters, minor case`,
	)

	test.AssetEqual(
		false,
		l.IsValidBuiltInCode("EN"),
		`Two letters, upper case`,
	)

	test.AssetEqual(
		true,
		l.IsValidBuiltInCode("tyv"),
		`Three letters`,
	)

	test.AssetEqual(
		true,
		l.IsValidBuiltInCode("be-tarask"),
		`With dash`,
	)

	test.AssetEqual(
		true,
		l.IsValidBuiltInCode("be-x-old"),
		`With extension (two dashes)`,
	)

	test.AssetEqual(
		false,
		l.IsValidBuiltInCode("be_tarask"),
		`Reject underscores`,
	)
}