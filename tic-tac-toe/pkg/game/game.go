package tictactoeGame

import (
	"github.com/big-smiles/golang-boardgames/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
)

type TicTacToeGame struct {
	managerGame *game.ManagerGame
}

func NewTicTacToeGame(data TicTacToeData, callbackInteraction interaction.Callback, callbackOutput output.Callback) (*TicTacToeGame, error) {
	managerGame, err := game.NewGame(*data.g, callbackOutput, callbackInteraction)
	if err != nil {
		return nil, err
	}
	return &TicTacToeGame{
		managerGame: managerGame,
	}, nil
}
func (game *TicTacToeGame) Start() error {
	return game.managerGame.Start()
}
func (game *TicTacToeGame) SelectInteraction(selectedInteractions []interaction.SelectedInteraction) error {
	return game.managerGame.SelectInteraction(selectedInteractions)
}
