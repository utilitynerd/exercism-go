package collatzconjecture

import "fmt"

// CollatzConjecture returns the number of steps it takes to reach 1 using the Collatz Conjecture
func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, fmt.Errorf("input must be greater than 0: %d", n)
	}

	var steps int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		steps++
	}
	return steps, nil
}
