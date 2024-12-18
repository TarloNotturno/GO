package main

import (
	"main.go/board"
)

func main() {
	var game board.ChessBoard
	game.Initialize()
	game.UpdateBoard()
}
