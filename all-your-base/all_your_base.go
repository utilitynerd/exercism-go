package allyourbase

import "math"

import "errors"

func ConvertToBase(inBase int, digits []int, outBase int) ([]int, error) {
	var out []int
	if inBase < 2 {
		return out, errors.New("input base must be >= 2")
	}
	if outBase < 2 {
		return out, errors.New("output base must be >= 2")
	}

	var num int
	for i, pow := len(digits)-1, 0; i >= 0; i-- {
		if digits[i] < 0 || digits[i] >= inBase {
			return out, errors.New("all digits must satisfy 0 <= d < input base")
		}
		num += digits[i] * int(math.Pow(float64(inBase), float64(pow)))
		pow++
	}

	for num > 0 {
		remain := num % outBase
		out = prependInt(out, remain)
		num = num / outBase
	}

	if out == nil {
		out = append(out, 0)
	}

	return out, nil
}

func prependInt(x []int, y int) []int {
	x = append(x, 0)
	copy(x[1:], x)
	x[0] = y
	return x
}
