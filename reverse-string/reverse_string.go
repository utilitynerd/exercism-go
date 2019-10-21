// Package reverse provides a function for reversing strings
package reverse

// Reverse returns the input string reversed
func Reverse(s string) string {
	runes := []rune(s)

	for i := len(runes)/2 - 1; i >= 0; i-- {
		opp := len(runes) - 1 - i
		runes[i], runes[opp] = runes[opp], runes[i]
	}

	return string(runes)
}
