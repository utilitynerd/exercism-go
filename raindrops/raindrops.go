// Package raindrops provides soothing simulated raindrop sounds
package raindrops

import (
	"strconv"
)

// Convert converts an integer to a simulated rain sound
func Convert(i int) string {
	var output string
	if i%3 == 0 {
		output += "Pling"
	}
	if i%5 == 0 {
		output += "Plang"
	}
	if i%7 == 0 {
		output += "Plong"
	}
	if len(output) == 0 {
		return strconv.Itoa(i)
	}
	return output
}
