package atbash

import (
	"math"
	"regexp"
	"strings"
)

var nonAplha = regexp.MustCompile("[^a-z0-9]+")

func Atbash(s string) string {
	s = strings.ToLower(s)
	s = nonAplha.ReplaceAllString(s, "")

	var ct strings.Builder

	for i := range s {
		if i > 0 && i%5 == 0 {
			ct.WriteString(" ")
		}
		letter := float64(s[i])
		if 97 <= letter && letter <= 122 {
			letter = math.Abs(letter-97-25) + 97
		}
		ct.WriteByte(byte(letter))
	}

	return ct.String()
}
