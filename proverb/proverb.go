// Package proverb supplies all your proverb needs
package proverb

import "fmt"

// Proverb constructs a deeply meaningful proverb from a slice of strings
func Proverb(rhyme []string) []string {
	inLen := len(rhyme)
	// We can allocate a slice of exactly required length
	output := make([]string, inLen)

	for i := range rhyme {
		if i+1 < inLen {
			// Since proper slice is already allocated, can directly assign values
			// instead of using append
			output[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
		if i+1 == inLen {
			output[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		}
	}
	return output
}
