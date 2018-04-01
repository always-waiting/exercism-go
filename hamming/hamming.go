package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	var e error
	if len(a) != len(b) {
		e = errors.New("长度不同")
		return -1, e
	}
	byteA := []byte(a)
	byteB := []byte(b)
	var count int
	for index, value := range byteA {
		if value != byteB[index] {
			count++
		}
	}
	return count, e
}
