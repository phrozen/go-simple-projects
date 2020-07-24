package myRules

import (
	"github.com/santtana92/chess/myBoards"
	"github.com/santtana92/chess/myGameUtils"
)

//ChessRules ...
type ChessRules struct {
	board myBoards.ChessBoard
	turn  myGameUtils.MyColor
}

//TurnMove ...
func (c *ChessRules) TurnMove(move Movement) (bool, error) {
	return true, nil
}

//PlayerLoose ...
func (c *ChessRules) PlayerLoose(myGameUtils.MyColor) bool {
	return false
}

//Check ...
func (c *ChessRules) Check() bool {
	return false
}

//CheckMate ...
func (c *ChessRules) CheckMate() bool {
	return false
}

//ConsoleMovement ...
type ConsoleMovement struct {
}

//Parse ...
func (m *ConsoleMovement) Parse(move interface{}) (x int, y int, piece myGameUtils.PieceType) {
	move = move.(string)
	fmt.Printf("%s", s)
	return 0, 0, myGameUtils.EmptyPiece
}
