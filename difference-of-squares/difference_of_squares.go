package diffsquares

func SquareOfSums(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	a1 := SquareOfSums(n)
	a2 := SumOfSquares(n)
	return a1 - a2
}
