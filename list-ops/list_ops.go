package listops

// IntList is a []int with additional methods attached
type IntList []int

type unaryFunc func(int) int
type binFunc func(x, y int) int
type predFunc func(int) bool

// Length returns the length of the InList
func (l IntList) Length() int {
	return len(l)
}

// Reverse returns the reverse of the IntList
func (l IntList) Reverse() IntList {
	out := make(IntList, len(l))
	for i := len(l) - 1; i >= 0; i-- {
		out[len(l)-i-1] = l[i]
	}
	return out
}

// Map returns a new InList with f applied to each member
func (l IntList) Map(f unaryFunc) IntList {
	out := make(IntList, len(l))
	for i, v := range l {
		out[i] = f(v)
	}
	return out
}

// Append returns a new Intist with list appended
func (l IntList) Append(list IntList) IntList {
	out := make(IntList, len(l)+len(list))

	for i, v := range l {
		out[i] = v
	}
	for i, v := range list {
		out[i+len(l)] = v
	}
	return out
}

// Concat returns a new IntList with all the gives lists concatinated
func (l IntList) Concat(lists []IntList) IntList {
	out := make(IntList, len(l))
	for i, v := range l {
		out[i] = v
	}
	if len(lists) > 0 {
		for _, list := range lists {
			for _, v := range list {
				out = append(out, v)
			}
		}
	}
	return out
}

// Foldl fold each element of l into accumulator from the left using f
func (l IntList) Foldl(f binFunc, accum int) int {
	for _, v := range l {
		accum = f(accum, v)
	}
	return accum
}

// Foldr fold each element of l into accumulator from the right using f
func (l IntList) Foldr(f binFunc, accum int) int {
	if len(l) > 0 {
		for i := len(l) - 1; i >= 0; i-- {
			accum = f(l[i], accum)
		}
	}
	return accum
}

// Filter returns a new InList filtered by f
func (l IntList) Filter(f predFunc) IntList {
	out := make(IntList, 0)
	for _, v := range l {
		if f(v) {
			out = append(out, v)
		}
	}
	return out
}
