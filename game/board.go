package game

type Board [][]int

// GetValueAt safely access a value at x, y
// It returns the value if found, or 0 if not
// present or invalid
func (b Board) GetValueAt(row, col int) int {
	if row >= 0 &&
		col >= 0 &&
		row < len(b) &&
		col < len(b[row]) {
		return b[row][col]
	} else {
		return 0
	}

}

// IsEmpty returns a boolean representing whether the game has completed
// and all cells have died
// TODO - this can be more efficient by storing an array of all live cells
// and only iterating over them instead of the entire board
func (b Board) IsEmpty() bool {
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[row]); col++ {
			if b[row][col] == 1 {
				return false
			}
		}
	}
	return true
}
