package main

//import "fmt"

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
			case Knight:
				moves = append(moves, KnightMoves(p, from)...)
			case King:
				moves = append(moves, KingMoves(p, from)...)
			}
		}
	}
	return moves
}

func (p *Position) ApplyMove(m Move) Position {

	newBoard := Board(p.board)

	if m.promotion != NoPiece {
		newBoard[m.to] = newBoard[m.promotion]
	} else {
		newBoard[m.to] = newBoard[m.from]
	}
	newBoard[m.from] = NoPiece

	newPosition := Position{
		newBoard,
		^p.turn,
		p.score,
		p.castling,
		p.enPassant,
	}
	return newPosition
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
	if a <= b {
		return a
	} else {
		return b
	}

}

func KnightMoves(p *Position, from int) []Move {
	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()
	var moves = []Move{}

	rank, file := GetRankFile(from)

	dirs := [4]int{N, E, S, W}
	multLimits := [4]int{7 - rank, 7 - file, rank, file} // N E S W

	for i, dir := range dirs {
		var a [2]int
		if dir == N || dir == S {
			a = [2]int{1, 3}
		} else {
			a = [2]int{0, 2}
		}
		if multLimits[i] >= 2 {
			for _, j := range a {
				if multLimits[j] >= 1 {
					to := from + 2*dir + dirs[j]
					if p.board[to] == NoPiece || p.board[to].Color() != myColour {
						moves = append(moves, Move{from, to, NoPiece})
					}
				}
			}
		}
	}
	return moves
}

func KingMoves(p *Position, from int) []Move {

	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

	rank, file := GetRankFile(from)
	directions := [8]int{N, E, S, W, N + E, E + S, S + W, W + N}
	multLimits := [8]int{7 - rank,
		7 - file,
		rank,
		file,
		min(7-rank, 7-file),
		min(7-file, rank),
		min(rank, file),
		min(7-rank, file),
	}

	var moves = []Move{}
	for dir, _ := range directions {
		for mult := 1; mult <= min(multLimits[dir], 1); mult++ {
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

	rank, file := GetRankFile(from)

	var moves = []Move{}
	var localFirstRank, localPromoRank, localN int

	if myColour == White {
		localFirstRank, localPromoRank = 1, 6
		localN = N
	} else {
		localFirstRank, localPromoRank = 6, 1
		localN = S
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
	if to := from + localN + E; file != 7 && p.board[to] != NoPiece && p.board[to].Color() != myColour {
		moves = append(moves, Move{from, to, NoPiece})
	}
	// Take NW
	if to := from + localN + W; file != 0 && p.board[to] != NoPiece && p.board[to].Color() != myColour {
		moves = append(moves, Move{from, to, NoPiece})
	}

	return moves
}
