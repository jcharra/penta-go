package pentago

import "testing"

func TestBoard(t *testing.T) {
	b := NewBoard()
	if b.fields[5][5] != 0 {
		t.Error("Nope")
	}
}

func TestSetAt(t *testing.T) {
	b := NewBoard()

	if !b.WhitesTurn {
		t.Error("Must be white's turn")
	}

	b, _ = b.SetAt(5, 5)
	if b.fields[5][5] != WHITE {
		t.Error("Incorrect color at 5|5:", b.fields[5][5])
	}

	if b.WhitesTurn {
		t.Error("Must be black's turn")
	}

	b, _ = b.SetAt(4, 4)
	if b.fields[4][4] != BLACK {
		t.Error("Incorrect color at 4|4:", b.fields[4][4])
	}

	b, err := b.SetAt(4, 4)
	if err == nil {
		t.Error("This should not have worked")
	}
}

func TestFindSuccessors(t *testing.T) {
	b := NewBoard()
	successors := b.findSuccessors()

	if len(successors) != 9 {
		t.Errorf("Expected 9 successors, found %v: %v", len(successors), successors)
	}
}

func TestRotateClockwise(t *testing.T) {
	b := NewBoard()

	b.fields = [6][6]int{
		[6]int{1, 2, 0, 0, 2, 0},
		[6]int{0, 0, 0, 0, 2, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 1, 0, 2, 1},
		[6]int{0, 0, 0, 0, 1, 0},
		[6]int{1, 0, 0, 0, 2, 0},
	}

	b = b.Rotate(UPPERLEFT, CLOCKWISE)
	b = b.Rotate(UPPERRIGHT, CLOCKWISE)
	b = b.Rotate(LOWERLEFT, CLOCKWISE)
	b = b.Rotate(LOWERRIGHT, CLOCKWISE)

	expected := NewBoard()
	expected.fields = [6][6]int{
		[6]int{0, 0, 1, 0, 0, 0},
		[6]int{0, 0, 2, 0, 2, 2},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 2, 1, 2},
		[6]int{0, 0, 1, 0, 0, 1},
	}

	if !b.Equals(expected) {
		t.Error("Unexpected rotation: ", b)
	}
}

func TestRotateCounterClockwise(t *testing.T) {
	b := NewBoard()

	b.fields = [6][6]int{
		[6]int{1, 2, 0, 0, 2, 0},
		[6]int{0, 0, 0, 0, 2, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 2, 1, 0, 2, 1},
		[6]int{0, 0, 0, 0, 1, 0},
		[6]int{1, 0, 0, 0, 2, 0},
	}

	b = b.Rotate(UPPERLEFT, COUNTERCLOCKWISE)
	b = b.Rotate(UPPERRIGHT, COUNTERCLOCKWISE)
	b = b.Rotate(LOWERLEFT, COUNTERCLOCKWISE)
	b = b.Rotate(LOWERRIGHT, COUNTERCLOCKWISE)

	expected := NewBoard()
	expected.fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{2, 0, 0, 2, 2, 0},
		[6]int{1, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 1, 0, 0},
		[6]int{2, 0, 0, 2, 1, 2},
		[6]int{0, 0, 1, 0, 0, 0},
	}

	if !b.Equals(expected) {
		t.Error("Unexpected rotation: ", b)
	}
}

func TestEquality(t *testing.T) {
	b := NewBoard()
	b.fields = [6][6]int{
		[6]int{1, 1, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
	}

	b2 := NewBoard()
	b2.fields = [6][6]int{
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
	b3.fields = [6][6]int{
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
	b4.fields = [6][6]int{
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
	bdiff.fields = [6][6]int{
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{0, 0, 2, 0, 0, 0},
		[6]int{0, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
		[6]int{1, 0, 0, 0, 0, 0},
	}

	if b.EqualsIgnoreRotation(bdiff) {
		t.Error("Board should be different from bdiff")
	}
}
