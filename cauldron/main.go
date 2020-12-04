package main

import (
	"fmt"
	"net/http"
)

func handlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "works")
}

func main() {
	fmt.Println("main init: listening on 8080")
	fmt.Println("does it work?")
	http.HandleFunc("/", handlerTest)
	http.ListenAndServe(":8080", nil)
}
