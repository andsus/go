package alphametics

import (
	"fmt"
	"strings"
)

type alphametic struct {
	firstChrs   [26]bool
	usedDigits  [10]bool
	chrs        []rune
	sum         string
	operands    []string
	runeToDigit [26]int
}

func (a *alphametic) solveAlphaMetic(depth int) bool {
	if depth == len(a.chrs) {
		sumVal, operandsTotal := 0, 0
		for _, r := range a.sum {
			sumVal = 10*sumVal + a.runeToDigit[r-'A']
		}
		for _, operand := range a.operands {
			n := 0
			for _, r := range operand {
				n = 10*n + a.runeToDigit[r-'A']
			}
			operandsTotal += n
		}
		return operandsTotal == sumVal
	}

	chr := a.chrs[depth] - 'A'
	for digit := 0; digit < 10; digit++ {
		if (digit != 0 || !a.firstChrs[chr]) && !a.usedDigits[digit] {
			a.usedDigits[digit] = true
			a.runeToDigit[chr] = digit
			if a.solveAlphaMetic(depth + 1) {
				return true
			}
			a.usedDigits[digit] = false
		}
	}
	return false
}

func Solve(puzzle string) (map[string]int, error) {
	split := strings.Split(strings.ReplaceAll(puzzle, " ", ""), "==")

	a := alphametic{
		chrs: make([]rune, 0, 26),
	}

	a.sum, a.operands = split[1], strings.Split(split[0], "+")
	a.firstChrs[a.sum[0]-'A'] = true
	for _, c := range a.sum {
		a.runeToDigit[c-'A'] = 1
	}

	for _, n := range a.operands {
		a.firstChrs[n[0]-'A'] = true
		for _, c := range n {
			a.runeToDigit[c-'A'] = 1
		}
	}
	for i, v := range a.runeToDigit {
		if v != 0 {
			a.chrs = append(a.chrs, rune(i)+'A')
		}
	}
	if !a.solveAlphaMetic(0) {
		return nil, fmt.Errorf("no solution")
	}

	res := make(map[string]int, len(a.chrs))

	for _, v := range a.chrs {
		res[string(v)] = a.runeToDigit[v-'A']

	}
	return res, nil
	// panic("Please implement the Solve function")
}
