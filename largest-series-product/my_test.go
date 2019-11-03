package lsproduct

import (
	"reflect"
	"testing"
)

func TestMakeIntSlice(t *testing.T) {
	i, err := makeintSlice("1234")
	if err != nil {
		t.Errorf("input: 1234, output %v", i)
	}
	if reflect.DeepEqual(i, []int{1, 2, 3, 4}) {
		t.Errorf("expected []int{1,2,3,4}, got: %v", i)
	}
}

var cases = []struct {
	start int
	span  int
	sum   int64
}{
	{1, 2, 6},
	{0, 2, 2},
	{0, 3, 6},
	{1, 3, 24},
	{0, 4, 24},
}

func TestMultiplySeries(t *testing.T) {
	i, err := makeintSlice("1234")
	if err != nil {
		t.Errorf("input: 1234, output %v", i)
	}
	for _, test := range cases {

		sum, err := i.multiplySeries(test.start, test.span)
		if err != nil {
			t.Errorf("%v", err)
		}
		if sum != test.sum {
			t.Errorf("expected: %v, got: %v", test.sum, sum)
		}
	}

}
