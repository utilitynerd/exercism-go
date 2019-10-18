// Package acronym provides the ability to create an acronym from a string
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns an acronym built from its input string
func Abbreviate(s string) string {
	var output strings.Builder
	r, _ := regexp.Compile(`([[:alpha:]])(?:[[:alpha:]]|')*`)
	for _, match := range r.FindAllStringSubmatch(s, -1) {
		output.WriteString(match[1])
	}
	return strings.ToUpper(output.String())
}
