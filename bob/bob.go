// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.Trim(remark, "\n\r \t")
	question := isQuestion(remark)
	yell := isYell(remark)
	slience := isSlience(remark)
	if question && yell {
		return "Calm down, I know what I'm doing!"
	}
	if question {
		return "Sure."
	}
	if yell {
		return "Whoa, chill out!"
	}
	if slience {
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func isYell(remark string) bool {
	var valid = regexp.MustCompile(`[A-Z]+`)
	if !valid.MatchString(remark) {
		return false
	}
	if strings.ToUpper(remark) == remark {
		return true
	}
	return false
}
func isQuestion(remark string) bool {
	if strings.HasSuffix(remark, "?") {
		return true
	}
	return false
}
func isSlience(remark string) bool {
	if remark == "" {
		return true
	}
	return false
}
