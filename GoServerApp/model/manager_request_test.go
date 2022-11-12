package model

import "testing"

// Tests that CreateManagerRequest() returns no errors
// in the case of a valid request from a player to an admin.
func Test_CreateManagerRequest_ValidRequest(t *testing.T) {
	initTestDB()

	DBConn.Create(&User{
		Username: "MeytonPanning",
		Role:     PLAYER,
	})
	DBConn.Create(&User{
		Username: "WillBelichick",
		Role:     ADMIN,
	})

	newRequest := &ManagerRequest{
		PlayerUsername: "MeytonPanning",
		AdminUsername:  "WillBelichick",
	}

	err := CreateManagerRequest(newRequest)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	DBConn.Exec("DELETE FROM users")
	// DBConn.Exec("DELETE FROM manager_requests")
}

// Test for producing error when sender is not a player
func Test_CreateManagerRequest_InvalidSender(t *testing.T) {
	initTestDB()

	DBConn.Create(&User{
		Username: "MeytonPanning",
		Role:     MANAGER,
	})
	DBConn.Create(&User{
		Username: "WillBelichick",
		Role:     ADMIN,
	})

	newRequest := &ManagerRequest{
		PlayerUsername: "MeytonPanning",
		AdminUsername:  "WillBelichick",
	}

	err := CreateManagerRequest(newRequest)
	if err == nil {
		t.Errorf("Error should have been produced since sender does not have valid role.")
	}

	DBConn.Exec("DELETE FROM users")
	// DBConn.Exec("DELETE FROM manager_requests")
}

func Test_CreateManagerRequest_InvalidReceiver(t *testing.T) {
	initTestDB()

	DBConn.Create(&User{
		Username: "MeytonPanning",
		Role:     PLAYER,
	})
	DBConn.Create(&User{
		Username: "WillBelichick",
		Role:     MANAGER,
	})

	newRequest := &ManagerRequest{
		PlayerUsername: "MeytonPanning",
		AdminUsername:  "WillBelichick",
	}

	err := CreateManagerRequest(newRequest)
	if err == nil {
		t.Errorf("Error should have been produced since receiver does not have valid role.")
	}

	DBConn.Exec("DELETE FROM users")
	// DBConn.Exec("DELETE FROM manager_requests")
}
