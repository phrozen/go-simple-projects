package pieces

import (
	"fmt"
	"github.com/santtana92/chess/myGameUtils"
)

type baseChessPiece struct {
	color myGameUtils.MyColor
	pType myGameUtils.PieceType
}

func (b *baseChessPiece) GetColor() myGameUtils.MyColor {
	return b.color
}

func (b *baseChessPiece) GetType() myGameUtils.PieceType {
	return b.pType
}

//consolePiece ...
type consoleChessPiece struct {
	baseChessPiece
}

func (p *consoleChessPiece) Draw() {
	fmt.Print(myGameUtils.ColorToConsole(p.GetColor()) + myGameUtils.PieceToLetter(p.pType))
}

//NewConsolePiece ...
func NewConsolePiece(c myGameUtils.MyColor, p myGameUtils.PieceType) Piece {
	return &consoleChessPiece{
		baseChessPiece: baseChessPiece{
			color: c,
			pType: p}}
}

var (
	emptyConsolePiece Piece
)

//GetConsoleEmptyPiece ...
func GetConsoleEmptyPiece() Piece {
	if emptyConsolePiece == nil {
		emptyConsolePiece = NewConsolePiece(myGameUtils.EmptyColor, myGameUtils.EmptyPiece)
	}
	return emptyConsolePiece
}
