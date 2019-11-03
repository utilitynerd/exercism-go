package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
)

type intSlice []int

func makeintSlice(s string) (intSlice, error) {
	iSlice := make(intSlice, 0, len(s))
	var digit int
	var err error
	for _, r := range s {
		if digit, err = strconv.Atoi(string(r)); err != nil {
			return nil, fmt.Errorf("invalid character: %c", r)
		}
		iSlice = append(iSlice, digit)
	}
	return iSlice, nil
}

func (iSlice intSlice) multiplySeries(start int, span int) (product int64, err error) {
	if start < 0 || start+span > len(iSlice) {
		return 0, fmt.Errorf("error start %v, span %v", start, span)
	}
	for i := start; i < start+span; i++ {
		if i == start {
			product = int64(iSlice[i])
		} else {
			product *= int64(iSlice[i])
		}
	}
	return product, nil
}

// LargestSeriesProduct returns the largest product of a series of span numbers
// in string of ascii digits s
func LargestSeriesProduct(s string, span int) (largest int64, err error) {
	if span > len(s) || span < 0 {
		return -1, errors.New("invalid span")
	}
	if span == 0 {
		return 1, nil
	}

	digits, err := makeintSlice(s)
	if err != nil {
		return -1, err
	}

	for i := 0; i+span <= len(s); i++ {
		product, err := digits.multiplySeries(i, span)
		if err != nil {
			return -1, err
		}
		if product > largest {
			largest = product
		}
	}
	return largest, nil
}
