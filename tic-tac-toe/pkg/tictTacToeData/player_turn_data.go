package tictTacToeData

import (
	"errors"
	"fmt"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	instructionControl "github.com/big-smiles/golang-boardgames/pkg/instructions/control"
	instructionEntity "github.com/big-smiles/golang-boardgames/pkg/instructions/entity"
	instructionEntityModifier "github.com/big-smiles/golang-boardgames/pkg/instructions/entity_modifier"
	instructionInteraction "github.com/big-smiles/golang-boardgames/pkg/instructions/interaction"
	instructionOutput "github.com/big-smiles/golang-boardgames/pkg/instructions/output"
	instruction_phase "github.com/big-smiles/golang-boardgames/pkg/instructions/phase"
	"github.com/big-smiles/golang-boardgames/pkg/interaction"
	"github.com/big-smiles/golang-boardgames/pkg/phaseData"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
	ValueModifierCommon "github.com/big-smiles/golang-boardgames/pkg/value_modifier/common"
)

func GetPlayersTurnPhase() *phaseData.DataPhase {
	dataPhase := phaseData.DataPhase{
		Name: PhasePlayerTurns,
		Turns: []phaseData.DataTurn{
			getPlayerTurn(PlayerId1),
			getPlayerTurn(PlayerId2),
			{
				ActivePlayers: []player.Id{},
				Stages: []phaseData.DataStage{
					{
						Instructions: instructionControl.NewDataInstructionArray(
							instruction_phase.NewDataInstructionSetNextPhase(
								resolveValueConstant.NewResolveConstant(PhasePlayerTurns),
							),
						),
					},
				},
			},
		},
	}
	return &dataPhase
}
func getPlayerTurn(playerId player.Id) phaseData.DataTurn {
	fmt.Printf("\ngetPlayerTurn %s\n", playerId)
	var variablePropertyName entity.NamePropertyId[[]entity.Id] = "filtered_entities"
	resolveEntitiesIds := resolveValueConstant.NewResolveValueFromVariable[[]entity.Id](variablePropertyName)
	playerInteraction, err := interaction.NewDataAvailableInteraction(
		playerId, resolveEntitiesIds,
		1,
		2,
	)
	if err != nil {
		panic(err)
	}
	turn := phaseData.DataTurn{
		ActivePlayers: []player.Id{playerId},
		Stages: []phaseData.DataStage{
			{
				Instructions: instructionControl.NewDataInstructionArray(
					instructionEntity.NewDataInstructionFilterEntities(
						selectEmptySquares,
						variablePropertyName,
					),
					instructionInteraction.NewDataClearAvailableInteraction(),
					instructionInteraction.NewDataAvailableInteractionData(
						*playerInteraction,
						instructionControl.NewDataInstructionArray(
							instructionEntityModifier.NewDataInstructionAddEntityModifierWithResolvedTarget(
								resolveValueConstant.NewResolveValueFromVariable[[]entity.Id](instruction.SelectedEntities),
								*getModifierSelected(playerId),
							),
							NewDataInstructionCheckEndOfGame(),
							instructionOutput.NewDataInstructionSendOutput(),
						),
					),
					instructionInteraction.NewDataWaitForInteractionData(),
				),
			},
		},
	}
	return turn
}
func selectEmptySquares(
	_ entity.Entity,
	managerPropertyId *entity.ManagerPropertyId,
	e entity.Entity,
) (bool, error) {
	managerTypedPeropertyId, err := entity.GetManagerTypedPropertyId[int](managerPropertyId)
	if err != nil {
		return false, err
	}
	variablePropertyId, err := managerTypedPeropertyId.GetId(IntPropertyNameState)
	if err != nil {
		return false, err
	}
	state, err := entity.GetValueFromEntity(e, variablePropertyId)
	if err != nil {
		var errorPropertyNotFound *entity.ErrorPropertyNotFound[int]
		if errors.As(err, &errorPropertyNotFound) {
			//here we can just ignore propertyNotFound
			return false, nil
		} else {
			return false, err
		}

	}
	canBeSelected := state == StateEmpty
	return canBeSelected, nil

}
func getModifierSelected(playerId player.Id) *entity.DataModifier {
	var value int
	if playerId == PlayerId1 {
		value = StatePlayer1
	} else {
		value = StatePlayer2
	}
	valueModifier, err := ValueModifierCommon.NewDataModifierSetValue[int](resolveValueConstant.NewResolveConstant(value))
	if err != nil {
		panic(err)
	}
	data, err := entity.NewDataEntityModifier(
		entity.DataPropertiesModifier{
			IntModifiers: entity.MapDataModifierProperties[int]{
				IntPropertyNameState: valueModifier,
			},
		})
	if err != nil {
		panic(err)
	}
	return data
}
