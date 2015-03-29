package main

type Move struct {
	from      int
	to        int
	promotion Piece
}

func (p *Position) GetMoves() []Move {
	var moves = []Move{}
	for from, piece := range p.board {
		if piece.Color() == White {
			switch piece.Type() {
			case Pawn:
				moves = append(moves, PawnMoves(p, from)...)
			case Queen:
				moves = append(moves, QueenMoves(p, from)...)
			case Rook:
				moves = append(moves, RookMoves(p, from)...)
			case Bishop:
				moves = append(moves, BishopMoves(p, from)...)
			}
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

func PawnMoves(p *Position, from int) []Move {
	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

	rank, _ := GetRankFile(from)

	var moves = []Move{}
	var localFirstRank, localPromoRank, localN, localE, localW int

	if myColour == White {
		localFirstRank, localPromoRank = 1, 6
		localN, localE, localW = N, E, W
	} else {
		localFirstRank, localPromoRank = 6, 1
		localN, localE, localW = -N, -E, -W
	}

	// Promotion
	if to := from + localN; rank == localPromoRank && p.board[to] == NoPiece {
		moves = append(moves, Move{from, to, Queen}) // Only queen for now.
	} else {
		// Single move
		if to := from + localN; p.board[to] == NoPiece {
			moves = append(moves, Move{from, to, NoPiece})

			// Double move
			if to = to + localN; rank == localFirstRank && p.board[to] == NoPiece {
				moves = append(moves, Move{from, to, NoPiece})
			}

		}

	}

	// Take NE
	if to := from + localN + localE; p.board[to] != NoPiece && p.board[to].Color() != myColour {
		moves = append(moves, Move{from, to, NoPiece})
	}
	// Take NW
	if to := from + localN + localW; p.board[to] != NoPiece && p.board[to].Color() != myColour {
		moves = append(moves, Move{from, to, NoPiece})
	}

	return moves
}
