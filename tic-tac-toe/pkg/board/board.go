package board

import (
	tictactoeGame "github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/output"
)

type Board struct {
	Squares [3][3]int
}

func NewBoard(o output.Game) *Board {
	b := Board{}
	b.Squares = [3][3]int{}
	idX := o.PropertyIds.Int[tictactoeGame.IntPropertyNameX]
	idY := o.PropertyIds.Int[tictactoeGame.IntPropertyNameY]
	idState := o.PropertyIds.Int[tictactoeGame.IntPropertyNameState]

	for _, e := range o.Entities {
		x := e.Properties.IntProperties[idX]
		y := e.Properties.IntProperties[idY]
		state := e.Properties.IntProperties[idState]
		b.Squares[x][y] = state
	}
	return &b
}
