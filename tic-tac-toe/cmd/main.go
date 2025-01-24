package main

import (
	"fmt"
	tictactoeGame "github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
)

func main() {
	fmt.Println("Hello tic-tac-toe")
	var callbackOutput output.Callback = func(output *output.Game) {
		fmt.Println("output callback")
		fmt.Println(output)
	}
	var callbackInput interaction.Callback = func(interactions []interaction.OutputInteraction) {
		fmt.Println("interactions callback")
	}
	data := tictactoeGame.NewTicTacToeData()
	g, err := tictactoeGame.NewTicTacToeGame(*data, callbackInput, callbackOutput)
	if err != nil {
		panic(err)
	}
	err = g.Start()
	if err != nil {
		panic(err)
	}
}
