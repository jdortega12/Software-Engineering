package model

import (
	"testing"
)

// Tests CreateTeam() creates the team in the DB correctly when
// the struct passed is correct.
func Test_CreateTeam_Valid(t *testing.T) {
	testInfo := &Team{
		Name:         "test_teamname",
		TeamLocation: "test_teamlocation,",
		// Height and weight not included (unlike in users.go)
	}

	err := CreateTeam(testInfo)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

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

	cleanUpDB()
}

// Tests that DeleteTeam() deletes team correctly.
func Test_DeleteTeam(t *testing.T) {
	team := &Team{
		Name:         "Badgers",
		TeamLocation: "Badgertown",
	}

	err := DBConn.Create(team).Error
	if err != nil {
		t.Error(err)
	}

	teamID := team.ID

	err = DeleteTeam(team)
	if err != nil {
		t.Error(err)
	}

	teamAfter := &Team{}
	err = DBConn.Where("id = ?", teamID).First(teamAfter).Error
	if err == nil {
		t.Error("Team should not have been findable in DB after deletion")
	}

	cleanUpDB()
}

// Tests that DeleteTeam() does not return an error even if team doesn't exist.
func Test_DeleteTeam_NotInDB(t *testing.T) {
	team := &Team{
		ID:           5,
		Name:         "Badgers",
		TeamLocation: "Badgertown",
	}

	err := DeleteTeam(team)
	if err != nil {
		t.Error(err)
	}
}

// Tests that GetTeamByID() returns correct team when it exists.
func Test_GetTeamByID_Valid(t *testing.T) {
	team := &Team{
		ID:   5,
		Name: "Badgers",
	}
	DBConn.Create(team)

	teamFromDB, err := GetTeamByID(5)
	if err != nil {
		t.Error(err)
	}

	team.CreatedAt = teamFromDB.CreatedAt
	team.UpdatedAt = teamFromDB.UpdatedAt
	team.DeletedAt = teamFromDB.DeletedAt

	if *teamFromDB != *team {
		t.Error("team taken from db was not the same as what was created")
	}

	cleanUpDB()
}

// Makes sure GetTeamByID() returns an error when team doesn't exist.
func Test_GetTeamByID_Invalid(t *testing.T) {
	_, err := GetTeamByID(5)
	if err == nil {
		t.Error("should have produced an error when team doesn't exist")
	}
}
