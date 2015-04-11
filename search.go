package main

import "fmt"

func (p *Position) GetBestMove(depth int, colour Colour) (Move, Position) {
	var bestEval int
	var bestPosition Position
	var bestMove Move

	if colour == White {
		bestEval = -EVAL_MAX
	} else {
		bestEval = +EVAL_MAX
	}

	var evalHits int = 0

	for _, move := range p.GetMoves(colour) {
		newPosition := p.ApplyMove(move)

		eval := alphabeta(&newPosition, depth, colour.Switch(), -EVAL_MAX, +EVAL_MAX, &evalHits)
		//eval := iterDeep(&newPosition, depth, colour.Switch())

		if colour == White {
			if eval > bestEval {
				bestEval = eval
				bestPosition = newPosition
				bestMove = move
			}
		} else {
			if eval < bestEval {
				bestEval = eval
				bestPosition = newPosition
				bestMove = move
			}
		}
	}
	fmt.Println("Best eval:", bestEval)
	fmt.Println("evalHits:", evalHits)
	return bestMove, bestPosition
}

type TreeKey struct {
	position Position
	move     Move
	cutoff   bool
}

func iterDeep(newPosition *Position, maxDepth int, colour Colour) int {
	var evalHits int = 0
	var eval int

	for depth := 0; depth <= maxDepth; depth++ {
		eval = alphabeta(newPosition, depth, colour, -EVAL_MAX, +EVAL_MAX, &evalHits)
	}

	return eval
}

func alphabeta(p *Position, depth int, colour Colour, alpha int, beta int, evalHits *int) int {
	if depth == 0 {
		*evalHits++
		return p.QuickEval()
	}

	moves := p.GetMoves(colour)

	if len(moves) == 0 {
		*evalHits++
		return p.QuickEval()
	}

	if colour == White {
		v := -EVAL_MAX
		for _, move := range moves {
			child := p.ApplyMove(move)

			v = max(v, alphabeta(&child, depth-1, Black, alpha, beta, evalHits))
			alpha = max(alpha, v)
			if beta <= alpha {
				break // beta cut off
			}
		}
		return v
	} else {
		v := EVAL_MAX
		for _, move := range moves {
			child := p.ApplyMove(move)
			v = min(v, alphabeta(&child, depth-1, White, alpha, beta, evalHits))
			beta = min(beta, v)

			if beta <= alpha {
				break // alpha cut off
			}
		}
		return v
	}
	return 0
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
