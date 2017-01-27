package pentago

import (
	"errors"
	"fmt"
)

type Board struct {
	WhitesTurn bool
	fields     [6][6]int
}

func NewBoard() Board {
	var rows [6][6]int
	return Board{fields: rows, WhitesTurn: true}
}

func (b Board) Repr() string {
	s := ""
	for i := 0; i < 6; i++ {
		s += fmt.Sprintf("\n%v", b.fields[i])
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
				b2.fields[offX+i][offY+j] = b.fields[offX+2-j][offY+i]
			}
		}
	} else if direction == COUNTERCLOCKWISE {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				b2.fields[offX+2-j][offY+i] = b.fields[offX+i][offY+j]
			}
		}
	}
	return b2
}

func (b Board) Equals(b2 Board) bool {
	return b.fields == b2.fields
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
		for rowIdx, row := range b.fields {
			for colIdx, val := range row {
				if b2.fields[colIdx][5-rowIdx] != val {
					return false
				}
			}
		}
	} else if rotDegree == 2 {
		for rowIdx, row := range b.fields {
			for colIdx, val := range row {
				if b2.fields[5-rowIdx][5-colIdx] != val {
					return false
				}
			}
		}
	} else if rotDegree == 3 {
		for rowIdx, row := range b.fields {
			for colIdx, val := range row {
				if b2.fields[5-colIdx][rowIdx] != val {
					return false
				}
			}
		}
	}

	return true
}

func (b Board) Copy() Board {
	bnew := NewBoard()
	bnew.WhitesTurn = b.WhitesTurn
	for rowIdx, row := range b.fields {
		for colIdx, val := range row {
			bnew.fields[rowIdx][colIdx] = val
		}
	}
	return bnew
}

func (b Board) SetAt(row, col int) (Board, error) {
	if b.fields[row][col] != 0 {
		return b, errors.New("Field is occupied")
	}

	bnew := b.Copy()

	if b.WhitesTurn {
		bnew.fields[row][col] = WHITE
	} else {
		bnew.fields[row][col] = BLACK
	}

	bnew.WhitesTurn = !b.WhitesTurn
	return bnew, nil
}
