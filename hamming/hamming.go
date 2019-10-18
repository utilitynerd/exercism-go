// Package hamming provides a function to calculate the hamming distance of 2 strings
package hamming

import (
	"errors"
)

// Distance returns the hamming distance of 2 strings and an optional error
func Distance(a, b string) (int, error) {
	// distance gets initialized as int types zero value, 0
	var distance int
	aRunes, bRunes := []rune(a), []rune(b)

	if len(aRunes) != len(bRunes) {
		return distance, errors.New("Strings must be the same length")
	}
	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			distance++
		}
	}
	return distance, nil
}
