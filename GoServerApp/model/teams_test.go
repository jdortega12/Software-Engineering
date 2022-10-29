package model

import (
	"testing"
)

// Tests CreateTeam() creates the team in the DB correctly when
// the struct passed is correct.
func Test_CreateTeam_Valid(t *testing.T) {
	initTestDB()

	testInfo := &Team{
		Name:         "test_teamname",
		TeamLocation: "test_teamlocation,",
		// Height and weight not included (unlike in users.go)
	}

	err := CreateTeam(testInfo)
	if err != nil {
		t.Errorf("Error: %s", err)
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
		t.Error("Structs should be equal")
	}
}
