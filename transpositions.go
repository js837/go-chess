package main

/*
Transpotion Table

This is used to store results of evalutions of moves.
*/

type TranspositionKey struct {
	board     Board
	turn      Colour
	castling  Castling // Castling rights
	enPassant int      // En passant square
}

type Transposition struct {
	depth int
	eval  float64
}

const TABLE_SIZE int = 131072

// Note we should build our own hash table to limit space.
type TranspositionTable map[TranspositionKey]Transposition

func (table TranspositionTable) LookupPosition(position *Position) (int, float64) {
	var key TranspositionKey = TranspositionKey{
		position.board,
		position.turn,
		position.castling,
		position.enPassant,
	}

	trans, found := table[key]
	if found {
		return trans.depth, trans.eval
	} else {
		return -1, 0
	}

}

func (table TranspositionTable) SetPosition(position *Position, depth int, eval float64) {
	var key TranspositionKey = TranspositionKey{
		position.board,
		position.turn,
		position.castling,
		position.enPassant,
	}
	trans, found := table[key]
	if found {
		// Is the stored value higher depth and more useful?
		if trans.depth >= depth {
			return
		}
	}
	// Add the position to the Table
	table[key] = Transposition{depth, eval}

}
