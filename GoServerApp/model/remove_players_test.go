package model

import "testing"

// Tests that ManagerRemovePlayer() returns no errors
func Test_ManagerRemovePlayer_ValidRemoval(t *testing.T) {

	DBConn.Create(&User{
		Username: "manny",
		Role:     MANAGER,
		TeamID:   17,
	})
	DBConn.Create(&User{
		Username: "brady",
		Role:     PLAYER,
		TeamID:   17,
	})

	removePlayer := &RemovePlayer{
		PlayerUsername:  "brady",
		ManagerUsername: "manny",
	}

	err := ManagerRemovePlayer(removePlayer)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	cleanUpDB()
}

// Tests that ManagerRemovePlayer() returns error when Manager has incorrect role
func Test_ManagerRemovePlayer_InvalidUser(t *testing.T) {
	DBConn.Create(&User{
		Username: "manny",
		Role:     MANAGER,
		TeamID:   17,
	})
	DBConn.Create(&User{
		Username: "brady",
		Role:     ADMIN,
		TeamID:   17,
	})

	removePlayer := &RemovePlayer{
		PlayerUsername:  "brady",
		ManagerUsername: "manny",
	}
	err := ManagerRemovePlayer(removePlayer)
	if err == nil {
		t.Error("Error should be produced since user to delete is ADMIN")
	}
	cleanUpDB()
}

// Tests that ManagerRemovePlayer returns error when Manager has incorrect role
func Test_ManagerRemovePlayer_InvalidManager(t *testing.T) {
	DBConn.Create(&User{
		Username: "manny",
		Role:     PLAYER,
		TeamID:   17,
	})
	DBConn.Create(&User{
		Username: "brady",
		Role:     PLAYER,
		TeamID:   17,
	})

	removePlayer := &RemovePlayer{
		PlayerUsername:  "brady",
		ManagerUsername: "manny",
	}
	err := ManagerRemovePlayer(removePlayer)
	if err == nil {
		t.Error("Error should be produced since user removing is PLAYER not MANAGER")
	}
	cleanUpDB()
}

// Tests that ManagerRemovePlayer returns error when Manager and PLAYER are not on the same team
func Test_ManagerRemovePlayer_InvalidTeam(t *testing.T) {
	DBConn.Create(&User{
		Username: "manny",
		Role:     MANAGER,
		TeamID:   18,
	})
	DBConn.Create(&User{
		Username: "brady",
		Role:     PLAYER,
		TeamID:   17,
	})

	removePlayer := &RemovePlayer{
		PlayerUsername:  "brady",
		ManagerUsername: "manny",
	}
	err := ManagerRemovePlayer(removePlayer)
	if err == nil {
		t.Error("Error should be produced since PLAYER and MANAGER are on different teams")
	}
	cleanUpDB()
}
