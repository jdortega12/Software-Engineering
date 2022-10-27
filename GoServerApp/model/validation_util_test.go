package model

import (
	"log"
	"testing"
)

// Test validate user success case.
func TestGoodValidateUser(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err = DBConn.Create(user).Error
	if err != nil {
		panic(err)
	}
	defer DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	_, _, err = ValidateUser(user.Username, user.Password)
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

// Test that error returned is
func TestBadValidateUser(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	user := &User{}

	err = DBConn.Create(user).Error
	if err != nil {
		panic(err)
	}
	defer DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	_, _, err = ValidateUser("test_username", "test_password")
	if err == nil {
		log.Println("validateUser() should have returned non-nil error when user doesn't exist")
		t.FailNow()
	}
}

// Tests that err is returned if provided username
// is an empty string.
func TestValidateUserNilUsername(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err = DBConn.Create(user).Error
	if err != nil {
		panic(err)
	}
	defer DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	_, _, err = ValidateUser("", user.Password)
	if err == nil {
		log.Println("validateUser() returned no error when username was nil")
		t.FailNow()
	}
}

// Tests that err is returned if provided password
// is an empty string.
func TestValidateUserNilPassword(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	user := &User{
		Username: "test_username",
		Password: "test_password",
		Role:     PLAYER,
	}

	err = DBConn.Create(user).Error
	if err != nil {
		panic(err)
	}
	defer DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	_, _, err = ValidateUser(user.Username, "")
	if err == nil {
		log.Println("validateUser() returned no error when password was nil")
		t.FailNow()
	}
}
