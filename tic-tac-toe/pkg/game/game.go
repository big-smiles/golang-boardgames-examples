package tictactoeGame

import (
	"fmt"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/board"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/console_draw"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/input"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/tictTacToeData"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/output"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	"sync"
)

type TicTacToeGame struct {
	managerGame *game.ManagerGame
	inputReader *input.InputReader
	board       *board.Board
	mu          sync.Mutex
}

func NewTicTacToeGame(data TicTacToeData) (*TicTacToeGame, error) {
	g := TicTacToeGame{}
	managerGame, err := game.NewGame(*data.g, g.callbackOutput, g.callbackInteraction)
	if err != nil {
		return nil, err
	}
	g.managerGame = managerGame
	return &g, nil
}
func (g *TicTacToeGame) Start() error {
	return g.managerGame.Start()
}
func (g *TicTacToeGame) SelectInteraction(selectedInteractions []interaction.SelectedInteraction) error {
	return g.managerGame.SelectInteraction(selectedInteractions)
}

func (g *TicTacToeGame) callbackOutput(o *output.Game) {
	g.mu.Lock()
	defer g.mu.Unlock()
	fmt.Println(o)
	g.board = board.NewBoard(*o)
	console_draw.Draw(*g.board)
}

func (g *TicTacToeGame) callbackInteraction(interactions []interaction.OutputInteraction) {
	g.mu.Lock()
	defer g.mu.Unlock()
	fmt.Println("interactions callback")
	fmt.Println(interactions)
	selectedId, err := g.getInteraction(interactions[0].PlayerId, interactions[0].AvailableEntities)
	for err != nil {
		fmt.Println(err)
		selectedId, err = g.getInteraction(interactions[0].PlayerId, interactions[0].AvailableEntities)
	}
	selectedInteraction, err := interaction.NewSelectedInteraction(
		interactions[0].Id,
		interactions[0].PlayerId,
		[]entity.Id{selectedId},
	)
	if err != nil {
		panic(err)
	}
	selectedInteractions := []interaction.SelectedInteraction{
		*selectedInteraction,
	}
	g.mu.Unlock()
	err = g.managerGame.SelectInteraction(selectedInteractions)
	if err != nil {
		panic(err)
	}
}
func (g *TicTacToeGame) getInteraction(playerId player.Id, available []entity.Id) (entity.Id, error) {

	if playerId == tictTacToeData.PlayerId1 {
		fmt.Println("Player 1 turn")
	} else {
		fmt.Println("Player 2 turn")
	}
	x, y, err := g.inputReader.GetInput()
	if err != nil {
		return 0, err
	}
	entityId := g.board.GetId(x, y)

	var found bool = false
	for _, id := range available {
		if entityId == id {
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("invalid entityId %d for positions %d, %d", entityId, x, y)
	}
	return entityId, nil
}
