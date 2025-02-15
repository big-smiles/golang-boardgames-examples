package tictactoeGame

import (
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/tictTacToeData"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	instructioncontrol "github.com/big-smiles/golang-boardgames/pkg/instructions/control"
	instructionentity "github.com/big-smiles/golang-boardgames/pkg/instructions/entity"
	instructionEntityModifier "github.com/big-smiles/golang-boardgames/pkg/instructions/entity_modifier"
	instructionOutput "github.com/big-smiles/golang-boardgames/pkg/instructions/output"
	instruction_phase "github.com/big-smiles/golang-boardgames/pkg/instructions/phase"
	"github.com/big-smiles/golang-boardgames/pkg/phaseData"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/golang-boardgames/pkg/value_modifier/common"
)

type TicTacToeData struct {
	g *game.DataGame
}

func NewTicTacToeData() (*TicTacToeData, error) {
	libraryEntity := getLibraryEntity()
	players := []player.Id{tictTacToeData.PlayerId1, tictTacToeData.PlayerId2}
	libraryPhase := getPhaseLibrary()
	firstPhase := tictTacToeData.PhaseCreateBoard
	g, err := game.NewDataGame(libraryEntity, libraryPhase, firstPhase, players)
	if err != nil {
		panic(err)
	}
	return &TicTacToeData{
		g: g,
	}, nil
}

func getLibraryEntity() entity.LibraryDataEntity {
	library := make(entity.LibraryDataEntity, 1)
	data := getSquareData()
	library[tictTacToeData.SquareData] = data
	return library
}
func getSquareData() entity.DataEntity {
	resolverName := resolveValueConstant.NewResolveConstant[entity.NameEntityId]("")
	dataId, err := entity.NewDataId(resolverName)
	if err != nil {
		panic(err)
	}

	dataProperties := entity.DataProperties{
		IntProperties: []entity.NamePropertyId[int]{
			tictTacToeData.IntPropertyNameY,
			tictTacToeData.IntPropertyNameX,
			tictTacToeData.IntPropertyNameState,
		},
	}
	dataEntity, err := entity.NewDataEntity(*dataId, dataProperties)
	if err != nil {
		panic(err)
	}
	return *dataEntity
}

func getPhaseLibrary() []phaseData.DataPhase {
	createBoard := getPhaseCreateBoard()
	playersTurns := tictTacToeData.GetPlayersTurnPhase()
	return []phaseData.DataPhase{
		*createBoard,
		*playersTurns,
	}
}
func getPhaseCreateBoard() *phaseData.DataPhase {
	p := phaseData.DataPhase{
		Name: tictTacToeData.PhaseCreateBoard,
		Turns: []phaseData.DataTurn{
			{
				ActivePlayers: []player.Id{},
				Stages: []phaseData.DataStage{
					{
						Instructions: instructioncontrol.NewDataInstructionArray(
							getCreateSquare(0, 0, tictTacToeData.StateEmpty),
							getCreateSquare(0, 1, tictTacToeData.StateEmpty),
							getCreateSquare(0, 2, tictTacToeData.StateEmpty),
							getCreateSquare(1, 0, tictTacToeData.StateEmpty),
							getCreateSquare(1, 1, tictTacToeData.StateEmpty),
							getCreateSquare(1, 2, tictTacToeData.StateEmpty),
							getCreateSquare(2, 0, tictTacToeData.StateEmpty),
							getCreateSquare(2, 1, tictTacToeData.StateEmpty),
							getCreateSquare(2, 2, tictTacToeData.StateEmpty),
							instructionOutput.NewDataInstructionSendOutput(),
						),
					},
					{
						Instructions: instructioncontrol.NewDataInstructionArray(
							instruction_phase.NewDataInstructionSetNextPhase(
								resolveValueConstant.NewResolveConstant(tictTacToeData.PhasePlayerTurns),
							),
						),
					},
				},
			},
		},
	}
	return &p
}
func getCreateSquare(x int, y int, state int) instruction.DataInstruction {
	xValueResolver := resolveValueConstant.NewResolveConstant[int](x)
	setXModifier, err := ValueModifierCommon.NewDataModifierSetValue(xValueResolver)
	if err != nil {
		panic(err)
	}

	yValueResolver := resolveValueConstant.NewResolveConstant[int](y)

	setYModifier, err := ValueModifierCommon.NewDataModifierSetValue(yValueResolver)

	stateValueResolver := resolveValueConstant.NewResolveConstant[int](state)

	setStateModifier, err := ValueModifierCommon.NewDataModifierSetValue(stateValueResolver)
	if err != nil {
		panic(err)
	}
	intPropertiesModifier := make(entity.MapDataModifierProperties[int], 3)
	intPropertiesModifier[tictTacToeData.IntPropertyNameX] = setXModifier
	intPropertiesModifier[tictTacToeData.IntPropertyNameY] = setYModifier
	intPropertiesModifier[tictTacToeData.IntPropertyNameState] = setStateModifier

	dataPropertiesModifier := entity.DataPropertiesModifier{
		IntModifiers: intPropertiesModifier,
	}
	dataEntityModifier, err := entity.NewDataEntityModifier(dataPropertiesModifier)
	if err != nil {
		panic(err)
	}
	resolverScalar := resolveValueConstant.NewResolveValueFromVariable[entity.Id](tictTacToeData.EntityIdPropertyNameStoreEntity)
	targetResolver := resolveValueConstant.NewResolveScalarToSlice[entity.Id](resolverScalar)

	dataArrayInstruction := instructioncontrol.NewDataInstructionArray(
		instructionentity.NewDataInstructionCreateEntityIntoVariable(tictTacToeData.SquareData, tictTacToeData.EntityIdPropertyNameStoreEntity),
		instructionEntityModifier.NewDataInstructionAddEntityModifierWithResolvedTarget(targetResolver, *dataEntityModifier),
	)
	return dataArrayInstruction
}
