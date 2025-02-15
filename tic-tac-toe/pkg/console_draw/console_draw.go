package console_draw

import (
	"fmt"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/board"
	"github.com/big-smiles/golang-boardgame-examples/tic-tac-toe/pkg/tictTacToeData"
)

func Draw(b board.Board) {
	square00 := getCharacter(b.Squares[0][0])
	square01 := getCharacter(b.Squares[0][1])
	square02 := getCharacter(b.Squares[0][2])
	square10 := getCharacter(b.Squares[1][0])
	square11 := getCharacter(b.Squares[1][1])
	square12 := getCharacter(b.Squares[1][2])
	square20 := getCharacter(b.Squares[2][0])
	square21 := getCharacter(b.Squares[2][1])
	square22 := getCharacter(b.Squares[2][2])

	println("")
	println("_______")
	println(fmt.Sprintf("|%s|%s|%s|", square00, square01, square02))
	println("_______")
	println(fmt.Sprintf("|%s|%s|%s|", square10, square11, square12))
	println("_______")
	println(fmt.Sprintf("|%s|%s|%s|", square20, square21, square22))
	println("_______")
}
func getCharacter(state int) string {
	switch state {
	case tictTacToeData.StateEmpty:
		return " "
	case tictTacToeData.StatePlayer1:
		return "X"
	case tictTacToeData.StatePlayer2:
		return "O"
	default:
		panic("invalid state")
	}

}
