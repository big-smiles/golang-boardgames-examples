package main

import (
	"fmt"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/board"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/console_draw"
	tictactoeGame "github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
)

func main() {
	fmt.Println("Hello tic-tac-toe")
	var callbackOutput output.Callback = func(o *output.Game) {
		b := board.NewBoard(*o)

		console_draw.Draw(*b)
	}
	var callbackInput interaction.Callback = func(interactions []interaction.OutputInteraction) {
		fmt.Println("interactions callback")
	}
	data, err := tictactoeGame.NewTicTacToeData()
	if err != nil {
		panic(err)
	}
	g, err := tictactoeGame.NewTicTacToeGame(*data, callbackInput, callbackOutput)
	if err != nil {
		panic(err)
	}
	err = g.Start()
	if err != nil {
		panic(err)
	}
}
