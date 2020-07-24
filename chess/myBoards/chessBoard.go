package myBoards

import (
	"errors"
	"fmt"
	"github.com/santtana92/chess/myGameUtils"
	"github.com/santtana92/chess/pieces"
)

type consoleChessSquare struct {
	baseSquare
}

func (c *consoleChessSquare) Draw() {
	if c.Empty() {
		fmt.Printf("%-3s|", myGameUtils.ColorToConsole(c.GetColor()))
	} else {
		fmt.Printf("%s", myGameUtils.ColorToConsole(c.GetColor()))
		c.Content().Draw()
		fmt.Print("|")
	}
}

//NewConsoleChessSquare ...
func NewConsoleChessSquare(c myGameUtils.MyColor) Square {
	return &consoleChessSquare{
		baseSquare: baseSquare{
			color:    c,
			optPiece: pieces.GetConsoleEmptyPiece()}}
}

//ChessBoard ...
type ChessBoard struct {
	squares [][]Square
}

//GetBoard ..
func (c *ChessBoard) GetBoard() [][]Square {
	return c.squares
}

//ConsoleChessBoard ...
type ConsoleChessBoard struct {
	ChessBoard
}

//Draw ...
func (b *ConsoleChessBoard) Draw() {
	for i := myGameUtils.ChessBoardH - 1; i >= 0; i-- {
		for j := 0; j < myGameUtils.ChessBoardW; j++ {
			b.squares[i][j].Draw()
		}
		fmt.Println()
	}
}

//InitBoard ...
func (b *ConsoleChessBoard) InitBoard() (bool, error) {
	if b.squares == nil {
		return false, errors.New("Board is empty. Cannot init a empty board")
	}
	otherPieces := []myGameUtils.PieceType{myGameUtils.Rock, myGameUtils.Knigth,
		myGameUtils.Bishop, myGameUtils.Queen, myGameUtils.King, myGameUtils.Bishop,
		myGameUtils.Knigth, myGameUtils.Rock}
	for _, color := range []myGameUtils.MyColor{myGameUtils.Black, myGameUtils.White} {
		var pawnC int
		var otherC int
		if color == myGameUtils.White {
			pawnC = 1
			otherC = 0
		} else {
			pawnC = 6
			otherC = 7
		}
		for i := 0; i < myGameUtils.ChessBoardW; i++ {
			b.squares[pawnC][i].SetContent(pieces.NewConsolePiece(color, myGameUtils.Pawn))
			b.squares[otherC][i].SetContent(pieces.NewConsolePiece(color, otherPieces[i]))
		}
	}

	return true, nil
}

//ResetBoard ...
func (b *ConsoleChessBoard) ResetBoard() (bool, error) {
	return true, nil
}

//NewConsoleCheesBoard Coments to avoid // WARNING:
func NewConsoleCheesBoard() *ConsoleChessBoard {
	chessConsoleSquares := make([][]Square, myGameUtils.ChessBoardH)
	for i := 0; i < myGameUtils.ChessBoardH; i++ {
		chessConsoleSquares[i] = make([]Square, myGameUtils.ChessBoardW)
		for j := 0; j < myGameUtils.ChessBoardW; j++ {
			chessConsoleSquares[i][j] = NewConsoleChessSquare(myGameUtils.BlackOrWhiteSquare(i, j))
		}
	}
	return &ConsoleChessBoard{
		ChessBoard: ChessBoard{
			squares: chessConsoleSquares}}
}
