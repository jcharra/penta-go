package ai

import "github.com/jcharra/pentago/core"

const centerBonus int = 10
const chainBonusMiddle int = 5
const chainBonusOuter int = 3
const winnerValue int = 1000000

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type EvaluatedMove struct {
	move  core.Move
	value int
}

func FindBestMove(b core.Board, breadth, depth int) EvaluatedMove {
	succs := core.FindSuccessors(b)

	// This is our list of <breadth> best moves, sorted by their evaluation desc
	bestMoves := make([]EvaluatedMove, breadth)
	for i := 0; i < breadth; i++ {
		bestMoves[i] = EvaluatedMove{value: getWorstValue(b.Turn)}
	}

	for board, move := range succs {
		val := evaluate(board)

		for i := 0; i < breadth; i++ {
			if better(val, bestMoves[i].value, b.Turn) {
				// insert move into list of best moves, pushing out the worst of them
				copy(bestMoves[i+1:], bestMoves[i:])
				bestMoves[i] = EvaluatedMove{move: move, value: val}
				break
			}
		}
	}

	// depth == 0 means we do not recurse and just pick the seemingly best move from our list.
	if depth == 0 {
		return bestMoves[0]
	}

	// Re-evaluate the current list's <breadth> elements by considering the
	// optimal opponent's move
	var bestMove EvaluatedMove
	bestEval := getWorstValue(b.Turn)
	for _, bm := range bestMoves {
		boardAfterMove := b.SetAt(bm.move.Row, bm.move.Col)
		opponentMove := FindBestMove(boardAfterMove, breadth, depth-1)
		bm.value = opponentMove.value

		if better(bm.value, bestEval, b.Turn) {
			bestMove = bm
			bestEval = bm.value
		}
	}

	return bestMove
}

// Returns whether <a> is a better value than <b> from <color>'s perspective
func better(a, b, color int) bool {
	if color == core.WHITE {
		return a > b
	} else {
		return a < b
	}
}

func getWorstValue(color int) int {
	if color == core.WHITE {
		return MinInt
	} else {
		return MaxInt
	}
}

func evaluate(b core.Board) int {
	winner := b.Winner()
	if winner == core.WHITE {
		return winnerValue
	} else if winner == core.BLACK {
		return -winnerValue
	}

	val := 0

	// Centers are important
	for _, col := range []int{b.Fields[1][1], b.Fields[1][4], b.Fields[4][1], b.Fields[4][4]} {
		if col != 0 {
			val += centerBonus * colorSign(col)
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
		return chainBonusMiddle
	} else {
		return chainBonusOuter
	}
}
