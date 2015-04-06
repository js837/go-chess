package main

import (
	"fmt"
	"testing"
)

func TestAlphaBeta(t *testing.T) {

	pos := FromFen("r4br1/3b1kpp/1q1P4/1pp1RP1N/p7/6Q1/PPB3PP/2KR4 w - - 0 1")
	pos1 := FromFen("r2q1k1r/ppp1bB1p/2np4/6N1/3PP1bP/8/PPP5/RNB2RK1 w - - 0 1")

	fmt.Println(alphabeta(&pos, 5, -EVAL_MAX, +EVAL_MAX, White))

	if minimax(&pos, 5, White) != alphabeta(&pos, 5, -EVAL_MAX, +EVAL_MAX, White) {
		t.Error("Minimax and alphabeta don't agree")
	}

	if minimax(&pos1, 3, White) != alphabeta(&pos1, 3, -EVAL_MAX, +EVAL_MAX, White) {
		t.Error("Minimax and alphabeta don't agree")
	}

}
