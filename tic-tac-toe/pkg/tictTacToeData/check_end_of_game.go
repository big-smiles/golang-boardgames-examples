package tictTacToeData

import (
	"errors"
	"fmt"
	"github.com/big-smiles/golang-boardgames/pkg/entity"
	"github.com/big-smiles/golang-boardgames/pkg/instruction"
	"github.com/big-smiles/golang-boardgames/pkg/player"
	resolveValueConstant "github.com/big-smiles/golang-boardgames/pkg/resolve_value/constant"
)

const (
	propertyNameSquares entity.NamePropertyId[[]entity.Id] = "squares"
)

type DataInstructionCheckEndOfGame struct {
}

func (d DataInstructionCheckEndOfGame) NewFromThisData() (instruction.Instruction, error) {
	return newInstructionCheckEndOfGame(d), nil
}

func NewDataInstructionCheckEndOfGame() *DataInstructionCheckEndOfGame {
	return &DataInstructionCheckEndOfGame{}
}

type InstructionCheckEndOfGame struct {
}

func newInstructionCheckEndOfGame(data DataInstructionCheckEndOfGame) *InstructionCheckEndOfGame {
	return &InstructionCheckEndOfGame{}

}
func (i InstructionCheckEndOfGame) Execute(ctx instruction.ExecutionContext) error {
	var predicate entity.Predicate = func(
		executionVariables entity.Entity,
		managerPropertiesId *entity.ManagerPropertyId,
		e entity.Entity,
	) (bool, error) {
		managerTypedPeropertyId, err := entity.GetManagerTypedPropertyId[int](managerPropertiesId)
		if err != nil {
			return false, err
		}
		variablePropertyId, err := managerTypedPeropertyId.GetId(IntPropertyNameState)
		if err != nil {
			return false, err
		}
		_, err = entity.GetValueFromEntity(e, variablePropertyId)
		if err != nil {
			var errorPropertyNotFound *entity.ErrorPropertyNotFound[int]
			if errors.As(err, &errorPropertyNotFound) {
				//here we can just ignore propertyNotFound
				return false, nil
			} else {
				return false, err
			}

		}
		return true, nil
	}
	err := ctx.Performer.Entity.FilterEntitiesIntoVariable(
		ctx.ExecutionVariables,
		predicate,
		propertyNameSquares,
	)
	if err != nil {
		return err
	}
	entityIds, err := instruction.ResolveValueResolver(
		ctx.ExecutionVariables,
		ctx.Performer.ValueResolver,
		resolveValueConstant.NewResolveValueFromVariable(propertyNameSquares),
	)
	if err != nil {
		return err
	}
	var board [3][3]int
	for _, entityId := range entityIds {
		e, err := ctx.Performer.Entity.GetById(entityId)
		if err != nil {
			return err
		}
		state, err := instruction.GetValueFromEntity(*ctx.Performer.Entity, *e, IntPropertyNameState)
		if err != nil {
			return err
		}
		x, err := instruction.GetValueFromEntity(*ctx.Performer.Entity, *e, IntPropertyNameX)
		if err != nil {
			return err
		}
		y, err := instruction.GetValueFromEntity(*ctx.Performer.Entity, *e, IntPropertyNameY)
		if err != nil {
			return err
		}
		if x < 0 || x > 2 {
			return errors.New("x must be between 0 and 2")
		}
		if y < 0 || y > 2 {
			return errors.New("y must be between 0 and 2")
		}
		board[x][y] = state

	}
	win, who, err := isVictory(board)
	if err != nil {
		return err
	}
	if win {
		if who == PlayerId1 {
			fmt.Println("PLAYER 1 WINS")
		} else {
			fmt.Println("PLAYER 2 WINS")
		}
	} else if isDraw(board) {
		fmt.Println("The Game is a draw")
	}
	return nil
}

type stateRow struct {
	one   int
	two   int
	three int
}

func isDraw(board [3][3]int) bool {
	return board[0][0] != StateEmpty &&
		board[0][1] != StateEmpty &&
		board[0][2] != StateEmpty &&
		board[1][0] != StateEmpty &&
		board[1][1] != StateEmpty &&
		board[1][2] != StateEmpty &&
		board[2][0] != StateEmpty &&
		board[2][1] != StateEmpty &&
		board[2][2] != StateEmpty
}
func isVictory(board [3][3]int) (win bool, who player.Id, err error) {
	rows := []stateRow{
		//first horizontal
		{board[0][0], board[0][1], board[0][2]},
		//second horizontal
		{board[1][0], board[1][1], board[1][2]},
		//thir horizontal
		{board[2][0], board[2][1], board[2][2]},
		//first vertical
		{board[0][0], board[1][0], board[2][0]},
		//second vertical
		{board[0][1], board[1][1], board[2][1]},
		//third vertical
		{board[0][2], board[1][2], board[2][2]},
		//first diagonal
		{board[0][0], board[1][1], board[2][2]},
		//second diagonal
		{board[0][2], board[1][1], board[2][0]},
	}
	for _, row := range rows {
		win, who, err = isVictoryThreeSquares(row.one, row.two, row.three)
		if err != nil {
			return false, who, err
		}
		if win {
			return true, who, nil
		}
	}
	return false, who, nil
}
func isVictoryThreeSquares(one int, two int, three int) (win bool, who player.Id, err error) {
	var zero player.Id
	if one == StateEmpty || two == StateEmpty || three == StateEmpty {
		return false, zero, nil
	}
	if one == two && two == three {
		if one == StatePlayer1 {
			return true, PlayerId1, nil
		} else if one == StatePlayer2 {
			return true, PlayerId2, nil
		} else {
			return false, zero, fmt.Errorf("invalid state %d", one)
		}
	}

	return false, zero, nil
}
