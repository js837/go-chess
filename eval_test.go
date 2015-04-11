package main

import (
	"fmt"
	"testing"
)

func TestAlphaBeta(t *testing.T) {

	var cutoffs int = 0

	pos := FromFen("r4br1/3b1kpp/1q1P4/1pp1RP1N/p7/6Q1/PPB3PP/2KR4 w - - 0 1")
	pos1 := FromFen("r2q1k1r/ppp1bB1p/2np4/6N1/3PP1bP/8/PPP5/RNB2RK1 w - - 0 1")

	if minimax(&pos, 3, White) != alphabeta(&pos, 3, White, -EVAL_MAX, +EVAL_MAX, &cutoffs) {
		t.Error("Minimax and alphabeta don't agree")
	}

	if minimax(&pos1, 3, White) != alphabeta(&pos1, 3, White, -EVAL_MAX, +EVAL_MAX, &cutoffs) {
		t.Error("Minimax and alphabeta don't agree")
	}

}

func TestCutoffs(t *testing.T) {

	var cutoffs int = 0

	pos := FromFen("rnb1kb1r/p1qpp2p/8/1p2Ppp1/1PB5/P2p2P1/Q4P1P/R1B1K1nR w KQkq - 0 0")

	alphabeta(&pos, 8, White, -EVAL_MAX, +EVAL_MAX, &cutoffs)
	fmt.Println("Cut offs:", cutoffs)

}
