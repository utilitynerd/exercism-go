// Package scrabble provides utilities to aid scoring scrabble
package scrabble

import (
	"fmt"
	"unicode"
)

var letterScores = map[rune]int{
	'A': 1,
	'B': 3,
	'C': 3,
	'D': 2,
	'E': 1,
	'F': 4,
	'G': 2,
	'H': 4,
	'I': 1,
	'J': 8,
	'K': 5,
	'L': 1,
	'M': 3,
	'N': 1,
	'O': 1,
	'P': 3,
	'Q': 10,
	'R': 1,
	'S': 1,
	'T': 1,
	'U': 1,
	'V': 4,
	'W': 4,
	'X': 8,
	'Y': 4,
	'Z': 10,
}

// Score calculate the scrabble score for a given word
func Score(s string) int {
	score := 0
	for _, r := range s {
		value, ok := letterScores[unicode.ToUpper(r)]
		if !ok {
			panic(fmt.Sprintf("invalid scrabble letter: %c", r))
		}
		score += value
	}
	return score
}

// ScoreCase calculates the scrabble score for a give word
func ScoreCase(s string) int {
	score := 0
	for _, r := range s {
		switch unicode.ToUpper(r) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score++
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		default:
			panic(fmt.Sprintf("invalid scrabble letter: %c", r))
		}
	}
	return score
}
