package core

import "fmt"

type Move struct {
	Row, Col  int
	Quadrant  int
	Direction int
}

func (m Move) Repr() string {
	return fmt.Sprintf("(%v|%v) Q%v R%v", m.Row, m.Col, m.Quadrant, m.Direction)
}

func FindSuccessors(b Board) map[Board]Move {
	moves := b.findMoves()

	found := make(map[Board]Move, 0)
	for _, move := range moves {
		bnew := b.SetAt(move.Row, move.Col).Rotate(move.Quadrant, move.Direction)

		present := false
		for bo := range found {
			if bo.EqualsIgnoreRotation(bnew) {
				present = true
				break
			}
		}

		if !present {
			// we found a new successor board and store the move leading there
			found[bnew] = move
		}
	}

	return found
}

func (b Board) findMoves() []Move {
	moves := make([]Move, 0)

	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if b.Fields[i][j] == 0 {
				for r := 0; r < 4; r++ {
					m1 := Move{Row: i, Col: j, Quadrant: r, Direction: CLOCKWISE}
					m2 := Move{Row: i, Col: j, Quadrant: r, Direction: COUNTERCLOCKWISE}
					moves = append(moves, m1, m2)
				}
			}
		}
	}

	return moves
}
