package ai

import "github.com/jcharra/pentago/core"

const CENTER_BONUS int = 10
const CHAIN_BONUS_MIDDLE int = 5
const CHAIN_BONUS_OUTER int = 3
const WINNER_VALUE int = 1000000

type EvaluatedMove struct {
	move  core.Move
	value int
}

func FindBestMove(b core.Board, breadth, depth int) core.Move {
	return FindBestMoves(b, breadth, depth)[0].move
}

func FindBestMoves(b core.Board, breadth, depth int) []EvaluatedMove {
	succs := core.FindSuccessors(b)

	// This is our list of <breadth> best moves, sorted by their evaluation desc
	bestMoves := make([]EvaluatedMove, breadth)
	for i := 0; i < breadth; i++ {
		bestMoves[i] = EvaluatedMove{value: -WINNER_VALUE}
	}

	for board, move := range succs {
		val := evaluate(board) * colorSign(b.Turn)

		for i := 0; i < breadth; i++ {
			if val > bestMoves[i].value {
				// insert move into ist of best moves, pushing out the worst of them
				copy(bestMoves[i+1:], bestMoves[i:])
				bestMoves[i] = EvaluatedMove{move: move, value: val}
				break
			}
		}
	}

	//fmt.Printf("\nBest moves: %v", bestMoves)
	return bestMoves
}

func evaluate(b core.Board) int {
	winner := b.Winner()
	if winner == core.WHITE {
		return WINNER_VALUE
	} else if winner == core.BLACK {
		return -WINNER_VALUE
	}

	val := 0

	// Centers are important
	for _, col := range []int{b.Fields[1][1], b.Fields[1][4], b.Fields[4][1], b.Fields[4][4]} {
		if col != 0 {
			val += CENTER_BONUS * colorSign(col)
		}
	}

	// Scan for horizontal/vertical chains, rewarding greater lengths
	lastSeenVertical := b.Fields[0]
	lastSeenHorizontal := [6]int{b.Fields[0][0], b.Fields[1][0], b.Fields[2][0], b.Fields[3][0], b.Fields[4][0], b.Fields[5][0]}
	for i := 1; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if lastSeenHorizontal[j] == b.Fields[j][i] {
				val += colorSign(b.Fields[j][i]) * chainBonus(j)
			} else {
				lastSeenHorizontal[j] = b.Fields[j][i]
			}

			if lastSeenVertical[j] == b.Fields[i][j] {
				val += colorSign(b.Fields[i][j]) * chainBonus(j)
			} else {
				lastSeenVertical[j] = b.Fields[i][j]
			}
		}
	}

	// TODO: how to evaluate the diagonals?

	return val
}

func colorSign(color int) int {
	if color == core.WHITE {
		return 1
	} else if color == core.BLACK {
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
