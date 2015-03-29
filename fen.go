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

func getChar(piece Piece) rune {

	switch piece {
	case WP:
		return 'P'
	case WN:
		return 'N'
	case WB:
		return 'B'
	case WR:
		return 'R'
	case WQ:
		return 'Q'
	case WK:
		return 'K'

	case BP:
		return 'p'
	case BN:
		return 'n'
	case BB:
		return 'b'
	case BR:
		return 'r'
	case BQ:
		return 'q'
	case BK:
		return 'k'
	default:
		return '-'
	}

}

func PositionFromBoardFen(boardFen string) Position {
	// boardFen eg. rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR
	// TODO: Add error checking.
	position := Position{}
	var index int
	for k, rankStr := range strings.Split(boardFen, "/") {
		index = (7 - k) * 8
		for _, char := range rankStr {
			piece := getPiece(char)
			if piece == NoPiece {
				n, _ := strconv.Atoi(string(char))
				for m := 0; m < n; m++ {
					position.board[index] = NoPiece
					index++
				}
			} else {
				position.board[index] = piece
				index++
			}
		}
	}
	return position
}

func PositionToBoardFen(p *Position) string {
	var blankCount = 0
	boardFen := ""
	for i := 56; i >= 0; i -= 8 {

		for j := 0; j < 8; j++ {
			piece := p.board[i+j]
			if piece == NoPiece {
				blankCount++
			} else {
				// Write blanks
				if blankCount > 0 {
					boardFen += strconv.Itoa(blankCount)
				}

				// Write piece
				boardFen += string(getChar(piece))

				// Reset
				blankCount = 0
			}
		}
		if blankCount > 0 {
			boardFen += strconv.Itoa(blankCount)
			blankCount = 0
		}
		if i != 0 {
			boardFen += "/"
		}
	}
	return boardFen
}

//func PositionToFen(position *Position) string {
//	// Examples:
//	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
//	// rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2

//	components := strings.Split("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c6 0 2", " ")

//}
