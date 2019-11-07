// Package cryptosquare provides a function to encode strings as a crypto square
package cryptosquare

import (
	"regexp"
	"strings"
)

var nonAplha = regexp.MustCompile("[^a-z0-9]+")

// Encode returns the cryto square encoding of string s
func Encode(s string) string {
	s = strings.ToLower(s)
	s = nonAplha.ReplaceAllString(s, "")
	cols := numCols(s)
	out := make([]string, 0)
	for i := 0; i < cols; i++ {
		var str strings.Builder
		for j := i; j < len(s); j += cols {
			str.WriteByte(s[j])
		}
		if i > 0 && str.Len() < len(out[0]) {
			str.WriteString(" ")
		}
		out = append(out, str.String())
	}

	return strings.Join(out, " ")
}

func numCols(s string) int {
	var square int
	var i int
	for i = 0; square < len(s); i++ {
		square = i * i
	}
	return i - 1
}
