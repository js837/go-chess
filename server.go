package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func SetFen(w http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	fen, _ := ioutil.ReadAll(request.Body)

	p := PositionFromBoardFen(string(fen))
	for _, move := range p.GetMoves(p.turn) {
		fmt.Println(move)
	}

	moves := p.GetMoves(p.turn)
	if len(moves) > 0 {
		rand.Seed(time.Now().Unix())
		firstMove := moves[rand.Intn(len(moves))]
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
