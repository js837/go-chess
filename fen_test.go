package main

import (
	"testing"
)

func TestPositionFromBoardFen(t *testing.T) {

	board := PositionFromBoardFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR").board
	starting_board := Board{
		WR, WN, WB, WQ, WK, WB, WN, WR,
		WP, WP, WP, WP, WP, WP, WP, WP,
		OO, OO, OO, OO, OO, OO, OO, OO,
		OO, OO, OO, OO, OO, OO, OO, OO,
		OO, OO, OO, OO, OO, OO, OO, OO,
		OO, OO, OO, OO, OO, OO, OO, OO,
		BP, BP, BP, BP, BP, BP, BP, BP,
		BR, BN, BB, BQ, BK, BB, BN, BR,
	}

	if board != starting_board {
		t.Error("Starting board incorrect.")
	}
}
