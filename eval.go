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

type EvalNode struct {
	position Position
	eval     float64
	children []EvalNode
}

func (node *EvalNode) GenerateTree(depth int) {

	if depth == 0 {
		return
	}
	//fmt.Println(depth)
	//fmt.Println(PositionToBoardFen(&node.position))

	for _, move := range node.position.GetMoves(node.position.turn) {

		newPosition := node.position.ApplyMove(move)

		//fmt.Println(PositionToBoardFen(&newPosition))

		eval := newPosition.PieceEval(node.position.turn)
		newNode := EvalNode{newPosition, eval, []EvalNode{}}

		//fmt.Println(move)
		//fmt.Println(eval)

		node.children = append(node.children, newNode)
		newNode.GenerateTree(depth - 1)
	}

}

//func minimax(node *Position, depth int, maximizingPlayer Colour) float64 {
//	if depth == 0 {
//		v := node.PieceEval(maximizingPlayer.Switch())
//		return v
//	}

//	//fmt.Println(depth, node)

//	moves := node.GetMoves(maximizingPlayer)
//	if len(moves) == 0 {
//		return node.PieceEval(maximizingPlayer.Switch())
//	}

//	var bestValue float64

//	if maximizingPlayer == White {
//		bestValue = -1000000
//		for _, move := range moves {
//			newPosition := node.ApplyMove(move)
//			val := minimax(&newPosition, depth-1, Black)
//			bestValue = max64(bestValue, val)
//		}
//		return bestValue
//	} else {
//		bestValue = +1000000
//		for _, move := range moves {
//			newPosition := node.ApplyMove(move)
//			val := minimax(&newPosition, depth-1, White)
//			bestValue = min64(bestValue, val)
//		}
//		return bestValue
//	}
//}

func BestMove(node *Position, depth int, maximizingPlayer Colour) float64 {
	if depth == 0 {
		v := node.PieceEval(maximizingPlayer.Switch())
		return v
	}

	moves := node.GetMoves(maximizingPlayer)
	if len(moves) == 0 {
		return node.PieceEval(maximizingPlayer)
	}

	var bestValue float64

}

func (position *Position) PieceEval(colour Colour) float64 {
	// Gives evaluation heurisitc of board w.r.t. colour

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
		if piece.Colour() == colour {
			eval += +1 * pieceValue
		} else {
			eval += -1 * pieceValue
		}
	}
	return eval
}
