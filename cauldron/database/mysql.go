package database

import (
	"cauldron/config"
	"database/sql"
	"fmt"
	"time"

	// database sql package is used by mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const mysql = "mysql"

var mysqlDB *sql.DB

// Init initializes data
func initMySQL(conf *config.MySQLConfig) {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s)/%s", conf.Username, conf.Password, conf.Hostname, conf.DBName)
	mysqlDB, err := sql.Open(mysql, mysqlURI)
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
