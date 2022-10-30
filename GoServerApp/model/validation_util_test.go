package model

import (
	"testing"
)

// Tests ValidateUser() when user exists.
func Test_ValidateUser_Exists(t *testing.T) {
	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = ValidateUser(user.Username, user.Password)
	if err != nil {
		t.Errorf("Error %s", err)
	}

	cleanUpDB()
}

// Tests that ValidateUser() returns error when credentials are bad.
func Test_ValidateUser_BadCredentials(t *testing.T) {
	user := &User{}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = ValidateUser("test_username", "test_password")
	if err == nil {
		t.Error("validateUser() should have returned non-nil error when user doesn't exist")
	}

	cleanUpDB()
}

// Tests that ValidateUser() returns err if provided username is an empty string.
func Test_ValidateUser_NilUsername(t *testing.T) {
	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = ValidateUser("", user.Password)
	if err == nil {
		t.Error("validateUser() returned no error when username was nil")
	}

	cleanUpDB()
}

// Tests that ValidateUser() returns error when password is empty string.
func Test_ValidateUser_NilPassword(t *testing.T) {
	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = ValidateUser(user.Username, "")
	if err == nil {
		t.Error("validateUser() returned no error when password was nil")
	}

	cleanUpDB()
}
