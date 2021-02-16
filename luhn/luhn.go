package luhn

import (
	"strings"
	"unicode"
)

// Valid is single loop Luhn verification
func Valid(s string) bool {
	// remove whitespace
	s = strings.ReplaceAll(s, " ", "")

	if len(s) < 2 {
		return false
	}

	sum := 0
	// if even len compute first digit sequence
	parity := len(s)%2 == 0

	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}

		digit := int(r - '0')
		if parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		parity = !parity // flip flop

	}
	return sum%10 == 0

}
