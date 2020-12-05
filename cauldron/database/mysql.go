package database

import (
	"cauldron/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	// database sql package is used by mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const mysql = "mysql"

var mysqlDB *sqlx.DB

// Init initializes data
func initMySQL(conf *config.MySQLConfig) {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Username, conf.Password, conf.Hostname, conf.Port, conf.DBName)
	var err error
	mysqlDB, err = sqlx.Connect(mysql, mysqlURI)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	mysqlDB.SetConnMaxLifetime(time.Minute * 3)
	mysqlDB.SetMaxOpenConns(10)
	mysqlDB.SetMaxIdleConns(10)
}

// ReadRow executes a raw query
func ReadRow(dest interface{}, query string, args ...interface{}) error {
	return mysqlDB.Get(dest, query, args...)
}

// ReadMultipleRows reads multiple rows in a query
func ReadMultipleRows(dest interface{}, query string, args ...interface{}) error {
	return mysqlDB.Select(dest, query, args...)
}

// InsertRow inserts a single row
func InsertRow(query string, values map[string]interface{}) error {
	_, err := mysqlDB.NamedExec(query, values)
	return err
}

// Stop database
func stopMySQL() {
	mysqlDB.Close()
}
