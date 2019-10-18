// Package hamming provides a function to calculate the hamming distance of 2 strings
package hamming

import "errors"

// Distance returns the hamming distance of 2 strings and an optional error
func Distance(a, b string) (int, error) {
	aRunes, bRunes := []rune(a), []rune(b)

	if len(aRunes) != len(bRunes) {
		return 0, errors.New("strings must be the same length")
	}

	distance := 0
	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			distance++
		}
	}
	return distance, nil
}
