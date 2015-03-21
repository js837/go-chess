package main

import (
	"fmt"
)

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

const N, E, S, W = +8, +1, -1, -8

type Piece uint8
type Colour uint8

func (p Piece) Color() int { return int(p) & 0x01 }
func (p Piece) Type() int  { return int(p) &^ 0x01 }

type Castling struct {
	whiteKingside  bool
	whiteQueenside bool
	blackKingside  bool
	blackQueenside bool
}
type Move struct {
	from      int
	to        int
	promotion Piece
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
	return i % 8, i / 8
}

func RookMoves(p *Position, from int) []Move {
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

	rank, file := GetRankFile(from)
	directions := [4]int{N, E, S, W}
	multLimits := [4]int{7 - rank, 7 - file, rank, file} // N E S W

	var moves = []Move{}
	for dir := 0; dir < 4; dir++ {
		for mult := 1; mult <= multLimits[dir]; mult++ {
			to := from + mult*directions[dir]
			if p.board[to] == NoPiece {
				// Empty space
				moves = append(moves, Move{dir, to, NoPiece})
			} else {
				// Hit a piece
				if p.board[to].Color() != myColour {
					moves = append(moves, Move{dir, to, NoPiece})
				}
				break
			}
		}
	}
	return moves
}

func QueenMoves(p *Position, i uint8) []Move {

	return []Move{}
}

func main() {
	fmt.Println(GetRankFile(63))
	p := Position{Board{}, White, 0, Castling{}, 0}

	for _, move := range RookMoves(&p, 18) {
		fmt.Println(move)

	}
}
