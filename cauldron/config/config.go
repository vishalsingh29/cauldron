package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// MySQLConfig variables
type MySQLConfig struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	// TODO: put password in vault
	Password string `json:"password"`
	Username string `json:"username"`
	DBName   string `json:"dbname"`
}

// Config for the app
var Config struct {
	MySQL MySQLConfig `json:"mysql"`
}

// Init config
func Init() {
	// TODO: pick AppMode = local from env
	configFile, err := os.Open("./config/local.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&Config); err != nil {
		fmt.Println("parsing config file", err.Error())
	}
}
