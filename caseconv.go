package caseconv

import (
	"regexp"
	"strings"
)

var snake1 = regexp.MustCompile("([A-Z])([A-Z][a-z])")
var snake2 = regexp.MustCompile("([a-z])([A-Z])")

// CamelToSnake converts CamelCase to snake_case
func CamelToSnake(str string) string {
	return strings.ToLower(snake2.ReplaceAllString(snake1.ReplaceAllString(str, "${1}_${2}"), "${1}_${2}"))
}
