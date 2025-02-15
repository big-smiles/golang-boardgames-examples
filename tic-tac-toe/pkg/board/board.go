package board

import (
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/tictTacToeData"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/output"
)

type Board struct {
	Squares [3][3]int
	Ids     [3][3]entity.Id
}

func NewBoard(o output.Game) *Board {
	b := Board{}
	b.Squares = [3][3]int{}
	b.Ids = [3][3]entity.Id{}
	idX := o.PropertyIds.Int[tictTacToeData.IntPropertyNameX]
	idY := o.PropertyIds.Int[tictTacToeData.IntPropertyNameY]
	idState := o.PropertyIds.Int[tictTacToeData.IntPropertyNameState]

	for _, e := range o.Entities {
		x := e.Properties.IntProperties[idX]
		y := e.Properties.IntProperties[idY]
		state := e.Properties.IntProperties[idState]
		id := e.Id
		b.Squares[x][y] = state
		b.Ids[x][y] = id
	}
	return &b
}
func (b *Board) GetId(x int, y int) entity.Id {
	return b.Ids[x][y]
}
