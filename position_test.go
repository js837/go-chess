package main

import (
	"testing"
)

func TestGetRankFile(t *testing.T) {

	a, b := GetRankFile(63)

	if a != 7 || b != 7 {
		t.Error("GetRankFile did not work as expected.")
	}

	c, d := GetRankFile(3)

	if c != 0 || d != 3 {
		t.Error("GetRankFile did not work as expected.")
	}
}

func TestSwitchColour(t *testing.T) {

	w := Colour(White)
	b := Colour(Black)

	if w.Switch() != b || b.Switch() != w {
		t.Error("OtherColour did not work as expected.")
	}
}
