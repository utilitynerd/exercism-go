// Package anagram provides a utility function detecting anagrams
package anagram

import (
	"sort"
	"strings"
)

type runeSlice []rune

func (r runeSlice) Len() int           { return len(r) }
func (r runeSlice) Less(i, j int) bool { return r[i] < r[j] }
func (r runeSlice) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

// Detect returns a slice of strings containing all strings in candidate that are anagrams of word
func Detect(word string, candidates []string) []string {
	out := make([]string, 0)

	word = strings.ToLower(word)
	var wordRunes runeSlice = []rune(word)
	sort.Sort(wordRunes)

	for _, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if len(word) == len(candidate) && word != candidateLower {
			var candidateRunes runeSlice = []rune(candidateLower)
			sort.Sort(candidateRunes)
			if string(wordRunes) == string(candidateRunes) {
				out = append(out, candidate)
			}
		}
	}
	return out
}
