package board

type knight struct {
	piece Piece
}

func (k *knight) move(b *ChessBoard) {

}

func (c *ChessBoard) Initialize() {
	c.black = make(map[pos]Piece)
	towerBlack1 := Piece{x: 0, y: 0, role: 1, plot: "♖"}
	c.black[pos{x: 0, y: 0}] = towerBlack1
	towerBlack2 := Piece{x: 7, y: 0, role: 1, plot: "♖"}
	c.black[pos{x: 7, y: 0}] = towerBlack2
	knightBlack1 := Piece{x: 1, y: 0, role: 2, plot: "♘"}
	c.black[pos{x: 1, y: 0}] = knightBlack1
	knightBlack2 := Piece{x: 6, y: 0, role: 2, plot: "♘"}
	c.black[pos{x: 6, y: 0}] = knightBlack2
	bishopBlack1 := Piece{x: 2, y: 0, role: 3, plot: "♗"}
	c.black[pos{x: 2, y: 0}] = bishopBlack1
	bishopBlack2 := Piece{x: 5, y: 0, role: 3, plot: "♗"}
	c.black[pos{x: 5, y: 0}] = bishopBlack2
	kingBlack := Piece{x: 3, y: 0, role: 4, plot: "♔"}
	c.black[pos{x: 4, y: 0}] = kingBlack
	queenBlack := Piece{x: 4, y: 0, role: 5, plot: "♕"}
	c.black[pos{x: 3, y: 0}] = queenBlack

	pawnBlack0 := Piece{x: 0, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 0, y: 1}] = pawnBlack0
	pawnBlack1 := Piece{x: 1, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 1, y: 1}] = pawnBlack1
	pawnBlack2 := Piece{x: 2, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 2, y: 1}] = pawnBlack2
	pawnBlack3 := Piece{x: 3, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 3, y: 1}] = pawnBlack3
	pawnBlack4 := Piece{x: 4, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 4, y: 1}] = pawnBlack4
	pawnBlack5 := Piece{x: 5, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 5, y: 1}] = pawnBlack5
	pawnBlack6 := Piece{x: 6, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 6, y: 1}] = pawnBlack6
	pawnBlack7 := Piece{x: 7, y: 1, role: 0, plot: "♙"}
	c.black[pos{x: 7, y: 1}] = pawnBlack7

	c.white = make(map[pos]Piece)
	towerWhite1 := Piece{x: 0, y: 7, role: 1, plot: "♜"}
	c.white[pos{x: 0, y: 7}] = *&towerWhite1
	towerWhite2 := Piece{x: 7, y: 7, role: 1, plot: "♜"}
	c.white[pos{x: 7, y: 7}] = *&towerWhite2
	knightWhite1 := Piece{x: 1, y: 7, role: 2, plot: "♞"}
	c.white[pos{x: 1, y: 7}] = *&knightWhite1
	knightWhite2 := Piece{x: 6, y: 7, role: 2, plot: "♞"}
	c.white[pos{x: 6, y: 7}] = *&knightWhite2
	bishopWhite1 := Piece{x: 2, y: 7, role: 3, plot: "♝"}
	c.white[pos{x: 2, y: 7}] = *&bishopWhite1
	bishopWhite2 := Piece{x: 5, y: 7, role: 3, plot: "♝"}
	c.white[pos{x: 5, y: 7}] = *&bishopWhite2
	kingWhite := Piece{x: 3, y: 7, role: 4, plot: "♚"}
	c.white[pos{x: 3, y: 7}] = *&kingWhite
	queenWhite := Piece{x: 4, y: 7, role: 5, plot: "♛"}
	c.white[pos{x: 4, y: 7}] = *&queenWhite

	pawnWhite0 := Piece{x: 0, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 0, y: 6}] = pawnWhite0
	pawnWhite1 := Piece{x: 1, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 1, y: 6}] = pawnWhite1
	pawnWhite2 := Piece{x: 2, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 2, y: 6}] = pawnWhite2
	pawnWhite3 := Piece{x: 3, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 3, y: 6}] = pawnWhite3
	pawnWhite4 := Piece{x: 4, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 4, y: 6}] = pawnWhite4
	pawnWhite5 := Piece{x: 5, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 5, y: 6}] = pawnWhite5
	pawnWhite6 := Piece{x: 6, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 6, y: 6}] = pawnWhite6
	pawnWhite7 := Piece{x: 7, y: 6, role: 0, plot: "♟"}
	c.white[pos{x: 7, y: 6}] = pawnWhite7

}
