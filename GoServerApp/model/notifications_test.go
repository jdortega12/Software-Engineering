package model

import "testing"

// Tests that CreateTeamNotification returns no errors
// in the case of a valid invite from a manager to a player.
func Test_CreateTeamNotification_ValidInvite(t *testing.T) {
	DBConn, _ = InitDB(TEST_DB_PATH)

	DBConn.Create(&User{
		Username: "jaluhrman",
		Role:     MANAGER,
	})
	DBConn.Create(&User{
		Username: "colbert",
		Role:     PLAYER,
	})

	notification := &TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	err := CreateTeamNotification(notification)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}

// Tests that CreateTeamNotification produces no error in the case
// of a valid request from a player to a manager.
func Test_CreateTeamNotification_ValidRequest(t *testing.T) {
	DBConn, _ = InitDB(TEST_DB_PATH)

	DBConn.Create(&User{
		Username: "jaluhrman",
		Role:     PLAYER,
	})
	DBConn.Create(&User{
		Username: "colbert",
		Role:     MANAGER,
	})

	notification := &TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	err := CreateTeamNotification(notification)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}

// Tests that CreateTeamNotification produces an error when
// the sender and receiver are both players.
func Test_CreateTeamNotification_BothPlayers(t *testing.T) {
	DBConn, _ = InitDB(TEST_DB_PATH)

	DBConn.Create(&User{
		Username: "jaluhrman",
		Role:     PLAYER,
	})
	DBConn.Create(&User{
		Username: "colbert",
		Role:     PLAYER,
	})

	notification := &TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	err := CreateTeamNotification(notification)
	if err == nil {
		t.Error("Error should have been produced when both users are players")
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}

func Test_CreateTeamNotification_BothManagers(t *testing.T) {
	DBConn, _ = InitDB(TEST_DB_PATH)

	DBConn.Create(&User{
		Username: "jaluhrman",
		Role:     MANAGER,
	})
	DBConn.Create(&User{
		Username: "colbert",
		Role:     MANAGER,
	})

	notification := &TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	err := CreateTeamNotification(notification)
	if err == nil {
		t.Error("Error should have been produced when both users are managers")
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}

func Test_CreateTeamNotification_BothAdmins(t *testing.T) {
	DBConn, _ = InitDB(TEST_DB_PATH)

	DBConn.Create(&User{
		Username: "jaluhrman",
		Role:     ADMIN,
	})
	DBConn.Create(&User{
		Username: "colbert",
		Role:     ADMIN,
	})

	notification := &TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	err := CreateTeamNotification(notification)
	if err == nil {
		t.Error("Error should have been produced when both users are admins")
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}

func Test_CreateTeamNotification_InvalidSender(t *testing.T) {
	DBConn, _ = InitDB(TEST_DB_PATH)

	DBConn.Create(&User{
		Username: "colbert",
		Role:     MANAGER,
	})

	notification := &TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	err := CreateTeamNotification(notification)
	if err == nil {
		t.Error("Error should have been produced when sender isn't valid")
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}

func Test_CreateTeamNotification_InvalidReceiver(t *testing.T) {
	DBConn, _ = InitDB(TEST_DB_PATH)

	DBConn.Create(&User{
		Username: "colbert",
		Role:     MANAGER,
	})

	notification := &TeamNotification{
		SenderUsername:   "colbert",
		ReceiverUsername: "jaluhrman",
	}

	err := CreateTeamNotification(notification)
	if err == nil {
		t.Error("Error should have been produced when receiver isn't valid")
	}

	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM team_notifications")
}
