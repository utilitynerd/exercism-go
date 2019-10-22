package protein

import (
	"errors"
)

// ErrStop represents encountering a stop codon
var ErrStop = errors.New("stop codon encountered")

// ErrInvalidBase represents encountering an invalid codon base
var ErrInvalidBase = errors.New("invalid codon base encountered")

// FromCodon return the protein encouded by the input codon
func FromCodon(rna string) (string, error) {

	switch rna {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}
	return "", ErrInvalidBase

}

// FromRNA returns a slice of protein names encoded by the input rna string
func FromRNA(rna string) ([]string, error) {
	protiens := make([]string, 0)

	// Iterate through rna string, 1 codon (3 characters) at a time
	for i := 0; i < len(rna); i += 3 {
		protien, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrInvalidBase {
				return protiens, err
			}
			break
		}
		protiens = append(protiens, protien)
	}

	return protiens, nil
}
