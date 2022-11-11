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

// Tests GetUserByUsername() when user exists.
func Test_GetUserByUsername_UserExists(t *testing.T) {
	user := &User{
		Username: "weenjeen",
	}
	DBConn.Create(user)

	userFromDB, err := GetUserByUsername("weenjeen")
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if userFromDB == nil {
		t.Error("User is nil when it exists in the DB")
	}

	cleanUpDB()
}

// Tests that GetUserByUsername() returns error when user doesn't exist.
func Test_GetUserByUsername_NoUser(t *testing.T) {
	_, err := GetUserByUsername("weenjeen")
	if err == nil {
		t.Error("Err should be non-nil, user doesn't exist")
	}
}

// Tests that UpdateUserTeam() correctly updates user in DB.
func Test_UpdateUserTeam(t *testing.T) {
	user := &User{
		Username: "jaluhrman",
		Password: "peepee",
		Role:     PLAYER,
		TeamID:   0,
	}

	DBConn.Create(user)

	UpdateUserTeam(user, 5)

	userUpdated := &User{}
	DBConn.Where("id = ?", user.ID).First(userUpdated)

	if userUpdated.TeamID != 5 {
		t.Error("user's TeamID was not updated correctly")
	}

	cleanUpDB()
}

// Tests that GetUserPersonalInfoByID() works when info exists.
func Test_GetUserPersonalInfoByID_Valid(t *testing.T) {
	testUserID := uint(1)

	testInfo := &UserPersonalInfo{
		ID:        testUserID,
		Firstname: "Joe",
		Lastname:  "Luhrman",
		Height:    50,
		Weight:    50,
	}
	if err := DBConn.Create(testInfo).Error; err != nil {
		t.Error(err)
	}

	testInfoCpy, err := GetUserPersonalInfoByID(testUserID)
	if err != nil {
		t.Error(err)
	}

	testInfoCpy.CreatedAt = testInfo.CreatedAt
	testInfoCpy.UpdatedAt = testInfo.UpdatedAt
	testInfoCpy.DeletedAt = testInfo.DeletedAt

	if *testInfo != *testInfoCpy {
		t.Error("structs should be the same before creating and after reading from DB")
	}

	cleanUpDB()
}

// Makes sure that GetUserPersonalInfoByID() returns error when info
// does not exist in the DB.
func Test_GetUserPersonalInfoByID_NotExists(t *testing.T) {
	testUserID := uint(1)

	_, err := GetUserPersonalInfoByID(testUserID)
	if err == nil {
		t.Error("should have produced an error when user info not in DB")
	}

}


// Insert a manager into the DB with a TeamId of 1 and check that the record can be recovered
func TestGetManagerByTeam(t *testing.T) {
	// insert the manager
	user := User {
		TeamID: 1,
		Username: "Joe Douglas",
		Password: "allgasnobreaks",
		Role: MANAGER,
	}

	DBConn.Create(&user)

	// Now, look for the manager by calling GetManagerByTeamId
	manager, err := GetManagerByTeamID(1)

	if err != nil {
		t.Error(err)
	}

	if manager.Username != "Joe Douglas" {
		t.Error("Did not retrieve record")
	}

	cleanUpDB()
}

// Insert a player into the DB and make sure it can't retrieve a record
func TestGetManagerBad(t * testing.T) {
	// insert the player 
	user := User {
		TeamID: 1,
		Username: "Sauce Gardner",
		Password: "allgasnobreaks",
	}

	DBConn.Create(&user)

	_, err := GetManagerByTeamID(1)

	if err == nil {
		t.Error("This is supposed to fail")
	}

	cleanUpDB()

}

// Insert Players belonging to a team into the DB and make sure they can be recovered
func TestGetPlayersByTeam(t *testing.T) {
	user := User {
		TeamID: 1, 
		Username: "saucegardner",
		Password: "allgasnobreaks",
	}

	user2 := User{
		TeamID: 1,
		Username: "zachwilson",
		Password: "imbad@football",
	}

	DBConn.Create(&user)
	DBConn.Create(&user2)

	players, err := GetPlayersByTeamID(1)

	if err != nil {
		t.Error("Error retrieving players")
	}

	if players[0].Username != "saucegardner" || players[1].Username != "zachwilson" {
		t.Error("I didn't retrieve players correctly")
	}

	cleanUpDB()
}

// Make sure the DB does not retrieve players belonging to the wrong team
func TestGetPlayersWrongTeam(t *testing.T) {
	user := User{
		TeamID: 2,
		Username: "imnotintherightteam",
		Password: "imnotinteam1",
	}

	DBConn.Create(&user)

	players, err := GetPlayersByTeamID(1)

	if err != nil {
		t.Error(err)
	}

	if len(players) != 0 {
		t.Error("I'm not supposed to be here")
	}

	cleanUpDB()
}