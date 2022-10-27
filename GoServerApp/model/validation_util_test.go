package model

import "testing"

// Test validate user success case.
func TestValidateUser(t *testing.T) {
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

// Test ValidateUser() fail case.
func TestValidateUserFail(t *testing.T) {
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
	if err != nil {
		t.Errorf("Error %s", err)
	}
}
