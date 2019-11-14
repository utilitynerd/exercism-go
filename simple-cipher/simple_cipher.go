package cipher

import (
	"regexp"
	"strings"
)

type shift struct {
	distance int
}

// NewCaesar returns a Cipher which implements a Ceaser Cipher with a fixed shift of 3
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift returns a Cipher which implements a Ceaser Cipher with a shift between -25 and 25
func NewShift(distance int) Cipher {
	if distance == 0 || distance < -25 || distance > 25 {
		return nil
	}
	return &shift{distance}
}

func (c shift) Encode(s string) string {
	s = strings.ToLower(s)
	var out strings.Builder
	for _, letter := range s {
		if 'a' <= letter && letter <= 'z' {
			letter = shiftLetter(letter, c.distance)
			out.WriteRune(letter)
		}
	}
	return out.String()
}

func (c shift) Decode(s string) string {
	var out strings.Builder
	for _, letter := range s {
		letter = shiftLetter(letter, c.distance*-1)
		out.WriteRune(letter)
	}
	return out.String()
}

type vigenere struct {
	key string
}

var key = regexp.MustCompile("^a*[b-z]+[a-z]*$")

// NewVigenere returns a Cipher implementing a vigenere ciphere
func NewVigenere(s string) Cipher {
	if !key.MatchString(s) {
		return nil
	}
	return &vigenere{s}
}

func (c vigenere) Encode(s string) string {
	s = strings.ToLower(s)
	var out strings.Builder

	for i, j := 0, 0; i < len(s); i++ {
		letter := rune(s[i])
		if 'a' <= letter && letter <= 'z' {
			shift := int(c.key[j%len(c.key)]) - 97
			letter = shiftLetter(letter, shift)
			out.WriteRune(letter)
			j++
		}
	}

	return out.String()
}

func (c vigenere) Decode(s string) string {
	var out strings.Builder
	for i := 0; i < len(s); i++ {
		shift := int(c.key[i%len(c.key)]) - 97
		letter := shiftLetter(rune(s[i]), shift*-1)
		out.WriteRune(letter)
	}
	return out.String()
}

func shiftLetter(letter rune, distance int) rune {
	letter = rune(int(letter) + distance)
	if letter < 'a' {
		letter += 26
	} else if letter > 'z' {
		letter -= 26
	}
	return letter
}
