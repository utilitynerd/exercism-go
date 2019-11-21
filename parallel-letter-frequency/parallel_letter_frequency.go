package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency concurrently counts the frequency on runes in a slice of input texts.
// The results are merged into a single FreqMap
func ConcurrentFrequency(texts []string) FreqMap {
	results := make(chan FreqMap)

	for _, s := range texts {
		go func(s string) {
			fm := Frequency(s)
			results <- fm
		}(s)
	}

	out := make(FreqMap)
	for range texts {
		fm := <-results
		for k, v := range fm {
			out[k] += v
		}
	}
	return out
}
