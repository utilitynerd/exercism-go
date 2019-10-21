package isogram

import "unicode"

// IsIsogram returns whether a given string is an isogram
func IsIsogram(s string) bool {
	letters := make(map[rune]bool)
	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}
		if _, ok := letters[unicode.ToUpper(r)]; ok {
			return false
		}
		letters[unicode.ToUpper(r)] = true
	}
	return true
}

// IsIsogram2 returns whether a given string is an isogram
func IsIsogram2(s string) bool {
	letters := make(map[rune]struct{})
	for _, r := range s {
		if r == '-' || r == ' ' {
			continue
		}
		if _, ok := letters[unicode.ToUpper(r)]; ok {
			return false
		}
		letters[unicode.ToUpper(r)] = struct{}{}
	}
	return true
}
