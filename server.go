package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SetFen(w http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	body, _ := ioutil.ReadAll(request.Body)

	p := PositionFromBoardFen(string(body))
	for _, move := range p.GetMoves() {
		fmt.Println(move)
	}

	firstMove := p.GetMoves()[0]
	newP := p.ApplyMove(firstMove)

	fen := PositionToBoardFen(&newP)

	fmt.Fprintf(w, string(fen))
}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
