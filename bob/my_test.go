package bob

import "testing"

func TestHasLetter(t *testing.T) {
	cases := map[string]bool{
		"WHAT THE HEL": true,
		"1,2,3":        false,
		".@$@":         false,
	}

	for k, v := range cases {
		if v != hasLetter(k) {
			t.Fatalf("%s got %v, expected %v", k, hasLetter(k), v)
		}
	}
}
