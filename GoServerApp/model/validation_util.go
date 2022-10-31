package model

// validation_util.go -> utilities for authenticating user permissions

// Authenticates a user based on username and password. Returns pointer
// to User if succeeded and error if could not find user.
func AuthenticateUser(username string, password string) (*User, error) {
	user := &User{}

	err := DBConn.Where("username = ? AND password = ?", username, password).
		First(user).Error

	return user, err
}
