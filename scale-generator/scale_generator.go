package scale

import (
	"fmt"
	"strings"
)

type notes []string

var useSharps = notes{"C", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#", "a"}
var useFlats = notes{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}
var chromaticSharps = notes{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var chromaticFlats = notes{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

func (n notes) pos(target string) (int, error) {
	for i, note := range n {
		if note == target {
			return i, nil
		}
	}
	return -1, fmt.Errorf("note not found: %s", target)
}

var intervalsToSteps = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

// Scale returns the string representaion of a diatonic or chromatic scale startin on tonic and following the interval pattern
func Scale(tonic string, intervals string) []string {
	var n notes
	var err error

	// Determine which notes to use (sharps or flats)
	if n, err = getNotes(tonic); err != nil {
		panic(fmt.Sprintf("invalid note: %s", tonic))
	}

	// find index of tonic in notes
	var start int
	if start, err = n.pos(strings.Title(tonic)); err != nil {
		panic(fmt.Sprintf("invalid note: %s", tonic))
	}

	// The scale start with the tonic
	var scale = []string{n[start]}

	// if interval is "", then generate a chromatic scale
	if intervals == "" {
		for i := start + 1; i%len(n) != start; i++ {
			scale = append(scale, n[i%len(n)])
		}
	} else {
		i := start
		for _, interval := range intervals[:len(intervals)-1] {
			i = i + intervalsToSteps[interval]
			scale = append(scale, n[i%len(n)])
		}
	}

	return scale
}

func getNotes(tonic string) (notes, error) {
	// Determine whether to use flats or sharps
	var n notes
	if _, err := useFlats.pos(tonic); err == nil {
		n = chromaticFlats
	} else if _, err := useSharps.pos(tonic); err == nil {
		n = chromaticSharps
	} else {
		return nil, fmt.Errorf("invalid tonic: %s", tonic)
	}
	return n, nil
}
