package letter

import "sync"

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

func ConcurrentFrequency(texts []string) FreqMap {
	results := make(chan FreqMap)
	var wg sync.WaitGroup

	for _, s := range texts {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			results <- Frequency(s)
		}(s)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	out := FreqMap{}
	for fm := range results {
		for k := range fm {
			out[k] += fm[k]
		}
	}
	return out
}
