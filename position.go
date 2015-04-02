package main

// Colours
const (
	White = iota
	Black
)

// Piece types
const (
	NoPiece = iota << 1
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

// Pieces
const (
	WP = Pawn | White
	WN = Knight | White
	WB = Bishop | White
	WR = Rook | White
	WQ = Queen | White
	WK = King | White
	BP = Pawn | Black
	BN = Knight | Black
	BB = Bishop | Black
	BR = Rook | Black
	BQ = Queen | Black
	BK = King | Black
)

const OO = NoPiece // Alias for NoPiece
const N, E, S, W = +8, +1, -8, -1

type Piece uint8
type Colour uint8

func (p Piece) Colour() Colour { return Colour(p) & 0x01 }
func (p Piece) Type() Piece    { return Piece(p) &^ 0x01 }

func (c Colour) Switch() Colour {
	return Colour(c ^ 0x01)
}

type Castling struct {
	whiteKingside  bool
	whiteQueenside bool
	blackKingside  bool
	blackQueenside bool
}

type Position struct {
	board     Board    // 120 char representation of the board
	turn      Colour   // Whose turn is it?
	score     int      // Board evaluation
	castling  Castling // Castling rights
	enPassant int      // En passant square
}

type Board [64]Piece

func GetRankFile(i int) (int, int) {
	return i / 8, i % 8
}
