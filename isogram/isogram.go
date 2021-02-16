package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram function
func IsIsogram(input string) bool {
	isoMap := make(map[rune]bool)

	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			if _, ok := isoMap[r]; ok {
				return false
			}
			isoMap[r] = false

		}
	}

	return true
}
