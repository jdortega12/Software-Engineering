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
