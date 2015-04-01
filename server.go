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

	//var moves []Move

	output := p.PieceEval()

	fmt.Fprintf(w, fmt.Sprintf("x = %.6f\n", output))
}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
