package spiralmatrix

/*
Let us notice one clue property about our spiral matrix: first we need to go to the right and rotate clockwise 90 degrees, then we go down and again when we reached bottom, we rotate 90 degrees clockwise and so on. So, all we need to do is to rotate 90 degrees clockwise when we need:

    When we reached border of our matrix
    When we reached cell which is already filled.

Let x, y be coordinates on our grid and dx, dy is current direction we need to move. In geometrical sense, rotate by 90 degrees clockwise is written as dx, dy = -dy, dx.

Note, that matrix[y][x] is cell with coordinates (x,y), which is not completely obvious.

Complexity: time complexity is O(n^2), we process each element once. Space complexity is O(n^2) as well.
*/
func SpiralMatrix(size int) [][]int {
	matrix := [][]int{}
	for row := 0; row < size; row++ {
		matrix = append(matrix, make([]int, size))
	}
	row, col := 0, 0
	dr, dc := 0, 1
	for i := 1; i <= size*size; i++ {
		matrix[row][col] = i
		if row+dr < 0 || size <= row+dr || col+dc < 0 ||
			size <= col+dc || matrix[row+dr][col+dc] != 0 {
			dr, dc = dc, -dr // rotate -90
		}
		row, col = row+dr, col+dc
	}
	return matrix
}
