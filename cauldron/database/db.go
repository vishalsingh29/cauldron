package database

import "cauldron/config"

// InitAll initializes database
func InitAll(mysqlConf *config.MySQLConfig) {
	initMySQL(mysqlConf)
}

// StopAll database
func StopAll() {
	stopMySQL()
}
