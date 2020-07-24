package pieces

import (
	"github.com/santtana92/chess/myGameUtils"
)

//Piece ..
type Piece interface {
	GetColor() myGameUtils.MyColor
	GetType() myGameUtils.PieceType
	Draw()
}
