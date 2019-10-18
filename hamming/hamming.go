// Package hamming provides a function to calculate the hamming distance of 2 strings
package hamming

import "fmt"

// Distance returns the hamming distance of 2 strings and an optional error
func Distance(a, b string) (int, error) {
	// distance gets initialized as int types zero value, 0
	var distance int
	if len(a) != len(b) {
		return distance, fmt.Errorf("Strings must be the same length")
	}
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
