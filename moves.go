package main

import (
	"fmt"
)

type Move struct {
	from      int
	to        int
	promotion Piece
}

func (p *Position) GetMoves(colour Colour) {

	moves := make(chan Move)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-moves
			if more {
				fmt.Println(j)
			} else {
				done <- true
				return
			}
		}
	}()

	for from, piece := range p.board {
		go func(from int, piece Piece) {
			if piece.Color() == colour {
				switch piece.Type() {
				case Pawn:
					PawnMoves(p, from, moves)
				case Queen:
					QueenMoves(p, from, moves)
				case Rook:
					RookMoves(p, from, moves)
				case Bishop:
					BishopMoves(p, from, moves)
				case Knight:
					KnightMoves(p, from, moves)
				case King:
					KingMoves(p, from, moves)
				}
			}
		}(from, piece)
	}

	<-done

}

func (p *Position) ApplyMove(m Move) Position {

	newBoard := Board(p.board)

	if m.promotion != NoPiece {
		newBoard[m.to] = m.promotion
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

func QueenMoves(p *Position, from int, moves chan Move) {
	RookMoves(p, from, moves)
	BishopMoves(p, from, moves)
}

func RookMoves(p *Position, from int, moves chan Move) {
	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

	rank, file := GetRankFile(from)
	directions := [4]int{N, E, S, W}
	multLimits := [4]int{7 - rank, 7 - file, rank, file} // N E S W

	for dir := 0; dir < 4; dir++ {
		for mult := 1; mult <= multLimits[dir]; mult++ {
			to := from + mult*directions[dir]
			if p.board[to] == NoPiece {
				// Empty space
				moves <- Move{from, to, NoPiece}
			} else {
				// Hit a piece
				if p.board[to].Color() != myColour {
					moves <- Move{from, to, NoPiece}
				}
				break
			}
		}
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}

}

func KnightMoves(p *Position, from int, moves chan Move) {
	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

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
						moves <- Move{from, to, NoPiece}
					}
				}
			}
		}
	}
}

func KingMoves(p *Position, from int, moves chan Move) {

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

	for dir, _ := range directions {
		for mult := 1; mult <= min(multLimits[dir], 1); mult++ {
			to := from + mult*directions[dir]
			if p.board[to] == NoPiece {
				// Empty space
				moves <- Move{from, to, NoPiece}
			} else {
				// Hit a piece
				if p.board[to].Color() != myColour {
					moves <- Move{from, to, NoPiece}
				}
				break
			}
		}
	}
}

func BishopMoves(p *Position, from int, moves chan Move) {
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

	for dir := 0; dir < 4; dir++ {
		for mult := 1; mult <= multLimits[dir]; mult++ {
			to := from + mult*directions[dir]
			if p.board[to] == NoPiece {
				// Empty space
				moves <- Move{from, to, NoPiece}
			} else {
				// Hit a piece
				if p.board[to].Color() != myColour {
					moves <- Move{from, to, NoPiece}
				}
				break
			}
		}
	}
}

func PawnMoves(p *Position, from int, moves chan Move) {
	// Get the possible moves o
	// TODO: Get the colour from the position
	myColour := p.board[from].Color()

	rank, file := GetRankFile(from)

	var localFirstRank, localPromoRank, localFinalRank, localN int

	if myColour == White {
		localFirstRank, localPromoRank, localFinalRank = 1, 6, 7
		localN = N
	} else {
		localFirstRank, localPromoRank, localFinalRank = 6, 1, 0
		localN = S
	}

	// Not legal position, but included for safety
	if rank == localFinalRank {
		return
	}

	// Promotion (only Queen for now)
	var promotion Piece
	if rank == localPromoRank {
		promotion = Queen
	} else {
		promotion = NoPiece
	}

	// Single move
	if to := from + localN; p.board[to] == NoPiece {
		moves <- Move{from, to, promotion}

		// Double move
		if to = to + localN; rank == localFirstRank && p.board[to] == NoPiece {
			moves <- Move{from, to, NoPiece}
		}

	}
	// Take NE
	if to := from + localN + E; file != 7 && p.board[to] != NoPiece && p.board[to].Color() != myColour {
		moves <- Move{from, to, promotion}
	}
	// Take NW
	if to := from + localN + W; file != 0 && p.board[to] != NoPiece && p.board[to].Color() != myColour {
		moves <- Move{from, to, promotion}
	}
}
