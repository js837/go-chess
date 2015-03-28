package main

import (
	"strconv"
	"strings"
)

//func PositionFromFen(fen string) Position {
//	// Examples:
//	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
//	// rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2

//	components := strings.Split("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2", " ")
//}

func getPiece(char rune) Piece {

	switch char {
	case 'P':
		return WP
	case 'N':
		return WN
	case 'B':
		return WB
	case 'R':
		return WR
	case 'Q':
		return WQ
	case 'K':
		return WK

	case 'p':
		return BP
	case 'n':
		return BN
	case 'b':
		return BB
	case 'r':
		return BR
	case 'q':
		return BQ
	case 'k':
		return BK

	default:
		return NoPiece
	}
}

func PositionFromBoardFen(boardFen string) Board {
	// boardFen eg. rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR
	// TODO: Add error checking.
	board := Board{}
	var index int
	for k, rankStr := range strings.Split(boardFen, "/") {
		index = (7 - k) * 8
		for _, char := range rankStr {
			piece := getPiece(char)
			if piece == NoPiece {
				n, _ := strconv.Atoi(string(char))
				for m := 0; m < n; m++ {
					board[index] = NoPiece
					index++
				}
			} else {
				board[index] = piece
				index++
			}
		}
	}
	return board
}

//func PositionToFen(position *Position) string {
//	// Examples:
//	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
//	// rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2

//	components := strings.Split("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2", " ")

//}
