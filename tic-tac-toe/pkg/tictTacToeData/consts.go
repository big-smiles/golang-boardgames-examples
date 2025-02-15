package tictTacToeData

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/player"
)

const (
	SquareData                      entity.NameDataEntity            = "squareData"
	PhaseCreateBoard                phase.NamePhase                  = "phase_create_board"
	PhasePlayerTurns                phase.NamePhase                  = "phase_player_turns"
	IntPropertyNameX                entity.NamePropertyId[int]       = "x"
	IntPropertyNameY                entity.NamePropertyId[int]       = "y"
	IntPropertyNameState            entity.NamePropertyId[int]       = "state"
	EntityIdPropertyNameStoreEntity entity.NamePropertyId[entity.Id] = "storeEntity"
	StateEmpty                      int                              = 3
	StatePlayer1                    int                              = 1
	StatePlayer2                    int                              = 2
	PlayerId1                       player.Id                        = "player_id_1"
	PlayerId2                       player.Id                        = "player_id_2"
)
