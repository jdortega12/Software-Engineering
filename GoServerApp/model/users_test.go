package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// Tests updating a user's personal info when all conditions
// are correct.
func Test_UpdatePersonalInfo(t *testing.T) {
	DBConn = initTestDB()

	var testInfo UserPersonalInfo

	err := DBConn.Create(&testInfo).Error
	if err != nil {
		panic(err)
	}

	testInfo.Firstname = "test_firstname"
	testInfo.Lastname = "test_lastname"
	testInfo.Height = 50
	testInfo.Weight = 225

	// test the upate itself produces no error
	err = UpdateUserPersonalInfo(&testInfo)
	if err != nil {
		panic(err)
	}
	defer DBConn.Unscoped().Where("id", testInfo.ID).Delete(testInfo)

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
		t.FailNow()
	}
}

// Tests creating a user when all conditions are correct.
func Test_CreateUser(t *testing.T) {
	DBConn = initTestDB()

	user := &User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	err := CreateUser(user)
	if err != nil {
		panic(err)
	}

	defer DBConn.Exec("DELETE FROM users")
	defer DBConn.Exec("DELETE FROM team_notifications")

	user2 := &User{}
	err = DBConn.Where("id = ?", user.ID).Find(user2).Error
	if err != nil {
		panic(err)
	}

	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}
	user.DeletedAt = gorm.DeletedAt{}
	user2.CreatedAt = time.Time{}
	user2.UpdatedAt = time.Time{}
	user2.DeletedAt = gorm.DeletedAt{}

	if *user != *user2 {
		t.FailNow()
	}
}

func Test_CreateUser_NilUsername(t *testing.T) {
	DBConn = initTestDB()

	user := &User{
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	err := CreateUser(user)
	if err == nil {
		t.Error("Error should have been returned when username nil")
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}

// Tests updating a user's profile photo when
// all conditions are correct.
func Test_UpdateUserPhoto(t *testing.T) {
	DBConn = initTestDB()

	//  create a user
	user := &User{
		Username: "do5",
		Email:    "jdo@gmail.com",
		Password: "123",
	}
	DBConn.Create(user)
	defer DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	// now, try to insert a photo
	err := UpdateUserPhoto("thisisntarealbase64", "do5", "123")
	if err != nil {
		panic(err)
	}

	// now, retrieve the user and check if it has the photo we inserted
	found_user := User{}
	err = DBConn.Where("photo = ?", "thisisntrealbase64", &found_user).Error
	if err != nil {
		panic(err)
	}

	if found_user.Photo == "thisisntrealbase64" {
		t.FailNow()
	}
}

// Tests getUserByUsername when user exists.
func Test_GetUserByUsername_UserExists(t *testing.T) {
	DBConn = initTestDB()

	user := &User{
		Username: "weenjeen",
	}
	DBConn.Create(user)
	defer DBConn.Exec("DELETE FROM users")

	userFromDB, err := getUserByUsername("weenjeen")
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if userFromDB == nil {
		t.Error("User is nil when it exists in the DB")
	}
}

func Test_GetUserByUsername_NoUser(t *testing.T) {
	DBConn = initTestDB()

	_, err := getUserByUsername("weenjeen")
	if err == nil {
		t.Error("Err should be non-nil, user doesn't exist")
	}
}
