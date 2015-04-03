package main

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func min64(a, b float64) float64 {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

func max64(a, b float64) float64 {
	if a <= b {
		return b
	} else {
		return a
	}
}

func getPiece(char rune) Piece {

	switch char {
	case 'P':
		return WP
	case 'N':
		return WN
	case 'B':
		return WB
	case 'R':
		return WR
	case 'Q':
		return WQ
	case 'K':
		return WK

	case 'p':
		return BP
	case 'n':
		return BN
	case 'b':
		return BB
	case 'r':
		return BR
	case 'q':
		return BQ
	case 'k':
		return BK

	default:
		return NoPiece
	}
}

func getChar(piece Piece) rune {

	switch piece {
	case WP:
		return 'P'
	case WN:
		return 'N'
	case WB:
		return 'B'
	case WR:
		return 'R'
	case WQ:
		return 'Q'
	case WK:
		return 'K'

	case BP:
		return 'p'
	case BN:
		return 'n'
	case BB:
		return 'b'
	case BR:
		return 'r'
	case BQ:
		return 'q'
	case BK:
		return 'k'
	default:
		return '-'
	}

}
