package database

import (
	"database/sql"
	"time"

	// database sql package is used by mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB

// Init initializes data
func initMySQL() {
	mysqlDB, err := sql.Open("mysql", "root:zhomsg3d@tcp(127.0.0.1:3306)/cauldron")
	if err != nil {
		panic(err)
	} // See "Important settings" section.
	mysqlDB.SetConnMaxLifetime(time.Minute * 3)
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetMaxIdleConns(10)
}

// Stop database
func stopMySQL() {
	mysqlDB.Close()
}
