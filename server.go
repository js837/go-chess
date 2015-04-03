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

	fmt.Println("Best")

	fmt.Fprintf(w, PositionToBoardFen(&p))

}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8005", nil)
}
