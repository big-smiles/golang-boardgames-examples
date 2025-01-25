package tictactoeGame

import (
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/game"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	instructioncontrol "github.com/big-smiles/golang-boardgames/pkg/instructions/control"
	instructionentity "github.com/big-smiles/golang-boardgames/pkg/instructions/entity"
	instructionEntityModifier "github.com/big-smiles/golang-boardgames/pkg/instructions/entity_modifier"
	instructionOutput "github.com/big-smiles/golang-boardgames/pkg/instructions/output"
	"github.com/big-smiles/golang-boardgames/pkg/phase"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/golang-boardgames/pkg/value_modifier/common"
)

const (
	SquareData                      entity.NameDataEntity            = "squareData"
	PhaseCreateBoard                phase.NamePhase                  = "phase_create_board"
	IntPropertyNameX                entity.NamePropertyId[int]       = "x"
	IntPropertyNameY                entity.NamePropertyId[int]       = "y"
	IntPropertyNameState            entity.NamePropertyId[int]       = "state"
	EntityIdPropertyNameStoreEntity entity.NamePropertyId[entity.Id] = "storeEntity"
	StateEmpty                      int                              = 3
	StatePlayer1                    int                              = 1
	StatePlayer2                    int                              = 2
)

type TicTacToeData struct {
	g *game.DataGame
}

func NewTicTacToeData() (*TicTacToeData, error) {
	libraryEntity := getLibraryEntity()
	players := getPlayers()
	libraryPhase, err := getPhaseLibrary()
	if err != nil {
		return nil, err
	}
	firstPhase := getFirstPhase()
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
	library[SquareData] = data
	return library
}
func getSquareData() entity.DataEntity {
	resolverName, err := resolveValueConstant.NewResolveConstant[entity.NameEntityId]("")
	if err != nil {
		panic(err)
	}
	dataId, err := entity.NewDataId(resolverName)
	if err != nil {
		panic(err)
	}
	boolProperties := []entity.NamePropertyId[bool]{}
	intProperties := []entity.NamePropertyId[int]{IntPropertyNameY, IntPropertyNameX, IntPropertyNameState}
	stringProperties := []entity.NamePropertyId[string]{}
	entityIdProperties := []entity.NamePropertyId[entity.Id]{}
	arrayEntityIdProperties := []entity.NamePropertyId[[]entity.Id]{}
	dataProperties, err := entity.NewDataProperties(boolProperties, stringProperties, entityIdProperties, intProperties, arrayEntityIdProperties)
	if err != nil {
		panic(err)
	}
	dataEntity, err := entity.NewDataEntity(*dataId, *dataProperties)
	if err != nil {
		panic(err)
	}
	return *dataEntity
}
func getPlayers() []player.Id {
	var player1 player.Id = "player_1"
	var player2 player.Id = "player_2"
	players := []player.Id{player1, player2}
	return players
}

func getPhaseLibrary() (phase.LibraryPhase, error) {
	libraryPhase := make(phase.LibraryPhase, 1)
	a, err := getPhaseCreateBoard()
	if err != nil {
		return nil, err
	}
	libraryPhase[PhaseCreateBoard] = *a
	return libraryPhase, nil
}
func getFirstPhase() phase.NamePhase {
	return PhaseCreateBoard
}
func getPhaseCreateBoard() (*phase.DataPhase, error) {
	createEntity00, err := getCreateSquare(0, 0, StateEmpty)
	createEntity01, err := getCreateSquare(0, 1, StateEmpty)
	createEntity02, err := getCreateSquare(0, 2, StateEmpty)
	createEntity10, err := getCreateSquare(1, 0, StateEmpty)
	createEntity11, err := getCreateSquare(1, 1, StateEmpty)
	createEntity12, err := getCreateSquare(1, 2, StateEmpty)
	createEntity20, err := getCreateSquare(2, 0, StateEmpty)
	createEntity21, err := getCreateSquare(2, 1, StateEmpty)
	createEntity22, err := getCreateSquare(2, 2, StateEmpty)
	sendOutputDataInstruction, err := instructionOutput.NewDataInstructionSendOutput()
	if err != nil {
		return nil, err
	}
	a := make([]instruction.DataInstruction, 10)
	a[0] = createEntity00
	a[1] = createEntity01
	a[2] = createEntity02
	a[3] = createEntity10
	a[4] = createEntity11
	a[5] = createEntity12
	a[6] = createEntity20
	a[7] = createEntity21
	a[8] = createEntity22
	a[9] = sendOutputDataInstruction

	dataArrayInstruction, err := instructioncontrol.NewDataInstructionArray(a)
	if err != nil {
		return nil, err
	}
	stages, err := phase.NewDataStage(dataArrayInstruction)
	if err != nil {
		return nil, err
	}

	turns, err := phase.NewDataTurn([]player.Id{}, []phase.DataStage{*stages})
	if err != nil {
		return nil, err
	}

	p, err := phase.NewDataPhase([]phase.DataTurn{*turns})
	if err != nil {
		return nil, err
	}
	return p, nil
}
func getCreateSquare(x int, y int, state int) (instruction.DataInstruction, error) {
	createEntity1DataInstruction, err := instructionentity.NewDataInstructionCreateEntityIntoVariable(SquareData, EntityIdPropertyNameStoreEntity)
	if err != nil {
		panic(err)
	}

	xValueResolver, err := resolveValueConstant.NewResolveConstant[int](x)
	if err != nil {
		panic(err)
	}
	setXModifier, err := ValueModifierCommon.NewDataModifierSetValue(xValueResolver)
	if err != nil {
		panic(err)
	}

	yValueResolver, err := resolveValueConstant.NewResolveConstant[int](y)
	if err != nil {
		panic(err)
	}
	setYModifier, err := ValueModifierCommon.NewDataModifierSetValue(yValueResolver)
	if err != nil {
		panic(err)
	}

	stateValueResolver, err := resolveValueConstant.NewResolveConstant[int](state)
	if err != nil {
		panic(err)
	}
	setStateModifier, err := ValueModifierCommon.NewDataModifierSetValue(stateValueResolver)
	if err != nil {
		panic(err)
	}
	intPropertiesModifier := make(entity.MapDataModifierProperties[int], 3)
	intPropertiesModifier[IntPropertyNameX] = setXModifier
	intPropertiesModifier[IntPropertyNameY] = setYModifier
	intPropertiesModifier[IntPropertyNameState] = setStateModifier

	dataPropertiesModifier, err := entity.NewDataPropertiesModifier(
		&intPropertiesModifier,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}
	dataEntityModifier, err := entity.NewDataEntityModifier(*dataPropertiesModifier)
	if err != nil {
		panic(err)
	}
	resolverScalar, err := resolveValueConstant.NewResolveValueFromVariable[entity.Id](EntityIdPropertyNameStoreEntity)
	targetResolver, err := resolveValueConstant.NewResolveScalarToSlice[entity.Id](resolverScalar)
	if err != nil {
		panic(err)
	}
	setXYStateForSquareInstruction, err := instructionEntityModifier.NewDataInstructionAddEntityModifierWithResolvedTarget(targetResolver, *dataEntityModifier)
	if err != nil {
		panic(err)
	}

	a := make([]instruction.DataInstruction, 2)
	a[0] = createEntity1DataInstruction
	a[1] = setXYStateForSquareInstruction

	dataArrayInstruction, err := instructioncontrol.NewDataInstructionArray(a)
	if err != nil {
		panic(err)
	}
	return dataArrayInstruction, nil
}
