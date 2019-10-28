// Package diffsquares provides functions that calculate the sum of squares
// and square of sums for sequences from 1 to a given integer
package diffsquares

// SquareOfSumNaive returns the square of the sum of integers from 1 to n
func SquareOfSumNaive(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i
	}
	result *= result
	return result
}

// SquareOfSum returns the square of the sum of integers from 1 to n
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquaresNaive returns the  sum of the squares of  integers from 1 to n
func SumOfSquaresNaive(n int) int {
	var result int
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// SumOfSquares returns the  sum of the squares of  integers from 1 to n
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1) / 6)
}

// Difference returns the differece of the square of the sum and the sum of the squares
// of integers from 1 to n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
