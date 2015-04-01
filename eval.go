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

type EvalNode struct {
	position Position
	eval     float64
	children []EvalNode
}

func (node *EvalNode) GenerateTree(depth int) {

	if depth == 0 {
		return
	}
	fmt.Println(depth)
	for _, move := range node.position.GetMoves(node.position.turn) {

		newPosition := node.position.ApplyMove(move)
		eval := newPosition.PieceEval()
		newNode := EvalNode{newPosition, eval, []EvalNode{}}

		//fmt.Println(move)
		//fmt.Println(eval)

		node.children = append(node.children, newNode)
		newNode.GenerateTree(depth - 1)
	}

}

func (position *Position) PieceEval() float64 {
	// Requires a channel of pieces to work on

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
		default:
			pieceValue = 0
		}
		if piece.Colour() == position.turn {
			eval += +1 * pieceValue
		} else {
			eval += -1 * pieceValue
		}
	}
	return eval
}
