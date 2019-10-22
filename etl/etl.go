//Package etl provides tranformation functions
package etl

import "strings"

// Transform pivots the input, from a map[int][]string to a map[string]int
func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)

	for score, letters := range in {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}

	return out
}
