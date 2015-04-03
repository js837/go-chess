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

	p := FromFen(string(fen))

	moves := p.GetMoves(p.turn)

	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(moves))

	newPosition := p.ApplyMove(moves[i])

	fmt.Fprintf(w, PositionToBoardFen(&newPosition))

}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
