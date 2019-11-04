// Package luhn provides a function to validate strings using the Luhn algorithm
package luhn

import (
	"strconv"
	"strings"
)

// Valid expects an ascii encoded string of numbers and spaces
// and returns if the string is valid according to the Luhn algorithm
func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	if len(s) <= 1 {
		return false
	}
	var sum int
	toDouble := false
	for i := len(s) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return false
		}
		if toDouble {
			if digit >= 5 {
				sum += digit*2 - 9
			} else {
				sum += digit * 2
			}
			toDouble = false
		} else {
			sum += digit
			toDouble = true
		}
	}
	return sum%10 == 0
}
