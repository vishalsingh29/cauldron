package users

// User dao
type User struct {
	Name  string
	Email string
}

// Create new user
func Create(name string, email string) (User, error) {
	return User{}, nil
}
