package ai

import (
	"testing"

	"github.com/jcharra/pentago/core"
)

func TestFindWinningMove(t *testing.T) {
	b := core.NewBoard()
	var bestMoveWhite, bestMoveBlack, expected core.Move

	// First sample board allowing an immediate win for both colors
	// (even though it couldn't be black's turn, of course)
	b.Fields = [6][6]int{
		[6]int{1, 1, 1, 1, 0, 2},
		[6]int{0, 0, 0, 0, 2, 0},
		[6]int{0, 0, 0, 2, 0, 0},
		[6]int{2, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b.Turn = core.WHITE

	bestMoveWhite = FindBestMove(b, 1, 0)
	if bestMoveWhite.Row != 0 || bestMoveWhite.Col != 4 {
		t.Error("Wrong best move for white: ", bestMoveWhite)
	}

	b.Turn = core.BLACK

	bestMoveBlack = FindBestMove(b, 1, 0)
	expected = core.Move{Row: 4, Col: 1, Quadrant: core.LOWERLEFT, Direction: core.CLOCKWISE}
	if bestMoveBlack != expected {
		t.Error("Wrong best move for black: ", bestMoveBlack)
	}

	// Second sample
	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 1, 1, 1},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{2, 2, 2, 0, 0, 2},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b.Turn = core.WHITE

	bestMoveWhite = FindBestMove(b, 1, 0)
	expected = core.Move{Row: 1, Col: 1, Quadrant: core.UPPERLEFT, Direction: core.COUNTERCLOCKWISE}

	if bestMoveWhite != expected {
		t.Error("Wrong best move for white: ", bestMoveWhite)
	}

	b.Turn = core.BLACK

	bestMoveBlack = FindBestMove(b, 1, 0)

	expected = core.Move{Row: 4, Col: 5, Quadrant: core.LOWERRIGHT, Direction: core.COUNTERCLOCKWISE}
	if bestMoveBlack != expected {
		t.Error("Wrong best move for black: ", bestMoveBlack)
	}
}