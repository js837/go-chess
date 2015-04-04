package main

//import "fmt"

const (
	KING_WEIGHT   = 2000
	QUEEN_WEIGHT  = 90
	ROOK_WEIGHT   = 50
	BISHOP_WEIGHT = 30
	KNIGHT_WEIGHT = 25
	PAWN_WEIGHT   = 10
)

func (position *Position) QuickEval() int64 {

	var eval int64 = 0
	var pieceValue int64

	for _, piece := range position.board {
		switch piece.Type() {
		case Pawn:
			pieceValue = PAWN_WEIGHT
		case Queen:
			pieceValue = QUEEN_WEIGHT
		case Rook:
			pieceValue = ROOK_WEIGHT
		case Bishop:
			pieceValue = BISHOP_WEIGHT
		case Knight:
			pieceValue = KNIGHT_WEIGHT
		case King:
			pieceValue = KING_WEIGHT
		default:
			pieceValue = 0
		}
		if piece.Colour() == White {
			eval += +1 * pieceValue
		} else {
			eval += -1 * pieceValue
		}
	}
	return eval
}
