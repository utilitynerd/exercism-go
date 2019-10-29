// Package pangram provides a funtion to check if a string is a pangram
package pangram

import "unicode"

// IsPangram returns true if ascii string s is a case-insensitive pangram
func IsPangram(s string) bool {

	seen := make(map[rune]struct{})

	for _, r := range s {
		r = unicode.ToLower(r)
		if 'a' <= r && r <= 'z' {
			if _, ok := seen[r]; !ok {
				seen[r] = struct{}{}
			}
		}
	}
	if len(seen) == 26 {
		return true
	}
	return false
}
