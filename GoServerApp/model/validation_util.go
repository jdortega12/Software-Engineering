package model

// validation_util.go -> utilities for validating user permissions

// Validates a user based on username and password. Returns role and ID
// of user if validation succeeded and error if it did not succeed.
func ValidateUser(username string, password string) (uint, userRole, error) {
	user := &User{}

	err := DBConn.Where("username = ? AND password = ?", username, password).
		First(user).Error

	return user.ID, user.Role, err
}
