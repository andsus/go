package grains

import (
	"errors"
)

// Square function returns number of grains from chess board square
func Square(n int) (uint64, error) {
	if n > 64 || n <= 0 {
		return 0, errors.New("n is out of bound")
	}

	// n << x is "n times 2, x times"
	return (1 << uint64(n-1)), nil
}

// Total count of grains in 64 chess board
func Total() uint64 {
	return (1 << uint64(64)) - 1
}
