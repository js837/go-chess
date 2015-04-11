package main

const (
	KING_WEIGHT   = EVAL_MAX
	QUEEN_WEIGHT  = 900
	ROOK_WEIGHT   = 500
	BISHOP_WEIGHT = 300
	KNIGHT_WEIGHT = 250
	PAWN_WEIGHT   = 100
)

const EVAL_MAX int = 10000000

func (position *Position) QuickEval() int {

	var eval int = 0
	var pieceValue int

	// Obviously inefficient.
	eval += +1*len(position.GetMoves(White)) + -1*len(position.GetMoves(Black))

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
