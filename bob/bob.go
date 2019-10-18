// Package bob simulates bob, a teenager
package bob

import (
	"strings"
)

// Hey implements bob having a conversation
func Hey(remark string) string {
	if isQuestion(remark) {
		if hasLetter(remark) && isYelling(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if hasLetter(remark) && isYelling(remark) {
		return "Whoa, chill out!"
	} else if isSilent(remark) {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func hasLetter(s string) bool {
	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			return true
		}
	}
	return false
}

func isQuestion(s string) bool {
	return strings.HasSuffix(strings.TrimSpace(s), "?")
}

func isYelling(s string) bool {
	return strings.ToUpper(s) == s
}

func isSilent(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
