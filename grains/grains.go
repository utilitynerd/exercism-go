// Package grains provides functions to help solve the classic wheat and chessboard problem
// https://en.wikipedia.org/wiki/Wheat_and_chessboard_problem
package grains

import (
	"fmt"
	"math"
)

// Square retruns the number of grains on a specific chessboard square
func Square(i int) (grains uint64, err error) {
	if i <= 0 || i > 64 {
		return 0, fmt.Errorf("invalid square %d", i)
	}
	grains = uint64(math.Exp2(float64(i - 1)))
	return grains, nil
}

// Square2 retruns the number of grains on a specific chessboard square
func Square2(i int) (grains uint64, err error) {
	if i <= 0 || i > 64 {
		return 0, fmt.Errorf("invalid square %d", i)
	}
	grains = uint64(1 << (i - 1))
	return grains, nil
}

// Total returns the total number of wheat grains on a filled chessboard
func Total() (grains uint64) {
	return math.MaxUint64
}
