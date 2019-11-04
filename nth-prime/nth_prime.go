// Package prime provides utilities to work with prime numbers
package prime

// Nth returns the  nth prime number
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	primes := make([]int, 1, n)
	primes[0] = 2
	var isPrime bool
	for i := 3; len(primes) < n; i += 2 {
		isPrime = true
		for _, p := range primes[1:] {
			if i%p == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes[n-1], true
}
