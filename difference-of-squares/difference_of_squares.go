package diffsquares

// SquareOfSum (n int) int
func SquareOfSum(n int) int {
	square := (n * (n + 1) / 2)
	return square * square

}

// SumOfSquares (n int) int
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference (n int) int
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
