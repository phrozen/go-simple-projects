package myGameUtils

//MyColor ...
type MyColor int

const (
	//White ...
	White MyColor = iota
	//Black ...
	Black
	//EmptyColor ...
	EmptyColor
)

//ColorToConsole ..
func ColorToConsole(c MyColor) string {
	if c == Black {
		return "B"
	}
	return "W"
}

//BlackOrWhiteSquare if both are odd or even then the square is black
func BlackOrWhiteSquare(x int, y int) MyColor {
	if ((x+1)%2 == 0 && (y+1)%2 == 0) || ((x+1)%2 == 1 && (y+1)%2 == 1) {
		return Black
	}
	return White
}

//ChessBoardH ...
const ChessBoardH = 8

//ChessBoardW ...
const ChessBoardW = 8

//PieceType ...
type PieceType int

const (
	//Pawn ...
	Pawn PieceType = iota
	//Knigth ...
	Knigth
	//Rock ...
	Rock
	//Bishop ...
	Bishop
	//King ...
	King
	//Queen ..
	Queen
	//EmptyPiece ...
	EmptyPiece
)

//PieceToLetter ...
func PieceToLetter(p PieceType) string {

	return [...]string{"p", "N", "R", "b", "K", "Q", "E"}[p]
}

//LetterToPiece ..
func LetterToPiece(s string) PieceType {
	m := map[string]PieceType{
		"p": Pawn,
		"N": Knigth,
		"R": Rock,
		"b": Bishop,
		"K": King,
		"Q": Queen,
		"E": EmptyPiece,
	}
	return m[s]
}
