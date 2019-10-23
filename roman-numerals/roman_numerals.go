package romannumerals

import (
	"fmt"
	"sort"
)

var numberToRoman = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// ToRomanNumeral converts a decimal integer to its roman numeral representation
func ToRomanNumeral(num int) (string, error) {

	if num <= 0 || num > 3000 {
		return "", fmt.Errorf("invalid number %d", num)
	}

	// To store the keys in slice in sorted order
	var keys []int
	for k := range numberToRoman {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var out string
	for _, key := range keys {
		for num >= key {
			out += numberToRoman[key]
			num -= key
		}
	}

	return out, nil
}
