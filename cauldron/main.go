package main

import (
	"cauldron/config"
	"cauldron/database"
	"cauldron/users"
	"fmt"
	"net/http"
)

func doNothing(w http.ResponseWriter, r *http.Request) {}

func handlerTest(w http.ResponseWriter, r *http.Request) {
	allUsers, err := users.GetAllUsers()
	fmt.Println("len users ", len(allUsers))
	if err != nil {
		fmt.Println(fmt.Errorf("Error getting a user %v", err))
		return
	}
	err = users.Create("vishalsingh", "vishalrsingh@google.com")
	if err != nil {
		fmt.Println("err = ", err)
	}
	output := ""
	for _, v := range allUsers {
		output += fmt.Sprintf("%s, %s\n", v.Name, v.Email)
		fmt.Println(v.Name, v.Email, "here ")
	}
	fmt.Fprintf(w, output)
}

func main() {
	fmt.Println("main init: listening on 8080")

	// config init
	config.Init()

	// init database
	database.InitAll(&config.Config.MySQL)
	http.HandleFunc("/favicon.ico", doNothing)
	http.HandleFunc("/", handlerTest)
	http.ListenAndServe(":8080", nil)

	// stop database
	// defer database.StopAll()
}
