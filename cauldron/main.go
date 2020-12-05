package main

import (
	"cauldron/config"
	"cauldron/database"
	"cauldron/users"
	"fmt"
	"net/http"
)

func handlerTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "works")
	user, err := users.Create("vishal", "vishal.singh@lucifer.com")
	if err != nil {
		fmt.Println(fmt.Errorf("Error creating a user %v", err))
		return
	}
	fmt.Println(user.Name, user.Email, " created")
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
