package main

import (
	"fmt"
	"io/ioutil"
	//"math/rand"
	"net/http"
	//"time"
)

func SetFen(w http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	fen, _ := ioutil.ReadAll(request.Body)

	p := PositionFromBoardFen(string(fen))
	//	for _, move := range p.GetMoves(p.turn) {
	//		fmt.Println(move)
	//	}

	var moves []Move

	for i := 0; i < 1000000; i++ {
		moves = p.GetMoves(p.turn)
		for _, move := range moves {
			p.ApplyMove(move)
		}
	}

	fmt.Fprintf(w, string(fen))
}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
