package ch8

// OMG I did this all by myself, this is a backtracking problem!!!
// TODO: compare my solution with the book's
// Actually, the book asked to return a path, I tried to solve the subproblem
// if a path even exists
func robotInAGridHelper(matrix [][]int, row, col int) bool {
	if matrix[0][0] == 1 || matrix[len(matrix)-1][len(matrix)-1] == 1 {
		return false
	}
	if row == len(matrix) {
		return true
	}

	for row < len(matrix) && col < len(matrix[0]) {
		if matrix[row][col] == 1 {
			return false
		}
		// if robotInAGridHelper(matrix, row+1, col) || robotInAGridHelper(matrix, row, col+1) {
		// 	return true
		// }
		// add to path?
		// works just the same, makes sense since we are incrementing col by one
		if robotInAGridHelper(matrix, row+1, col) {
			return true
		}
		// remove from path?
		col++
	}
	return false
}

func RobotInAGrid(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	}
	return robotInAGridHelper(matrix, 0, 0)
}
