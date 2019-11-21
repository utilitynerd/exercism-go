package rotationalcipher

import (
	"strings"
)

func RotationalCipher(text string, shift int) string {

	var cipherText strings.Builder

	for _, r := range text {
		switch {
		case 'A' <= r && r <= 'Z':
			r = rune(int(r) + shift%26)
			if r > 'Z' {
				r -= 26
			} 
		case 'a' <= r && r <= 'z':
			r = rune(int(r) + shift%26)
			if r > 'z' {
				r -= 26
			}
		}
		cipherText.WriteRune(r)
	}

	return cipherText.String()
}
