package hamming

import (
	"errors"
)

//Distance returns (int,error)
func Distance(a, b string) (int, error) {
	// create rune slice first
	ar, br := []rune(a), []rune(b)
	totalDiff := 0

	if len(ar) != len(br) {
		return 0, errors.New("erros in distance() function input DNA string lengh differ")
	}

	for i := range ar {
		if ar[i] != br[i] {
			totalDiff++
		}
	}

	return totalDiff, nil
}
