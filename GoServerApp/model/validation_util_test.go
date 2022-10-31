package model

import (
	"testing"
)

// Tests AuthenticateUser() when user exists.
func Test_AuthenticateUser_Exists(t *testing.T) {
	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = AuthenticateUser(user.Username, user.Password)
	if err != nil {
		t.Errorf("Error %s", err)
	}

	cleanUpDB()
}

// Tests that AuthenticateUser() returns error when credentials are bad.
func Test_AuthenticateUser_BadCredentials(t *testing.T) {
	user := &User{}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = AuthenticateUser("test_username", "test_password")
	if err == nil {
		t.Error("AuthenticateUser() should have returned non-nil error when user doesn't exist")
	}

	cleanUpDB()
}

// Tests that AuthenticateUser() returns err if provided username is an empty string.
func Test_AuthenticateUser_NilUsername(t *testing.T) {
	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = AuthenticateUser("", user.Password)
	if err == nil {
		t.Error("AuthenticateUser() returned no error when username was nil")
	}

	cleanUpDB()
}

// Tests that AuthenticateUser() returns error when password is empty string.
func Test_AuthenticateUser_NilPassword(t *testing.T) {
	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err := DBConn.Create(user).Error
	if err != nil {
		t.Error(err)
	}

	_, err = AuthenticateUser(user.Username, "")
	if err == nil {
		t.Error("AuthenticateUser() returned no error when password was nil")
	}

	cleanUpDB()
}
