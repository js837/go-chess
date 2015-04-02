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

	//output := p.PieceEval(p.turn)

	root := EvalNode{p, 0, []EvalNode{}}
	root.GenerateTree(2)

	//	for i, el := range root.children {
	//		fmt.Printf("%d %.6f", i, el.eval)
	//	}

	best := minimax(&p, 3, White)
	var bestFen string

	for _, move := range p.GetMoves(p.turn) {
		newPosition := p.ApplyMove(move)
		if newPosition.PieceEval(p.turn) == best {
			bestFen = PositionToBoardFen(&newPosition)
			break
		}
	}

	fmt.Println("Best:", best)

	fmt.Fprintf(w, bestFen)

}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8005", nil)
}
