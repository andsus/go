package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix stores int in [row,col]
type Matrix [][]int

// New creates a matrix from a string.
func New(input string) (*Matrix, error) {
	var res Matrix

	var cols int

	for i, line := range strings.Split(input, "\n") {
		var row []int
		line = strings.TrimSpace(line)

		for _, strNum := range strings.Fields(line) {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				return nil, errors.New("error string converting number")
			}
			row = append(row, num)
		}
		if i != 0 && len(row) != cols {
			return nil, errors.New("rows with different number of columns")
		}
		cols = len(row)
		res = append(res, row)
	}

	return &res, nil
}

// Rows gets the matrix represented in rows.
func (m Matrix) Rows() [][]int {
	nRows := len(m)

	if nRows == 0 {
		return Matrix{}
	}

	res := make([][]int, nRows)

	for i, row := range m {
		res[i] = append(res[i], row...)
	}

	return res
}

// Cols gets the matrix represented in cols.
func (m Matrix) Cols() [][]int {

	nRows := len(m)

	if nRows == 0 {
		return [][]int{}
	}

	nCols := len(m[0])
	res := make([][]int, nCols)

	for _, row := range m {
		for j, val := range row {
			res[j] = append(res[j], val)
		}
	}

	return res
}

// Set sets the value of the matrix at point row, col.
func (m Matrix) Set(r, c, val int) bool {

	nRows := len(m)

	if r >= nRows || r < 0 {
		return false
	}

	nCols := len(m[0])

	if c >= nCols || c < 0 {
		return false
	}

	m[r][c] = val
	return true
}
