// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	// "[ -,]+"就会出错!
	reg := regexp.MustCompile("[-, ]+")
	s = strings.Title(s)
	aStr := reg.Split(s, -1)
	aUpperStr := Map(aStr, func(s string) string {
		// 用索引后变为byte类型了
		return string(s[0])
	})
	return strings.Join(aUpperStr, "")
}

func Map(vs []string, f func(string) string) []string {
	mvs := make([]string, len(vs))
	for i, v := range vs {
		mvs[i] = f(v)
	}
	return mvs
}
