// Package twofer is an implementaion of https://en.wikipedia.org/wiki/Two-fer
package twofer

import "fmt"

// ShareWith returns the string "One for <name>, one for me." if name is given,
// otherwise it returns the string "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
