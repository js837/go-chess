package main

import (
	"strconv"
	"strings"
)

func FromFen(fen string) Position {
	components := strings.Split(fen, " ")

	boardStr := components[0]
	turnStr := components[1]
	castlingStr := components[2]
	enPassantStr := components[3]
	halfMovesStr := components[4]
	fullMovesStr := components[5]

	board := boardFromFen(boardStr)
	turn := turnFromFen(turnStr)
	castling := castlingFromFen(castlingStr)
	enPassant := enPassantFromFen(enPassantStr)
	halfMoves, _ := strconv.Atoi(halfMovesStr)
	fullMoves, _ := strconv.Atoi(fullMovesStr)

	return Position{board, turn, castling, enPassant, halfMoves, fullMoves, 0}
}

var rankLookup = map[byte]int{'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7}
var fileLookup = map[byte]int{'1': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7}

func enPassantFromFen(enPassantStr string) int {
	if enPassantStr == "-" {
		return EMPTY_ENPASSANT
	}

	rank := rankLookup[enPassantStr[0]]
	file := fileLookup[enPassantStr[1]]

	return SquareFromRankFile(rank, file)
}

func castlingFromFen(castlingStr string) Castling {
	whiteKingside := strings.Contains(castlingStr, "K")
	whiteQueenside := strings.Contains(castlingStr, "Q")
	blackKingside := strings.Contains(castlingStr, "k")
	blackQueenside := strings.Contains(castlingStr, "q")

	return Castling{whiteKingside, whiteQueenside, blackKingside, blackQueenside}
}

func turnFromFen(turnStr string) Colour {
	if turnStr == "w" {
		return White
	} else {
		return Black
	}
}

func boardFromFen(boardStr string) Board {
	// boardStr eg. rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR
	board := Board{}
	var index int
	for k, rankStr := range strings.Split(boardStr, "/") {
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
		// Write blanks at the end of a row.
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
