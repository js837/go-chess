package main

import (
	"testing"
)

func TestPositionFromBoardFen(t *testing.T) {

	pos := FromFen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b Q b2 23 10")
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

	if pos.board != starting_board {
		t.Error("Board")
	}
	c := Castling{false, true, false, false}
	if pos.castling != c {
		t.Error("Castling")
	}
	if pos.enPassant != 9 {
		t.Error("En passant")
	}
	if pos.halfMoves != 23 {
		t.Error("Half moves")
	}
	if pos.fullMoves != 10 {
		t.Error("Full moves")
	}

}
