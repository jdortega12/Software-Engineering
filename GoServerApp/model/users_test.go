package model

import (
	"testing"
	"time"

	"gorm.io/gorm"
)

// Tests updating a user's personal info.
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

	testInfo = UserPersonalInfo{
		Firstname: "test_firstname",
		Lastname:  "test_lastname",

		Height: 50,
		Weight: 225,
	}

	// test the upate itself produces no error
	err = UpdateUserPersonalInfo(&testInfo)
	if err != nil {
		panic(err)
	}

	// pull record from DB to ensure it was saved correctly
	var testInfoCopy UserPersonalInfo
	err = DBConn.Where("user_personal_info_id = ?", testInfo.UserPersonalInfoID).
		Find(&testInfoCopy).Error

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// remove metadata before comparison because its supposed
	// to be different
	testInfo.CreatedAt = time.Time{}
	testInfo.UpdatedAt = time.Time{}
	testInfo.DeletedAt = gorm.DeletedAt{}
	testInfoCopy.CreatedAt = time.Time{}
	testInfoCopy.UpdatedAt = time.Time{}
	testInfoCopy.DeletedAt = gorm.DeletedAt{}

	if testInfo != testInfoCopy {
		t.FailNow()
	}
}

func TestCreateUser(t *testing.T) {
	//Initialize Database
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

	err = CreateUser(user)

	if err != nil {
		panic(err)
	}

	user2 := &User{}
	err = DBConn.Where("user_id = ?", user.UserID).Find(user2).Error
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

// test if we can create a user, and then insert a picture without errors
func TestInsertPhoto(t *testing.T) {
	//Initialize Database
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}
	
	// Insert a user
	user := &User{
		Username: "do5",
		Email:    "jdo@gmail.com",
		Password: "123",
	}
	
	err = CreateUser(user)
	
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
