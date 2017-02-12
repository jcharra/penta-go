package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jcharra/penta-go/ai"
	"github.com/jcharra/penta-go/core"
	"github.com/jcharra/penta-go/view"
)

func main() {
	fmt.Println("Welcome to Pentago")

	interactive := flag.Bool("i", false, "interactive")

	flag.Parse()

	if *interactive {
		fmt.Println("Starting interactive play ...")
		scanner := bufio.NewScanner(os.Stdin)

		var color int
		var err error
		for {
			fmt.Println("Will you play white (1) or black (-1)?")
			scanner.Scan()
			color, err = strconv.Atoi(scanner.Text())
			if err != nil || (color != 1 && color != -1) {
				fmt.Println("Invalid input")
			} else {
				break
			}
		}

		fmt.Println("Starting game")
		b := core.NewBoard()
		for b.Winner() == 0 {
			fmt.Printf("\nBoard:\n%v\n", b.Repr())

			if b.Turn == color {
				fmt.Println("Your move (row, col, e.g. '0 5')?")
				scanner.Scan()
				input := strings.Split(scanner.Text(), " ")
				if len(input) != 2 {
					continue
				}

				row, errRow := strconv.Atoi(input[0])
				col, errCol := strconv.Atoi(input[1])

				if errRow != nil || errCol != nil || row < 0 || row > 5 || col < 0 || col > 5 {
					continue
				}

				if b.Fields[row][col] != 0 {
					fmt.Println("Field is blocked")
					continue
				}

				fmt.Println("\nRotate which quadrant? \n0 1\t\t0 = clockwise\n2 3\t\t1 = counterclockwise\n?")
				scanner.Scan()
				input = strings.Split(scanner.Text(), " ")
				if len(input) != 2 {
					continue
				}

				quad, errQuad := strconv.Atoi(input[0])
				direction, errDirection := strconv.Atoi(input[1])

				if errQuad != nil || errDirection != nil || quad < 0 || quad > 3 || (direction != 0 && direction != 1) {
					continue
				}

				b = b.SetAt(row, col).Rotate(quad, direction)
			} else {

				move := ai.FindBestMove(b, 5, 3).Move
				fmt.Println("My move: ", move.Repr())
				b = b.SetAt(move.Row, move.Col).Rotate(move.Quadrant, move.Direction)
			}
		}

		w := b.Winner()
		if w == core.WHITE {
			fmt.Println("White wins")
		} else if w == core.BLACK {
			fmt.Println("Black wins")
		} else {
			fmt.Println("Game is drawn")
		}

	} else {
		view.RunUI()
	}
}
