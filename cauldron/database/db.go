package database

// InitAll initializes database
func InitAll() {
	initMySQL()
}

// StopAll database
func StopAll() {
	stopMySQL()
}
