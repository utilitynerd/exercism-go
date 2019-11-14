package sieve

func Sieve(limit int) []int {
	n := make(map[int]struct{})
	var s struct{}
	for i := 2; i <= limit+1; i++ {
		n[i] = s
	}
	for i = 2; i <= limit+1; i++ {

	}
	return make([]int, 0)
}
