package core

import "testing"

func TestBoard(t *testing.T) {
	b := NewBoard()
	if b.Fields[5][5] != 0 {
		t.Error("Nope")
	}
}

func TestSetAt(t *testing.T) {
	b := NewBoard()

	if b.Turn != WHITE {
		t.Error("Must be white's turn")
	}

	b = b.SetAt(5, 5)
	if b.Fields[5][5] != WHITE {
		t.Error("Incorrect color at 5|5:", b.Fields[5][5])
	}

	if b.Turn == WHITE {
		t.Error("Must be black's turn")
	}

	b = b.SetAt(4, 4)
	if b.Fields[4][4] != BLACK {
		t.Error("Incorrect color at 4|4:", b.Fields[4][4])
	}
}

func TestWinner(t *testing.T) {
	b := NewBoard()

	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 1, 1, 1, 1, 1},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	if b.Winner() != WHITE {
		t.Error("Expected winner white horizontally")
	}

	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{0, 1, 0, 0, 0, 0},
	}

	if b.Winner() != WHITE {
		t.Error("Expected winner white vertically")
	}

	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{0, 0, 1, 0, 0, 0},
		[6]int{0, 0, 0, 1, 0, 0},
		[6]int{0, 0, 0, 0, 1, 0},
		[6]int{0, 0, 0, 0, 0, 1},
	}

	if b.Winner() != WHITE {
		t.Error("Expected winner white diagonally")
	}

	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, -1, 0},
		[6]int{0, 0, 1, -1, 0, 0},
		[6]int{0, 0, -1, 1, 0, 0},
		[6]int{0, -1, 0, 0, 1, 0},
		[6]int{-1, 0, 0, 0, 0, 1},
	}

	if b.Winner() != BLACK {
		t.Error("Expected winner black diagonally")
	}

	b.Fields = [6][6]int{
		[6]int{0, 1, 0, 0, 0, 0},
		[6]int{0, 0, 1, 0, 0, 0},
		[6]int{0, 0, 0, 1, 0, 0},
		[6]int{0, 0, 0, 0, 1, 0},
		[6]int{0, 0, 0, 0, 0, 1},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	if b.Winner() != WHITE {
		t.Error("Expected winner white diagonally (small diagonal)")
	}

	b.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, -1},
		[6]int{0, 0, 0, 0, -1, 0},
		[6]int{0, 0, 0, -1, 0, 0},
		[6]int{0, 0, -1, 0, 0, 0},
		[6]int{0, -1, 0, 0, 0, 0},
	}

	if b.Winner() != BLACK {
		t.Error("Expected winner black diagonally (small diagonal)")
	}

	b.Fields = [6][6]int{
		[6]int{0, 1, 0, 1, 0, 0},
		[6]int{1, 1, 1, 1, -1, -1},
		[6]int{-1, -1, -1, 1, -1, 0},
		[6]int{0, 0, 0, -1, 0, 0},
		[6]int{0, 0, -1, 0, 0, 0},
		[6]int{0, 1, 0, 0, 0, 0},
	}

	if b.Winner() != 0 {
		t.Error("Expected to find no winner yet")
	}

	b.Fields = [6][6]int{
		[6]int{1, 1, 1, 1, 1, 0},
		[6]int{-1, -1, -1, -1, -1, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	if b.Winner() != DRAW {
		t.Error("Expected a draw, since both players have a winning position")
	}

	b.Fields = [6][6]int{
		[6]int{1, 1, 1, 1, -1, 1},
		[6]int{-1, -1, -1, -1, 1, -1},
		[6]int{1, 1, 1, 1, -1, 1},
		[6]int{-1, -1, -1, -1, 1, -1},
		[6]int{1, 1, 1, 1, -1, 1},
		[6]int{-1, -1, -1, -1, 1, -1},
	}

	if b.Winner() != DRAW {
		t.Error("Expected a draw, since board is full")
	}
}

func TestRotateClockwise(t *testing.T) {
	b := NewBoard()

	b.Fields = [6][6]int{
		[6]int{1, -1, 0, 0, -1, 0},
		[6]int{0, 0, 0, 0, -1, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 1, 0, -1, 1},
		[6]int{0, 0, 0, 0, 1, 0},
		[6]int{1, 0, 0, 0, -1, 0},
	}

	b = b.Rotate(UPPERLEFT, CLOCKWISE)
	b = b.Rotate(UPPERRIGHT, CLOCKWISE)
	b = b.Rotate(LOWERLEFT, CLOCKWISE)
	b = b.Rotate(LOWERRIGHT, CLOCKWISE)

	expected := NewBoard()
	expected.Fields = [6][6]int{
		[6]int{0, 0, 1, 0, 0, 0},
		[6]int{0, 0, -1, 0, -1, -1},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, -1, 1, -1},
		[6]int{0, 0, 1, 0, 0, 1},
	}

	if !b.Equals(expected) {
		t.Error("Unexpected rotation: ", b)
	}
}

func TestRotateCounterClockwise(t *testing.T) {
	b := NewBoard()

	b.Fields = [6][6]int{
		[6]int{1, -1, 0, 0, -1, 0},
		[6]int{0, 0, 0, 0, -1, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, -1, 1, 0, -1, 1},
		[6]int{0, 0, 0, 0, 1, 0},
		[6]int{1, 0, 0, 0, -1, 0},
	}

	b = b.Rotate(UPPERLEFT, COUNTERCLOCKWISE)
	b = b.Rotate(UPPERRIGHT, COUNTERCLOCKWISE)
	b = b.Rotate(LOWERLEFT, COUNTERCLOCKWISE)
	b = b.Rotate(LOWERRIGHT, COUNTERCLOCKWISE)

	expected := NewBoard()
	expected.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{-1, 0, 0, -1, -1, 0},
		[6]int{1, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 1, 0, 0},
		[6]int{-1, 0, 0, -1, 1, -1},
		[6]int{0, 0, 1, 0, 0, 0},
	}

	if !b.Equals(expected) {
		t.Error("Unexpected rotation: ", b)
	}
}

func TestEquality(t *testing.T) {
	b := NewBoard()
	b.Fields = [6][6]int{
		[6]int{1, 1, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b2 := NewBoard()
	b2.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 1},
		[6]int{0, 0, 0, 0, 0, 1},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	if !b.EqualsIgnoreRotation(b2) {
		t.Error("Board should be equal to b2")
	}

	b3 := NewBoard()
	b3.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 1, 1},
	}

	if !b.EqualsIgnoreRotation(b3) {
		t.Error("Board should be equal to b3")
	}
	b4 := NewBoard()
	b4.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
	}

	if !b.EqualsIgnoreRotation(b4) {
		t.Error("Board should be equal to b4")
	}

	bdiff := NewBoard()
	bdiff.Fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, -1, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
	}

	if b.EqualsIgnoreRotation(bdiff) {
		t.Error("Board should be different from bdiff")
	}
}
