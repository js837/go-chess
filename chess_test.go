package main

import (
	"testing"
)

func TestGetRankFile(t *testing.T) {

	a, b := GetRankFile(63)

	if a != 7 || b != 7 {
		t.Error("Add2Ints did not work as expected.")
	} else {
		t.Log("one test passed.") // log some info if you want
	}
}
