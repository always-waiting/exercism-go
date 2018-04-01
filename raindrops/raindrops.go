package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	var strRet string
	if number%3 == 0 {
		strRet += "Pling"
	}
	if number%5 == 0 {
		strRet += "Plang"
	}
	if number%7 == 0 {
		strRet += "Plong"
	}
	if strRet != "" {
		return strRet
	}
	return strconv.Itoa(number)
}

/*
package raindrops

import "strconv"

var factors = [3]struct {
	Value int
	Word  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(input int) string {
	response := ""
	for _, factor := range factors {
		if input%factor.Value == 0 {
			response += factor.Word
		}
	}
	if len(response) == 0 {
		response += strconv.Itoa(input)
	}
	return response
}
*/
