// Package raindrops provides soothing simulated raindrop sounds
package raindrops

import (
	"strconv"
	"strings"
)

// Convert converts an integer to a simulated rain sound
func Convert(i int) string {
	var output strings.Builder
	hasFactor := false
	if i%3 == 0 {
		output.WriteString("Pling")
		hasFactor = true
	}
	if i%5 == 0 {
		output.WriteString("Plang")
		hasFactor = true
	}
	if i%7 == 0 {
		output.WriteString("Plong")
		hasFactor = true
	}
	if hasFactor {
		return output.String()
	}
	return strconv.Itoa(i)
}
