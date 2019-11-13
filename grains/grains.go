// Package grains provides functions to help solve the classic wheat and chessboard problem
// https://en.wikipedia.org/wiki/Wheat_and_chessboard_problem
package grains

import (
	"fmt"
)

// Square retruns the number of grains on a specific chessboard square
func Square(i int) (uint64, error) {
	if i <= 0 || i > 64 {
		return 0, fmt.Errorf("invalid square %d", i)
	}
	return 1 << (i - 1), nil
}

// Total returns the total number of wheat grains on a filled chessboard
func Total() (grains uint64) {
	return 1<<64 - 1
}
