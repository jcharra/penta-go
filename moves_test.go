package pentago

import "testing"

func TestFindSuccessors(t *testing.T) {
	b := NewBoard()
	successors := FindSuccessors(b)

	if len(successors) != 9 {
		t.Errorf("Expected 9 successors, found %v: %v", len(successors), successors)
	}
}
