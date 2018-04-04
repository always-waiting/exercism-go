package grains

import (
	"errors"
	"math"
)

func Square(n int) (uint64, error) {
	if n > 0 && n < 65 {
		return uint64(math.Pow(float64(2), float64(n-1))), nil
	} else {
		return 0, errors.New("错误输入")
	}
}

func Total() uint64 {
	var sum uint64
	for i := 1; i < 65; i++ {
		number, _ := Square(i)
		sum += number
	}
	return sum
}
