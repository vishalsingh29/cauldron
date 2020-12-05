package users

import (
	"cauldron/database"
)

// User dao
type User struct {
	Name  string `db:"name"`
	Email string `db:"email"`
}

// Create new user
func Create(name string, email string) error {
	query := `insert into user(name, email) values (:name, :email)`
	values := map[string]interface{}{
		"name":  name,
		"email": email,
	}
	return database.InsertRow(query, values)
}

// Get user
func Get(email string) (User, error) {
	user := User{}
	query := "select name, email from user where email = ? LIMIT 1"
	err := database.ReadRow(&user, query, email)
	return user, err
}

// GetAllUsers gets all users
func GetAllUsers() ([]User, error) {
	objects := []User{}
	query := "select name, email from user"
	err := database.ReadMultipleRows(&objects, query)
	return objects, err
}
