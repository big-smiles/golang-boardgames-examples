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

func NewTicTacToeData() *TicTacToeData {
	libraryEntity := getLibraryEntity()
	players := getPlayers()
	libraryPhase := getPhaseLibrary()
	firstPhase := getFirstPhase()
	g, err := game.NewDataGame(libraryEntity, libraryPhase, firstPhase, players)
	if err != nil {
		panic(err)
	}
	return &TicTacToeData{
		g: g,
	}
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

func getPhaseLibrary() phase.LibraryPhase {
	libraryPhase := make(phase.LibraryPhase, 1)
	libraryPhase[PhaseCreateBoard] = getPhaseCreateBoard()
	return libraryPhase
}
func getFirstPhase() phase.NamePhase {
	return PhaseCreateBoard
}
func getPhaseCreateBoard() phase.DataPhase {
	createEntity1DataInstruction, err := instructionentity.NewDataInstructionCreateEntityIntoVariable(SquareData, EntityIdPropertyNameStoreEntity)
	if err != nil {
		panic(err)
	}

	xValueResolver, err := resolveValueConstant.NewResolveConstant[int](0)
	if err != nil {
		panic(err)
	}
	setXModifier, err := ValueModifierCommon.NewDataModifierSetValue(xValueResolver)
	if err != nil {
		panic(err)
	}

	yValueResolver, err := resolveValueConstant.NewResolveConstant[int](0)
	if err != nil {
		panic(err)
	}
	setYModifier, err := ValueModifierCommon.NewDataModifierSetValue(yValueResolver)
	if err != nil {
		panic(err)
	}

	stateValueResolver, err := resolveValueConstant.NewResolveConstant[int](StateEmpty)
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

	dataPropertiesModifer, err := entity.NewDataPropertiesModifier(
		&intPropertiesModifier,
		nil,
		nil,
		nil,
		nil,
	)
	if err != nil {
		panic(err)
	}
	dataEntityModifier, err := entity.NewDataEntityModifier(*dataPropertiesModifer)
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
	sendOutputDataInstruction, err := instructionOutput.NewDataInstructionSendOutput()
	if err != nil {
		panic(err)
	}
	a := make([]instruction.DataInstruction, 3)
	a[0] = createEntity1DataInstruction
	a[1] = setXYStateForSquareInstruction
	a[2] = sendOutputDataInstruction

	dataArrayInstruction, err := instructioncontrol.NewDataInstructionArray(a)
	if err != nil {
		panic(err)
	}

	stages, err := phase.NewDataStage(dataArrayInstruction)
	if err != nil {
		panic(err)
	}

	turns, err := phase.NewDataTurn([]player.Id{}, []phase.DataStage{*stages})
	if err != nil {
		panic(err)
	}

	p, err := phase.NewDataPhase([]phase.DataTurn{*turns})
	if err != nil {
		panic(err)
	}
	return *p
}
