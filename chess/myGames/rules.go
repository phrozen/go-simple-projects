package myRules

import (
	"github.com/santtana92/chess/myGameUtils"
)

//GameRules ...
type GameRules interface {
	TurnMove(move Movement) (bool, error)
	PlayerLoose(p myGameUtils.MyColor) bool
}

//Movement ...
type Movement interface {
	Parse(move interface{}) (x int, y int, piece myGameUtils.PieceType)
}
