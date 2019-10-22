// Package accumulate provides a funtion that applies an operation to a collection
package accumulate

// Accumulate applies the given function to each member of in, and returns the results in a new slice
func Accumulate(in []string, f func(string) string) []string {
	output := make([]string, len(in))
	for i, s := range in {
		output[i] = f(s)
	}
	return output
}
