package isogram

import (
	"sort"
	"strings"
	"unicode"
)

// IsIsogram4 unique byte
func IsIsogram4(input string) bool {
	var inputString = strings.ToLower(input)
	inputString = sortString(inputString)
	for i := range inputString {
		if i < len(inputString)-1 && isLetter(inputString[i]) && inputString[i] == inputString[i+1] {
			return false
		}
	}
	return true
}

func isLetter(s byte) bool {
	if 'a' <= s && s <= 'z' {
		return true
	}
	return false
}

type sortRunes []rune

// implement Len, Swap Less interface
func (s sortRunes) Len() int {
	return len(s)
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// IsIsogram3 function
func IsIsogram3(input string) bool {
	isoMap := make(map[rune]int)
	var runes = []rune(strings.ToLower(input))
	var res bool = true
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			if _, ok := isoMap[runes[i]]; ok {
				res = false
			} else {
				isoMap[runes[i]] = 1
			}

		}
	}
	return res

}
