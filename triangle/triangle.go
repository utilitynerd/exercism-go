// Package triangle should have a package comment that summarizes what it's about.
package triangle

import (
	"sort"
)

// Kind represents the kinds of trianges
type Kind int

const (
	// NaT is not a triangle
	NaT = iota
	// Sca is a scalene triange
	Sca
	// Iso is an isosceles triange
	Iso
	// Equ is an equilateral triange
	Equ
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {

	if !checkTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}

// Check if the triagle satifies the triangle inequality rule
func checkTriangle(a, b, c float64) bool {
	// If any side is 0, it's not a valid triange
	if a == 0 || b == 0 || c == 0 {
		return false
	}
	sides := []float64{a, b, c}
	sort.Float64s(sides)
	max := sides[len(sides)-1]
	if max <= sumFloats(sides)-max {
		return true
	}
	return false
}

// sumFloats returns the sum of the members of the input slice
func sumFloats(input []float64) float64 {
	var result float64
	for _, n := range input {
		result = result + n
	}
	return result
}
