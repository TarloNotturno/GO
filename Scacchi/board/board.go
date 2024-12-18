package board

import "fmt"

// additional function, just to print the maze layout
const FULLSQUARE = 11036
const EMPTYSQUARE = 11035

//pieceUnicodes = []string{" ", "♔", "♕", "♖", "♗", "♘", "♙", "♚", "♛", "♜", "♝", "♞", "♟"}

type Piece struct {
	y    int
	x    int
	role int
	plot string
}

type pos struct {
	y int
	x int
}

type ChessBoard struct {
	black map[pos]Piece
	white map[pos]Piece
}

func (p Piece) GetPos() (int, int) {
	return p.x, p.y
}

func (c ChessBoard) UpdateBoard() {
	N := 8
	fmt.Println("   A  B   C  D   E  F   G  H")
	fmt.Println("  ===========================")
	for y := 0; y < N; y++ {
		fmt.Print(y, "|")
		for x := 0; x < N; x++ {
			black, ok1 := c.black[pos{y: y, x: x}]
			white, ok2 := c.white[pos{y: y, x: x}]
			if ok1 {
				//if x%2 == 0 {
				//	fmt.Print(" ")
				//}
				fmt.Print(black.plot, "|")
			} else if ok2 {
				//if x%2 == 0 {
				//	fmt.Print(" ")
				//}
				fmt.Print(white.plot, "|")
			} else { //if x%2 == y%2 {
				fmt.Print("  |")
				//fmt.Print(string(rune(FULLSQUARE)))
				//} else {
				//
				//	fmt.Print("[]|")
				//fmt.Print(string(rune(EMPTYSQUARE)))
			}
		}
		fmt.Println()
	}
	fmt.Println("  ===========================")
	fmt.Println()
	fmt.Println()
}
