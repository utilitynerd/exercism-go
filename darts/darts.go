// Package darts provides functions to help score a theoritical dart game
package darts

const (
	innerCirle   float64 = 1
	middleCircle float64 = 5
	outerCircle  float64 = 10
)

// Score calculates the score for a single dart throw
func Score(x, y float64) int {
	if inCirlce(x, y, innerCirle) {
		return 10
	} else if inCirlce(x, y, middleCircle) {
		return 5
	} else if inCirlce(x, y, outerCircle) {
		return 1
	}
	return 0
}

// inCirlce returns whether coordianates x,y are on/in a cirlcle
// of radius centered at 0,0
func inCirlce(x, y, radius float64) bool {
	return (x*x)+(y*y) <= radius*radius
}
