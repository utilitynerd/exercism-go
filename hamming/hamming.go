// Package hamming provides a function to calculate the hamming distance of 2 strings
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance returns the hamming distance of 2 strings and an optional error
func Distance(a, b string) (int, error) {
	// distance gets initialized as int types zero value, 0
	var distance int
	lenA, lenB := utf8.RuneCountInString(a), utf8.RuneCountInString(b)
	if lenA != lenB {
		return distance, errors.New("Strings must be the same length")
	}
	for i1, i2, runes := 0, 0, 0; runes < lenA; runes++ {
		r1, width1 := utf8.DecodeRuneInString(a[i1:])
		r2, width2 := utf8.DecodeRuneInString(b[i2:])
		if r1 != r2 {
			distance++
		}
		i1 += width1
		i2 += width2
	}

	return distance, nil
}
