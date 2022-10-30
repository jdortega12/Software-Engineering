package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// Tests UpdateUserPersonalInfo() when all conditions are correct.
func Test_UpdateUserPersonalInfo(t *testing.T) {
	var testInfo UserPersonalInfo

	err := DBConn.Create(&testInfo).Error
	if err != nil {
		t.Error(err)
	}

	testInfo.Firstname = "test_firstname"
	testInfo.Lastname = "test_lastname"
	testInfo.Height = 50
	testInfo.Weight = 225

	// test the upate itself produces no error
	err = UpdateUserPersonalInfo(&testInfo)
	if err != nil {
		t.Error(err)
	}

	// pull record from DB to ensure it was saved correctly
	var testInfoCopy UserPersonalInfo
	err = DBConn.Where("id = ?", testInfo.ID).
		Find(&testInfoCopy).Error

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// remove metadata before comparison
	testInfo.CreatedAt = time.Time{}
	testInfo.UpdatedAt = time.Time{}
	testInfo.DeletedAt = gorm.DeletedAt{}
	testInfoCopy.CreatedAt = time.Time{}
	testInfoCopy.UpdatedAt = time.Time{}
	testInfoCopy.DeletedAt = gorm.DeletedAt{}

	if testInfo != testInfoCopy {
		t.Error("The updated info was not saved correctly in the DB")
	}

	cleanUpDB()
}

// Tests CreateUser() when all conditions are correct.
func Test_CreateUser(t *testing.T) {
	user := &User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	err := CreateUser(user)
	if err != nil {
		t.Error(err)
	}

	user2 := &User{}
	err = DBConn.Where("id = ?", user.ID).Find(user2).Error
	if err != nil {
		t.Error(err)
	}

	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}
	user.DeletedAt = gorm.DeletedAt{}
	user2.CreatedAt = time.Time{}
	user2.UpdatedAt = time.Time{}
	user2.DeletedAt = gorm.DeletedAt{}

	if *user != *user2 {
		t.Error()
	}

	cleanUpDB()
}

// Tests CreateUser() when username was not supplied.
func Test_CreateUser_NilUsername(t *testing.T) {
	user := &User{
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	err := CreateUser(user)
	if err == nil {
		t.Error("Error should have been returned when username nil")
	}

	cleanUpDB()
}

// Tests UpdateUserPhoto() when all conditions are correct.
func Test_UpdateUserPhoto(t *testing.T) {
	//  create a user
	user := &User{
		Username: "do5",
		Email:    "jdo@gmail.com",
		Password: "123",
	}
	DBConn.Create(user)

	// now, try to insert a photo
	err := UpdateUserPhoto("thisisntarealbase64", "do5", "123")
	if err != nil {
		t.Error(err)
	}

	// now, retrieve the user and check if it has the photo we inserted
	found_user := User{}
	err = DBConn.Where("photo = ?", "thisisntrealbase64", &found_user).Error
	if err != nil {
		t.Error(err)
	}

	if found_user.Photo == "thisisntrealbase64" {
		t.Fail()
	}

	cleanUpDB()
}

// Tests getUserByUsername() when user exists.
func Test_GetUserByUsername_UserExists(t *testing.T) {
	user := &User{
		Username: "weenjeen",
	}
	DBConn.Create(user)

	userFromDB, err := getUserByUsername("weenjeen")
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if userFromDB == nil {
		t.Error("User is nil when it exists in the DB")
	}

	cleanUpDB()
}

// Tests that getUserByUsername() returns error when user doesn't exist.
func Test_GetUserByUsername_NoUser(t *testing.T) {
	_, err := getUserByUsername("weenjeen")
	if err == nil {
		t.Error("Err should be non-nil, user doesn't exist")
	}
}
