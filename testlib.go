package main

import (
	"strings"
	"testing"

	"github.com/fatih/color"
)

type TestFunc func(*testing.T, string)

type testCase struct {
	length      int
	characters  string
	separator   string
	chunks      int
	description string
	verify      TestFunc
}

func makeBasicTestCase(
	length int,
	characters string,
	description string,
	verify func(*testing.T, string)) testCase {
	return testCase{length, characters, "", 0, description, verify}
}

func makeSeparatorTestCase(
	length int,
	characters string,
	separator string,
	chunks int,
	description string,
	verify func(*testing.T, string)) testCase {
	return testCase{length, characters, separator, chunks, description, verify}
}

func verifyLength(token string, length int) bool {
	return len(token) == length
}

func checkOccurence(token string, substr string) bool {
	return strings.Contains(token, substr)
}

func hasAtLeastNOccurences(token string, substr string, occurences int) bool {
	return strings.Count(token, substr) >= occurences
}

func tokenLengthShouldBe(length int) TestFunc {
	return func(t *testing.T, token string) {
		if !verifyLength(token, length) {
			t.Fatalf("%s is not of length %d", token, length)
		}
	}
}

func tokenLengthWithoutSeparatorShouldBe(length int, separator string) TestFunc {
	return func(t *testing.T, token string) {
		if !verifyLength(strings.ReplaceAll(token, separator, ""), length) {
			t.Fatalf("%s is not of length %d without separators", token, length)
		}
	}
}

func tokenIncludes(substr string) TestFunc {
	return func(t *testing.T, token string) {
		if !checkOccurence(token, substr) {
			t.Fatalf("%s does not contain %s", token, substr)
		}
	}
}

func tokenIncludesNOccurencesOf(substr string, occurences int) TestFunc {
	return func(t *testing.T, token string) {
		if !checkOccurence(token, substr) || !hasAtLeastNOccurences(token, substr, occurences) {
			t.Fatalf("%s does not contain %s", token, substr)
		}
	}
}

func printCaseDescription(msg string, description string) {
	c := color.New(color.FgHiBlack)
	c.Printf(msg, description)
}

func printCaseToken(token string) {
	c := color.New(color.FgHiMagenta)
	c.Printf(token)
}
