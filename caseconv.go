package caseconv

import (
	"regexp"
	"strings"
)

/*
CamelCase to snake_case
*/
func CamelToSnake(str string) string {
	snake1 := regexp.MustCompile("([A-Z])([A-Z][a-z])")
	snake2 := regexp.MustCompile("([a-z])([A-Z])")
	return strings.ToLower(snake2.ReplaceAllString(snake1.ReplaceAllString(str, "${1}_${2}"), "${1}_${2}"))
}
