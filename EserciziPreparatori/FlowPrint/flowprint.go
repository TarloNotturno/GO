package FlowPrint

import "fmt"

func printBoard(board [][]string) {
	for _, val := range board {
		for _, box := range val {
			fmt.Print(box)
		}
		fmt.Println()
	}

}

type coordinate struct {
	x int
	y int
}

func FillTheBoard(xIn int, yIn int, board [][]string) {
	input := coordinate{x: xIn, y: yIn}
	fmt.Println("board before flow")
	printBoard(board)
	if input.x < len(board[0]) && input.x >= 0 && input.y < len(board) && input.y >= 0 {
		if board[input.y][input.x] == "□" {
			queu := []coordinate{input}
			for len(queu) > 0 {
				x := queu[0].x
				y := queu[0].y
				if x >= 1 {
					if board[y][x-1] == "□" {
						queu = append(queu, coordinate{x: x - 1, y: y})
					}
				}
				if x < len(board[0])-1 {
					if board[y][x+1] == "□" {
						queu = append(queu, coordinate{x: x + 1, y: y})
					}
				}
				if y >= 1 {
					if board[y-1][x] == "□" {
						queu = append(queu, coordinate{x: x, y: y - 1})
					}
				}
				if y < len(board)-1 {
					if board[y+1][x] == "□" {
						queu = append(queu, coordinate{x: x, y: y + 1})
					}
				}
				board[y][x] = "■"
				queu = queu[1:]
				fmt.Print()
			}
		}
	}
	fmt.Printf("\n\n")
	fmt.Println("board after flow")
	printBoard(board)
}
