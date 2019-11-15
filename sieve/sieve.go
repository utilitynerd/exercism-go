package sieve

import "sort"

// Sieve implements the sieve of eratosthenes to return an []int of all primes up to limit
func Sieve(limit int) []int {
	if limit < 2 {
		return make([]int, 0)
	}
	candidates := make(map[int]struct{})
	var s struct{}
	for i := 2; i <= limit; i++ {
		candidates[i] = s
	}

	for i := 2; i <= limit; i++ {
		for j := i * 2; j <= limit; j += i {
			delete(candidates, j)
		}
	}

	primes := make([]int, len(candidates))
	i := 0
	for k := range candidates {
		primes[i] = k
		i++
	}
	sort.Ints(primes)
	return primes
}
