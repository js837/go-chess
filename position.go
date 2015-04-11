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
	castling  Castling // Castling rights
	enPassant int      // En passant square
	halfMoves int      // Number of moves since last capture or pawn advance
	fullMoves int      // Full move count - starts at 1, and is incremented after Black's move.
	score     int      // Board evaluation
}

type Board [64]Piece

const EMPTY_ENPASSANT int = -1

func GetRankFile(i int) (int, int) {
	return i / 8, i % 8
}

func SquareFromRankFile(rank, file int) int {
	return 8*rank + file
}

func (position *Position) GetBitBoard(piece Piece) uint64 {
	var bitboard uint64 = 0
	var mask uint64
	var i uint8
	for i = 0; i < 64; i++ {
		// If we have the correct piece at i we generate a bit string with 1 at posititon i
		// eg. i=4 00000...10000
		if position.board[i] == piece {
			mask = 0x01 << i
			bitboard ^= mask
		}
	}
	return bitboard
}

//func (p *Position) HashKey() int64 {

//}
