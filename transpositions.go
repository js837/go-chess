package main

type Transposition struct {
	depth int
	eval  float64
}

const TABLE_SIZE int = 131072

// Note we build our own hash table to limit space.
type TranspositionTable [TABLE_SIZE]Transposition

func (*TranspositionTable) GetEval(position *Position, depth int) float64 {
	key := position.HashKey() % TABLE_SIZE
	value := TranspositionTable[key]

	if value.depth == 0 {
		return 0
	}

}
