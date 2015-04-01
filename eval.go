package main

const (
	KING_WEIGHT   = 2000
	QUEEN_WEIGHT  = 90
	ROOK_WEIGHT   = 50
	BISHOP_WEIGHT = 30
	KNIGHT_WEIGHT = 25
	PAWN_WEIGHT   = 10
)

func (position *Position) PieceEval() float64 {

	var eval float64 = 0
	var pieceValue float64

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
		}

		if piece.Colour() == position.turn {
			eval += +1 * pieceValue
		} else {
			eval += -1 * pieceValue
		}

	}
	return eval
}
