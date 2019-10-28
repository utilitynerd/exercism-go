// Package summultiples provices the SumMultiples function, which return sthe sum of all the unique multiples of particular numbers up to but not including that number.
package summultiples

// SumMultiples returns the sum of all the unique multiples of particular numbers up to but not including that number.
func SumMultiples(n int, divisors ...int) int {
	var sum int
	for i := 1; i < n; i++ {
		for _, div := range divisors {
			if div > 0 && i%div == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
