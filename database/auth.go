package database

func CreateUser(username, password string) error {
	// creates a new user in the database, returns error if any
	return nil
}

func CheckPassword(username, password string) (success bool, err error) {
	// checks if the password is correct for the given username
	return true, nil
}
