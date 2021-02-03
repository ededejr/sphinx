package main

import (
	"testing"

	"github.com/fatih/color"
)

func TestSphinx(t *testing.T) {
	cases := []testCase{
		makeBasicTestCase(10, "A", "Token is of expected length", tokenLengthShouldBe(10)),
		makeSeparatorTestCase(20, "A", "-", 4, "Token includes separator", tokenIncludes("-")),
		makeSeparatorTestCase(
			20, "A", "-", 4,
			"Token includes the right number of separators", tokenIncludesNOccurencesOf("-", 3),
		),
		makeSeparatorTestCase(
			20, "A", "-", 4,
			"Number of characters should be (length - (#separators - 1))", tokenLengthWithoutSeparatorShouldBe(17, "-"),
		),
		makeSeparatorTestCase(
			32, "A", "-", 4,
			"Number of characters should be (length - (#separators - 1))", tokenLengthWithoutSeparatorShouldBe(29, "-"),
		),
		makeSeparatorTestCase(
			50, "A", "-", 10,
			"Number of characters should be (length - (#separators - 1))", tokenLengthWithoutSeparatorShouldBe(41, "-"),
		),
	}

	for _, test := range cases {
		printCaseDescription("Case: %s => ", test.description)
		token := generateToken(test.length, test.characters, test.separator, test.chunks)
		printCaseToken(token)
		test.verify(t, token)
		color.HiGreen(" ✔")
	}
}
