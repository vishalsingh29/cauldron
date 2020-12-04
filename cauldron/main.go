package main

import (
	"cauldron/db"
	"fmt"
	"net/http"
)

func handlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "works")
}

func main() {
	fmt.Println("main init: listening on 8080")
	http.HandleFunc("/", handlerTest)
	http.ListenAndServe(":8080", nil)
	db.GetData()
}
