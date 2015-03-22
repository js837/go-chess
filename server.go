package main

import (
	"fmt"
	"net/http"
)

func EndPointHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomeHandler")
}

func main() {

	http.HandleFunc("/endpoint", EndPointHandler) // homepage
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8000", nil)
}
