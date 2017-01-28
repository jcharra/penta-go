package pentago

import "fmt"

type Move struct {
	row, col  int
	quadrant  int
	direction int
}

func FindSuccessors(b Board) map[Board]Move {
	moves := b.findMoves()

	fmt.Printf("Checking %v moves", len(moves))

	found := make(map[Board]Move, 0)
	for _, move := range moves {
		bnew := b.SetAt(move.row, move.col).Rotate(move.quadrant, move.direction)

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
			if b.fields[i][j] == 0 {
				for r := 0; r < 4; r++ {
					m1 := Move{row: i, col: j, quadrant: r, direction: CLOCKWISE}
					m2 := Move{row: i, col: j, quadrant: r, direction: COUNTERCLOCKWISE}
					moves = append(moves, m1, m2)
				}
			}
		}
	}

	return moves
}
