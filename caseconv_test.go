package caseconv

import (
	"testing"
)

func TestCamelToSnake(t *testing.T) {
	shouldMatch(t, CamelToSnake, "SnakeCase", "snake_case")
	shouldMatch(t, CamelToSnake, "snakeCase", "snake_case")
	shouldMatch(t, CamelToSnake, "ABBR", "abbr")
	shouldMatch(t, CamelToSnake, "ABBRDescription", "abbr_description")
	shouldMatch(t, CamelToSnake, "LongABBRDescription", "long_abbr_description")
}

func shouldMatch(t *testing.T, convFn func(string) string, in, expected string) {
	if res := convFn(in); res != expected {
		t.Errorf("Expected %q to convert to %q, got %q", in, expected, res)
	}
}
