package luhn

import (
	"fmt"
	"regexp"
	"strconv"
	//"strings"
)

func Valid(input string) bool {
	if !inputCheck(input) {
		return false
	}
	fmt.Println("Begin:", input)
	input = trimSpace(input)
	//fmt.Println("End:", input)
	var sum int
	for i, v := range input {
		number, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			doubleNumber := number * 2
			if doubleNumber > 9 {
				doubleNumber -= 9
			}
			sum += doubleNumber
		} else {
			sum += number
		}
	}
	if sum%10 == 0 {
		return true
	} else {
		return false
	}
}

func trimSpace(input string) string {
	reg := regexp.MustCompile(" ")
	ret := reg.ReplaceAllString(input, "")
	return ret
}

func inputCheck(input string) bool {
	switch len(input) {
	case 0:
	case 1:
		return false

	}
	reg := regexp.MustCompile("[^0-9 ]")
	if reg.MatchString(input) {
		return false
	}
	return true
}
