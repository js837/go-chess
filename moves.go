package main

type Move struct {
	from      int
	to        int
	promotion Piece
}

func (p *Position) GetMoves() []Move {
	var moves = []Move{}
	for from, piece := range p.board {
		switch piece.Type() {
		case Queen:
			moves = append(moves, QueenMoves(p, from)...)
		case Rook:
			moves = append(moves, RookMoves(p, from)...)
		case Bishop:
			moves = append(moves, BishopMoves(p, from)...)
		}

	}
	return moves
}

func QueenMoves(p *Position, from int) []Move {
	return append(RookMoves(p, from), BishopMoves(p, from)...)
}

func RookMoves(p *Position, from int) []Move {
	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

	rank, file := GetRankFile(from)
	directions := [4]int{N, E, S, W}
	multLimits := [4]int{7 - rank, 7 - file, rank, file} // N E S W

	var moves = []Move{}
	for dir := 0; dir < 4; dir++ {
		for mult := 1; mult <= multLimits[dir]; mult++ {
			to := from + mult*directions[dir]
			if p.board[to] == NoPiece {
				// Empty space
				moves = append(moves, Move{from, to, NoPiece})
			} else {
				// Hit a piece
				if p.board[to].Color() != myColour {
					moves = append(moves, Move{from, to, NoPiece})
				}
				break
			}
		}
	}
	return moves
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func BishopMoves(p *Position, from int) []Move {
	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

	rank, file := GetRankFile(from)
	directions := [4]int{N + E, E + S, S + W, W + N}
	multLimits := [4]int{
		min(7-rank, 7-file),
		min(7-file, rank),
		min(rank, file),
		min(7-rank, file),
	}

	var moves = []Move{}
	for dir := 0; dir < 4; dir++ {
		for mult := 1; mult <= multLimits[dir]; mult++ {
			to := from + mult*directions[dir]
			if p.board[to] == NoPiece {
				// Empty space
				moves = append(moves, Move{from, to, NoPiece})
			} else {
				// Hit a piece
				if p.board[to].Color() != myColour {
					moves = append(moves, Move{from, to, NoPiece})
				}
				break
			}
		}
	}
	return moves
}
