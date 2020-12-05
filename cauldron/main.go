package main

import (
	"cauldron/database"
	"fmt"
	"net/http"
)

func handlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "works")
}

func main() {
	fmt.Println("main init: listening on 8080")

	// init database
	database.InitAll()
	fmt.Println("does it work?")
	http.HandleFunc("/", handlerTest)
	http.ListenAndServe(":8080", nil)
	// stop database
	defer database.StopAll()
}
