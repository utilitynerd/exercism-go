package wordcount

import (
	"regexp"
	"strings"
)

// Frequency maps a word to the number of times it appears in a string
type Frequency map[string]int

var findWord = regexp.MustCompile(`\w+('\w+)?|\d+`)

// WordCount returns a Frequency, representing a count of words and integers in string s
func WordCount(s string) Frequency {
	s = strings.ToLower(s)
	out := make(Frequency)

	words := findWord.FindAllString(s, -1)

	for _, w := range words {
		if count, ok := out[w]; ok {
			out[w] = count + 1
		} else {
			out[w] = 1
		}
	}
	return out
}
