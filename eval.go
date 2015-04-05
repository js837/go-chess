package main

import "fmt"

const (
	KING_WEIGHT   = 2000
	QUEEN_WEIGHT  = 90
	ROOK_WEIGHT   = 50
	BISHOP_WEIGHT = 30
	KNIGHT_WEIGHT = 25
	PAWN_WEIGHT   = 10
)

const EVAL_MAX int = 10000000

func (p *Position) GetBestMove(depth int, colour Colour) Position {
	var bestEval int
	var bestPosition Position

	if colour == White {
		bestEval = -EVAL_MAX
	} else {
		bestEval = +EVAL_MAX
	}

	for _, move := range p.GetMoves(colour) {
		newPosition := p.ApplyMove(move)
		eval := minimax(&newPosition, depth, colour.Switch())
		if colour == White {
			if eval > bestEval {
				bestEval = eval
				bestPosition = newPosition
			}
		} else {
			if eval < bestEval {
				bestEval = eval
				bestPosition = newPosition
			}
		}
	}
	fmt.Println(bestEval)
	return bestPosition
}

func minimax(p *Position, depth int, colour Colour) int {
	if depth == 0 {
		return p.QuickEval()
	}

	moves := p.GetMoves(colour)
	if len(moves) == 0 {
		return p.QuickEval()
	}

	if colour == White {
		v := -EVAL_MAX
		for _, move := range moves {
			child := p.ApplyMove(move)
			v = max(v, minimax(&child, depth-1, Black))
		}
		return v
	} else {
		v := EVAL_MAX
		for _, move := range moves {
			child := p.ApplyMove(move)
			v = min(v, minimax(&child, depth-1, White))
		}
		return v
	}
	return 0
}

func (position *Position) QuickEval() int {

	var eval int = 0
	var pieceValue int

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
