package connect

type Point struct {
	player    byte
	isVisited bool
}

type Board struct {
	row, col int
	points   [][]*Point
}

func (b *Board) IsWinner(row, col int, player byte) bool {
	if (row >= 0 && row < b.row) && (col >= 0 && col < b.col) {

		if point := b.points[row][col]; player == point.player && !point.isVisited {

			point.isVisited = true

			switch player {

			case 'O':
				// Player 'O' wins if the bottom row is reached
				if row == b.row-1 {
					return true
				}
			case 'X':
				// Player 'X' wins if the rightmost column is reached
				if col == b.col-1 {
					return true
				}
			}

			// go through all possible directions for winner
			// omit diagonal: (row+1, col+1), (row-1, col-1)
			isWinner := b.IsWinner(row-1, col, player) ||
				b.IsWinner(row-1, col+1, player) ||
				b.IsWinner(row, col+1, player) ||
				b.IsWinner(row+1, col, player) ||
				b.IsWinner(row+1, col-1, player) ||
				b.IsWinner(row, col-1, player)

			return isWinner
		}
	}
	return false
}

func ResultOf(input []string) (string, error) {

	// CReate the board
	b := &Board{
		row:    len(input),
		col:    len(input[0]),
		points: make([][]*Point, len(input)),
	}

	// Fill the board
	for i := range b.points {
		// Check for valid input
		if len(input[i]) != b.col {
			return "", nil
		}

		// Make correct lenght slices
		b.points[i] = make([]*Point, b.col)
		for j := range b.points[i] {

			switch input[i][j] {

			case 'X', 'O', '.':
				b.points[i][j] = &Point{input[i][j], false}

			default:
				return "", nil

			}
		}
	}

	// Check the winner
	for _, player := range []byte{'O', 'X'} {
		switch player {
		case 'X':
			// Check for rows for player 1 win
			for row := 0; row < b.row; row++ {
				if b.IsWinner(row, 0, player) {
					return string(player), nil
				}
			}
		case 'O':
			// Check for columns for player 2 win
			for col := 0; col < b.col; col++ {
				if b.IsWinner(0, col, player) {
					return string(player), nil
				}
			}
		}
	}
	return "", nil
}
