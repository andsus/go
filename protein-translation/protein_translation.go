package protein

import "errors"

// ErrStop is thrown when a STOP Codon encountered
var ErrStop = errors.New("Stop")

// ErrInvalidBase when a base RNA is invalid
var ErrInvalidBase = errors.New("Invalid base")

var codonMap map[string]string = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// FromCodon convert RNA into codon
func FromCodon(input string) (string, error) {

	if value, ok := codonMap[input]; ok {
		if value == "STOP" {
			return "", ErrStop
		}

		return value, nil
	}
	return "", ErrInvalidBase
}

// FromRNA translates polypetides to a sequence of proteins
// stops if a STOP codon is read
func FromRNA(in string) (rna []string, err error) {

	for i := 0; i < len(in); i++ {
		if (i+1)%3 == 0 {
			codon := in[i-2 : i+1]
			cur, err := FromCodon(codon)
			if err != nil {

				if err == ErrStop {
					return rna, nil
				}
				return rna, err
			}

			rna = append(rna, cur)
		}
	}
	return rna, nil
}
