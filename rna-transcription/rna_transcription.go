package strand

import (
	"fmt"
	"strings"
)

var dnaToRna = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA coverts a string of DNA nucliatides to a string of RNA nucliatides
func ToRNA(dna string) string {
	var output strings.Builder
	for _, c := range dna {
		if rna, ok := dnaToRna[c]; ok {
			output.WriteRune(rna)
		} else {
			panic(fmt.Sprintf("Invalid DNA nucliatides: %c", c))
		}
	}
	return output.String()
}
