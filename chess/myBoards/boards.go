package myBoards

import (
	"github.com/santtana92/chess/myGameUtils"
	"github.com/santtana92/chess/pieces"
)

//Square ...
type Square interface {
	Empty() bool
	Content() pieces.Piece
	GetColor() myGameUtils.MyColor
	SetContent(p pieces.Piece)
	Draw()
}

type baseSquare struct {
	color    myGameUtils.MyColor
	optPiece pieces.Piece
}

//Empty ...
func (s *baseSquare) Empty() bool {
	return s.optPiece.GetType() == myGameUtils.EmptyPiece
}

//Content ...
func (s *baseSquare) Content() pieces.Piece {
	return s.optPiece
}

//SetContent ...
func (s *baseSquare) SetContent(p pieces.Piece) {
	s.optPiece = p
}

//GetColor
func (s *baseSquare) GetColor() myGameUtils.MyColor {
	return s.color
}

//Board ...
type Board interface {
	Draw()
	GetBoard() [][]Square
	InitBoard() (bool, error)
	ResetBoard() (bool, error)
}
