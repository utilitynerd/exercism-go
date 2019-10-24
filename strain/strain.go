// Package strain contains types that have filtering methods
package strain

// Ints is a []int with filter methods attached
type Ints []int

// Keep returns a new Ints containing values where func f returns true
func (i Ints) Keep(f func(int) bool) Ints {
	var out Ints
	for _, n := range i {
		if f(n) {
			out = append(out, n)
		}
	}
	return out
}

// Discard returns a new Ints containing values where func f returns false
func (i Ints) Discard(f func(int) bool) Ints {
	var out Ints
	for _, n := range i {
		if !f(n) {
			out = append(out, n)
		}
	}
	return out
}

// Strings is a []int with filter methods attached
type Strings []string

// Keep returns a new Strings containing values where func f returns true
func (s Strings) Keep(f func(string) bool) Strings {
	var out Strings
	for _, n := range s {
		if f(n) {
			out = append(out, n)
		}
	}
	return out
}

// Discard returns a new Ints containing values where func f returns false
func (s Strings) Discard(f func(string) bool) Strings {
	var out Strings
	for _, n := range s {
		if !f(n) {
			out = append(out, n)
		}
	}
	return out
}

// Lists is a []struct{a, b, c int} with filtering methods attached
type Lists [][]int

// Keep returns a new Lists containing values where func f returns true
func (l Lists) Keep(f func([]int) bool) Lists {
	var out Lists

	for _, i := range l {
		if f(i) {
			out = append(out, i)
		}
	}
	return out
}

// Discard returns a new Lists containing values where func f returns false
func (l Lists) Discard(f func([]int) bool) Lists {
	var out Lists

	for _, i := range l {
		if !f(i) {
			out = append(out, i)
		}
	}
	return out
}
