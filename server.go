package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SetFen(w http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	fen, _ := ioutil.ReadAll(request.Body)

	p := PositionFromBoardFen(string(fen))
	for _, move := range p.GetMoves() {
		fmt.Println(move)
	}

	moves := p.GetMoves()
	if len(moves) > 0 {
		firstMove := moves[0]
		newP := p.ApplyMove(firstMove)

		fen = []byte(PositionToBoardFen(&newP))

	}

	fmt.Fprintf(w, string(fen))
}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
