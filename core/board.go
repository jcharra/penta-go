package core

const (
	WHITE = 1
	BLACK = -1
	DRAW  = 2
)

const (
	CLOCKWISE        = iota
	COUNTERCLOCKWISE = iota
)

const (
	UPPERLEFT  = iota
	UPPERRIGHT = iota
	LOWERLEFT  = iota
	LOWERRIGHT = iota
)

type Board struct {
	Turn   int
	Fields [6][6]int
}

func NewBoard() Board {
	var rows [6][6]int
	return Board{Fields: rows, Turn: WHITE}
}

func (b Board) Repr() string {
	ch := map[int]string{0: "_", 1: "O", -1: "X"}

	s := ""
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			s += ch[b.Fields[i][j]] + " "
		}
		s += "\n"
	}
	return s
}

func (b Board) Rotate(quadrant, direction int) Board {
	b2 := b.Copy()

	var offX, offY int
	switch quadrant {
	case UPPERLEFT:
		offX, offY = 0, 0
	case UPPERRIGHT:
		offX, offY = 3, 0
	case LOWERLEFT:
		offX, offY = 0, 3
	case LOWERRIGHT:
		offX, offY = 3, 3
	}

	if direction == CLOCKWISE {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				b2.Fields[offY+i][offX+j] = b.Fields[offY+2-j][offX+i]
			}
		}
	} else if direction == COUNTERCLOCKWISE {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				b2.Fields[offY+2-j][offX+i] = b.Fields[offY+i][offX+j]
			}
		}
	}
	return b2
}

func (b Board) Equals(b2 Board) bool {
	return b.Fields == b2.Fields
}

func (b Board) EqualsIgnoreRotation(b2 Board) bool {
	return b.equalsRot(b2, 0) || b.equalsRot(b2, 1) || b.equalsRot(b2, 2) || b.equalsRot(b2, 3)
}

/*
Returns whether board b2 equals b after rotation by 90*rotDegree degrees
*/
func (b Board) equalsRot(b2 Board, rotDegree int) bool {
	if rotDegree == 0 {
		return b.Equals(b2)
	} else if rotDegree == 1 {
		for rowIdx, row := range b.Fields {
			for colIdx, val := range row {
				if b2.Fields[colIdx][5-rowIdx] != val {
					return false
				}
			}
		}
	} else if rotDegree == 2 {
		for rowIdx, row := range b.Fields {
			for colIdx, val := range row {
				if b2.Fields[5-rowIdx][5-colIdx] != val {
					return false
				}
			}
		}
	} else if rotDegree == 3 {
		for rowIdx, row := range b.Fields {
			for colIdx, val := range row {
				if b2.Fields[5-colIdx][rowIdx] != val {
					return false
				}
			}
		}
	}

	return true
}

func (b Board) Copy() Board {
	bnew := NewBoard()
	bnew.Turn = b.Turn
	for rowIdx, row := range b.Fields {
		for colIdx, val := range row {
			bnew.Fields[rowIdx][colIdx] = val
		}
	}
	return bnew
}

func (b Board) SetAt(row, col int) Board {
	bnew := b.Copy()

	bnew.Fields[row][col] = b.Turn

	if b.Turn == WHITE {
		bnew.Turn = BLACK
	} else {
		bnew.Turn = WHITE
	}

	return bnew
}

func (b Board) Winner() int {
	// we need to collect all winning position, since multiple
	// may occur, even from both players, which would draw the game.
	wins := make(map[int]bool, 0)

	// Check row winner
	for _, row := range b.Fields {
		w := checkWinner(row)
		if w != 0 {
			wins[w] = true
		}
	}

	// Check column winner and get diagonals
	var diag1, diag2 [6]int
	for i := 0; i < 6; i++ {
		col := [6]int{b.Fields[0][i],
			b.Fields[1][i],
			b.Fields[2][i],
			b.Fields[3][i],
			b.Fields[4][i],
			b.Fields[5][i],
		}
		w := checkWinner(col)

		if w != 0 {
			wins[w] = true
		}

		// fill diagonals
		diag1[i] = b.Fields[i][i]
		diag2[i] = b.Fields[5-i][i]
	}

	// Check winner on 6er-diagonals
	for _, diag := range [][6]int{diag1, diag2} {
		wd := checkWinner(diag)
		if wd != 0 {
			wins[wd] = true
		}
	}

	// Small diagonals (5 Fields)
	var sd1, sd2, sd3, sd4 [5]int
	for i := 0; i < 5; i++ {
		sd1[i] = b.Fields[i][i+1]
		sd2[i] = b.Fields[i+1][i]
		sd3[i] = b.Fields[5-i][i+1]
		sd4[i] = b.Fields[4-i][i]
	}

	for _, sd := range [][5]int{sd1, sd2, sd3, sd4} {
		if sd[0] != 0 && allEqual(sd[0], sd[1:5]) {
			wins[sd[0]] = true
		}
	}

	if wins[WHITE] && wins[BLACK] {
		return DRAW
	} else if wins[WHITE] {
		return WHITE
	} else if wins[BLACK] {
		return BLACK
	}

	// Game is not decided yet, as long as there is still at least one field unoccupied
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if b.Fields[i][j] == 0 {
				return 0
			}
		}
	}

	// All fields are occupied => draw
	return DRAW
}

func checkWinner(arr [6]int) int {
	cand := arr[1] // only this color can win the row
	if arr[0] == cand {
		if allEqual(cand, arr[2:5]) {
			return cand
		}
	} else {
		if allEqual(cand, arr[2:6]) {
			return cand
		}
	}
	return 0
}

func allEqual(val int, arr []int) bool {
	for i := range arr {
		if arr[i] != val {
			return false
		}
	}
	return true
}
