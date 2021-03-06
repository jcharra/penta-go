package ai

import (
	"testing"

	"github.com/jcharra/penta-go/core"
)

func TestFindWinningMove(t *testing.T) {
	b := core.NewBoard()
	var bestMoveWhite, bestMoveBlack, expected core.Move

	// First sample board allowing an immediate win for both colors
	// (even though it couldn't be black's turn, of course)
	b.Fields = [6][6]int{
		[6]int{1, 1, 1, 1, 0, -1},
		[6]int{0, 0, 0, 0, -1, 0},
		[6]int{0, 0, 0, -1, 0, 0},
		[6]int{-1, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b.Turn = core.WHITE

	bestMoveWhite = FindBestMove(b, 1, 0).Move
	if bestMoveWhite.Row != 0 || bestMoveWhite.Col != 4 {
		t.Error("Wrong best move for white: ", bestMoveWhite)
	}

	b.Turn = core.BLACK

	bestMoveBlack = FindBestMove(b, 1, 0).Move
	expected = core.Move{Row: 4, Col: 1, Quadrant: core.LOWERLEFT, Direction: core.CLOCKWISE}
	if bestMoveBlack != expected {
		t.Error("Wrong best move for black: ", bestMoveBlack)
	}

	// Second sample
	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 1, 1, 1},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{-1, -1, -1, 0, 0, -1},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b.Turn = core.WHITE

	bestMoveWhite = FindBestMove(b, 1, 0).Move
	expected = core.Move{Row: 1, Col: 1, Quadrant: core.UPPERLEFT, Direction: core.COUNTERCLOCKWISE}

	if bestMoveWhite != expected {
		t.Error("Wrong best move for white: ", bestMoveWhite)
	}

	b.Turn = core.BLACK

	bestMoveBlack = FindBestMove(b, 1, 0).Move

	expected = core.Move{Row: 4, Col: 5, Quadrant: core.LOWERRIGHT, Direction: core.COUNTERCLOCKWISE}
	if bestMoveBlack != expected {
		t.Error("Wrong best move for black: ", bestMoveBlack)
	}
}

func TestFindMovesDepthOne(t *testing.T) {
	b := core.NewBoard()

	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 1, 0},
		[6]int{1, 1, 1, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, -1, 0, 0, 0},
		[6]int{0, -1, -1, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b.Turn = core.BLACK

	bestMoveBlack := FindBestMove(b, 5, 1)

	// No matter what the actual best move is, WHITE should not be
	// able to win immediately anymore after applying the move

	b = b.SetAt(bestMoveBlack.Move.Row, bestMoveBlack.Move.Col)
	bestMoveWhite := FindBestMove(b, 5, 1)

	if bestMoveWhite.value == winnerValue {
		t.Error("White should not be able to win after Black's move, but actually was. Black moved ", bestMoveBlack)
	}
}

func TestFindMovesDepthTwo(t *testing.T) {
	b := core.NewBoard()

	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 1, 1, 1, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, -1, 0, 0, -1, 0},
		[6]int{0, 0, -1, 0, 0, 0},
	}

	b.Turn = core.WHITE

	bestMoveWhite := FindBestMove(b, 5, 2)

	if bestMoveWhite.value != winnerValue {
		t.Error("White had a forced win, but moved ", bestMoveWhite)
	}
}
