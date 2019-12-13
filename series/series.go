package series

func All(n int, s string) []string {
	var series = make([]string, 0)
	for i, j := 0, n; j <= len(s); i, j = i+1, j+1 {
		series = append(series, s[i:j])
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
