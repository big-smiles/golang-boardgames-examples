package main

import (
	"fmt"
	tictactoeGame "github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/game"
)

func main() {
	fmt.Println("Hello tic-tac-toe")

	data, err := tictactoeGame.NewTicTacToeData()
	if err != nil {
		panic(err)
	}
	g, err := tictactoeGame.NewTicTacToeGame(*data)
	if err != nil {
		panic(err)
	}
	err = g.Start()
	if err != nil {
		panic(err)
	}
}
