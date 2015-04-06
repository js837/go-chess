package main

/*
Transpotion Table

This is used to store results of evalutions of moves.
*/

type TranspositionKey struct {
	turn      Colour
	board     Board
	castling  Castling // Castling rights
	enPassant int      // En passant square
}

type Transposition struct {
	depth int
	eval  int
}

// Note we should build our own hash table to limit space.
type TranspositionTable map[TranspositionKey]Transposition

func (table TranspositionTable) LookupPosition(position *Position, colour Colour) (bool, int, int) {
	var key TranspositionKey = TranspositionKey{
		colour,
		position.board,
		position.castling,
		position.enPassant,
	}

	trans, found := table[key]
	if found {
		return true, trans.depth, trans.eval
	} else {
		return false, 0, 0
	}

}

func (table TranspositionTable) SetPosition(colour Colour, position *Position, depth int, eval int) {
	var key TranspositionKey = TranspositionKey{
		colour,
		position.board,
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
