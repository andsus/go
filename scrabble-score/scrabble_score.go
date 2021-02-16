package scrabble

import "strings"

//Score scrabble word - Given a word, compute the Scrabble score for that word.
func Score(input string) int {

	var score int
	upCaseInput := strings.ToUpper(input)
	for i := range upCaseInput {

		if upCaseInput[i] >= 'A' && upCaseInput[i] <= 'Z' {
			score += lookupLetterValue(upCaseInput[i])
		}
	}

	return score

}

func lookupLetterValue(character byte) int {
	switch character {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}

}
