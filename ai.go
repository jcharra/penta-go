package pentago

const CENTER_BONUS int = 10
const CHAIN_BONUS_MIDDLE int = 5
const CHAIN_BONUS_OUTER int = 3
const WINNER_VALUE int = 1000000

func FindBestMove(b Board, breadth, depth int) Move {
	succs := FindSuccessors(b)
	bestEval := -WINNER_VALUE
	var bestMove Move

	for board, move := range succs {
		val := evaluate(board) * colorSign(b.Turn)
		if val > bestEval {
			bestMove = move
			bestEval = val
		}
	}

	//fmt.Printf("\nBest: %v with eval %v", bestMove, bestEval)
	return bestMove
}

func evaluate(b Board) int {
	winner := b.Winner()
	if winner == WHITE {
		return WINNER_VALUE
	} else if winner == BLACK {
		return -WINNER_VALUE
	}

	val := 0

	// Centers are important
	for _, col := range []int{b.fields[1][1], b.fields[1][4], b.fields[4][1], b.fields[4][4]} {
		if col != 0 {
			val += CENTER_BONUS * colorSign(col)
		}
	}

	// Scan for horizontal/vertical chains, rewarding greater lengths
	lastSeenVertical := b.fields[0]
	lastSeenHorizontal := [6]int{b.fields[0][0], b.fields[1][0], b.fields[2][0], b.fields[3][0], b.fields[4][0], b.fields[5][0]}
	for i := 1; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if lastSeenHorizontal[j] == b.fields[j][i] {
				val += colorSign(b.fields[j][i]) * chainBonus(j)
			} else {
				lastSeenHorizontal[j] = b.fields[j][i]
			}

			if lastSeenVertical[j] == b.fields[i][j] {
				val += colorSign(b.fields[i][j]) * chainBonus(j)
			} else {
				lastSeenVertical[j] = b.fields[i][j]
			}
		}
	}

	// TODO: how to evaluate the diagonals?

	return val
}

func colorSign(color int) int {
	if color == WHITE {
		return 1
	} else if color == BLACK {
		return -1
	}
	return 0
}

// Having multiple checkers in a row count differently on the row/column
// indexes 1 and 4, as they are more important.
func chainBonus(arrIdx int) int {
	if arrIdx == 1 || arrIdx == 4 {
		return CHAIN_BONUS_MIDDLE
	} else {
		return CHAIN_BONUS_OUTER
	}
}
