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
