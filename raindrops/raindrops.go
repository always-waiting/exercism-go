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
