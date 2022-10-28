package model

import (
	"fmt"
	"testing"
)

// Tests creating team info
func Test_CreateTeam_Valid(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	testInfo := &Team{
		Name:         "test_teamname",
		TeamLocation: "test_teamlocation,",
		// Height and weight not included (unlike in users.go)
	}

	err = CreateTeam(testInfo)
	if err != nil {
		panic(err)
	}
	defer DBConn.Unscoped().Where("id = ?", testInfo.ID).Delete(testInfo)

	var testInfoCopy Team
	err = DBConn.Where("id = ?", testInfo.ID).Find(&testInfoCopy).Error

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// Set metadata equal
	testInfo.CreatedAt = testInfoCopy.CreatedAt
	testInfo.UpdatedAt = testInfoCopy.UpdatedAt
	testInfo.DeletedAt = testInfoCopy.DeletedAt

	if *testInfo != testInfoCopy {
		fmt.Println("hello")
		t.FailNow()
	}
}
