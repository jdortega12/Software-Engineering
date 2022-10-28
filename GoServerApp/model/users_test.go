package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// Tests updating a user's personal info when all conditions
// are correct.
func TestUpdatePersonalInfo(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	var testInfo UserPersonalInfo

	err = DBConn.Create(&testInfo).Error
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
func TestCreateUser(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	user := &User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	DBConn.Create(user)
	defer DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	if err != nil {
		panic(err)
	}

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

// Tests updating a user's profile photo when
// all conditions are correct.
func TestUpdateUserPhoto(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	//  create a user
	user := &User{
		Username: "do5",
		Email:    "jdo@gmail.com",
		Password: "123",
	}
	DBConn.Create(user)
	defer DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	if err != nil {
		panic(err)
	}

	// now, try to insert a photo
	err = UpdateUserPhoto("thisisntarealbase64", "do5", "123")
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
