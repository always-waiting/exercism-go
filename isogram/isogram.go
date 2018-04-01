package isogram

import (
	"regexp"
	"strings"
)

func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	reg := regexp.MustCompile("[a-z]")
	for index, value := range input {
		if !reg.MatchString(string(value)) {
			continue
		}
		for i := index + 1; i < len(input); i++ {
			if string(input[i]) == string(value) {
				return false
			}
		}
	}
	return true
}
