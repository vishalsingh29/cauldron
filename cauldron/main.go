package main

import (
	"cauldron/config"
	"cauldron/database"
	"fmt"
	"net/http"
)

func handlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "works")
}

func main() {
	fmt.Println("main init: listening on 8080")

	// config init
	config.Init()

	// init database
	database.InitAll(&config.Config.MySQL)

	http.HandleFunc("/", handlerTest)
	http.ListenAndServe(":8080", nil)

	// stop database
	defer database.StopAll()
}
