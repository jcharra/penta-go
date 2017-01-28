package pentago

import "testing"

func TestFindWinningMove(t *testing.T) {
	b := NewBoard()
	var bestMoveWhite, bestMoveBlack, expected Move

	// First sample board allowing an immediate win for both colors
	// (even though it couldn't be black's turn, of course)
	b.fields = [6][6]int{
		[6]int{1, 1, 1, 1, 0, 2},
		[6]int{0, 0, 0, 0, 2, 0},
		[6]int{0, 0, 0, 2, 0, 0},
		[6]int{2, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b.Turn = WHITE

	bestMoveWhite = FindBestMove(b, 0, 0)
	if bestMoveWhite.row != 0 || bestMoveWhite.col != 4 {
		t.Error("Wrong best move for white: ", bestMoveWhite)
	}

	b.Turn = BLACK

	bestMoveBlack = FindBestMove(b, 0, 0)
	expected = Move{row: 4, col: 1, quadrant: LOWERLEFT, direction: CLOCKWISE}
	if bestMoveBlack != expected {
		t.Error("Wrong best move for black: ", bestMoveBlack)
	}

	// Second sample
	b.fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 1, 1, 1},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{2, 2, 2, 0, 0, 2},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b.Turn = WHITE

	bestMoveWhite = FindBestMove(b, 0, 0)
	expected = Move{row: 1, col: 1, quadrant: UPPERLEFT, direction: COUNTERCLOCKWISE}

	if bestMoveWhite != expected {
		t.Error("Wrong best move for white: ", bestMoveWhite)
	}

	b.Turn = BLACK

	bestMoveBlack = FindBestMove(b, 0, 0)

	expected = Move{row: 4, col: 5, quadrant: LOWERRIGHT, direction: COUNTERCLOCKWISE}
	if bestMoveBlack != expected {
		t.Error("Wrong best move for black: ", bestMoveBlack)
	}
}
