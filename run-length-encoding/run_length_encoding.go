package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// RunLengthEncode uses run length encoding to compress the input string s
func RunLengthEncode(s string) string {
	if len(s) == 0 {
		return ""
	}
	var out string
	var curCount int
	curRune, _ := utf8.DecodeRuneInString(s)
	for _, r := range s {
		if curRune != r {
			out = appendRLE(out, curRune, curCount)
			curRune = r
			curCount = 1
		} else {
			curCount++
		}
	}
	return appendRLE(out, curRune, curCount)
}

func appendRLE(s string, r rune, count int) string {
	if count == 1 {
		s += string(r)
	} else {
		s += fmt.Sprintf("%d%s", count, string(r))
	}
	return s
}

// RunLengthDecode decodes run length encode string s
func RunLengthDecode(s string) string {
	var out string
	var count int
	var digit int
	var err error

	for _, r := range s {
		if unicode.IsDigit(r) {
			if digit, err = strconv.Atoi(string(r)); err != nil {
				panic("invalid input string")
			}

			if count > 0 {
				count = count*10 + digit
			} else {
				count = digit
			}

		} else {
			if count > 0 {
				out += strings.Repeat(string(r), count)
			} else {
				out += string(r)
			}

			count = 0
		}
	}
	return out
}
