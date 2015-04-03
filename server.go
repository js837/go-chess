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

	p := FromFen(string(fen))

	fmt.Println(p)
	fmt.Println(string(fen))
	fmt.Println(p.turn == White)
	fmt.Println(p.turn == Black)

	fmt.Fprintf(w, "Reply")

}

func main() {

	http.HandleFunc("/setfen", SetFen)

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8000", nil)
}
