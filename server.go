package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func BestMove(w http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	fen, _ := ioutil.ReadAll(request.Body)
	p := FromFen(string(fen))

	bestMove, _ := p.GetBestMove(4, p.turn)

	fmt.Println(bestMove)
	fmt.Println(string(fen))

	fmt.Fprintf(w, PositionToBoardFen(&p))
}

func RandomMove(w http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()
	fen, _ := ioutil.ReadAll(request.Body)
	p := FromFen(string(fen))

	moves := p.GetMoves(p.turn)

	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(moves))
	newPosition := p.ApplyMove(moves[i])

	fmt.Fprintf(w, PositionToBoardFen(&newPosition))
}

func main() {

	http.HandleFunc("/random", RandomMove)
	http.HandleFunc("/best", BestMove)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":9000", nil)
}
