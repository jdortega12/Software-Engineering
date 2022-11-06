package model

import (
	"fmt"
	"testing"
)

// Tests that CreateTeamNotification() returns no errors
// in the case of a valid invite from a manager to a player.
func Test_CreateTeamNotification_ValidInvite(t *testing.T) {
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

	cleanUpDB()
}

// Tests that CreateTeamNotification() produces no error in the case
// of a valid request from a player to a manager.
func Test_CreateTeamNotification_ValidRequest(t *testing.T) {
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

	cleanUpDB()
}

// Tests that CreateTeamNotification() produces an error when
// the sender and receiver are both players.
func Test_CreateTeamNotification_BothPlayers(t *testing.T) {
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

	cleanUpDB()
}

// Tests that CreateTeamNotification() produces an error when
// the sender and receiver are both managers.
func Test_CreateTeamNotification_BothManagers(t *testing.T) {
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

	cleanUpDB()
}

// Tests that CreateTeamNotification() returns an error when
// both sender and receiver are admins.
func Test_CreateTeamNotification_BothAdmins(t *testing.T) {
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

	cleanUpDB()
}

// Tests that CreateTeamNotification() returns an error when
// SenderUsername is not the same as the logged in user.
func Test_CreateTeamNotification_InvalidSender(t *testing.T) {
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

	cleanUpDB()
}

// Tests that CreateTeamNotification() produces an error when the
// ReceiverUsername is not valid.
func Test_CreateTeamNotification_InvalidReceiver(t *testing.T) {
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

	cleanUpDB()
}

// Tests that CreatePromotionToManagerRequest() returns no errors when all
// required fields are present.
func Test_CreatePromotionToManagerRequest_Valid(t *testing.T) {
	err := CreatePromotionToManagerRequest(&PromotionToManagerRequest{
		SenderID:       0,
		SenderUsername: "doesnt matter",
	})

	if err != nil {
		t.Error(err)
	}

	cleanUpDB()
}

// Test that GetPromoToManReqBySendUsername returns no error when the request exists in the DB.
func Test_GetPromoToManReqBySendUsername_Exists(t *testing.T) {
	DBConn.Create(&PromotionToManagerRequest{
		SenderUsername: "jaluhrman",
	})

	_, err := GetPromoToManReqBySendUsername("jaluhrman")
	if err != nil {
		t.Error(err)
	}

	cleanUpDB()
}

// Tests that func returns an error when the user doesn't exist.
func Test_GetPromoToManReqBySendUsername_DoesntExist(t *testing.T) {
	_, err := GetPromoToManReqBySendUsername("jaluhrman")
	if err == nil {
		t.Error("should have produced an error when user doesn't exist")
	}
}

// Tests that func correctly retrieves all PromotionToManagerRequests.
func Test_GetAllPromotionToManagerRequests_Multiple(t *testing.T) {
	requests := []*PromotionToManagerRequest{
		{
			SenderID:       5,
			SenderUsername: "user 5",
			Message:        "test",
		},
		{
			SenderID:       3,
			SenderUsername: "jacob",
			Message:        "oof",
		},
		{
			SenderID:       1,
			SenderUsername: "teepee",
			Message:        "goof",
		},
	}
	DBConn.Create(&requests)

	requestsPulled, err := GetAllPromotionToManagerRequests()
	if err != nil {
		t.Error(err)
	}

	if len(requestsPulled) != 3 {
		t.Error("not all requests were returned")
	}

	for i := range requestsPulled {
		requestsPulled[i].CreatedAt = requests[i].CreatedAt
		requestsPulled[i].UpdatedAt = requests[i].UpdatedAt
		requestsPulled[i].DeletedAt = requests[i].DeletedAt

		if *requestsPulled[i] != *requests[i] {
			t.Errorf("requests at index %d where not equal", i)
		}
	}

	cleanUpDB()
}

// Tests that a good amount of requests can be pulled at once.
func Test_GetAllPromotionToManagerRequests_LotsOfThem(t *testing.T) {
	requests := make([]*PromotionToManagerRequest, 100)

	for i := 0; i < 100; i++ {
		requests[i] = &PromotionToManagerRequest{
			SenderID:       uint(i),
			SenderUsername: "user_" + fmt.Sprint(i),
			Message:        "message " + fmt.Sprint(i),
		}
	}
	DBConn.Create(&requests)

	requestsPulled, err := GetAllPromotionToManagerRequests()
	if err != nil {
		t.Error(err)
	}

	if len(requestsPulled) != len(requests) {
		t.Errorf("requests inserted: %d, requests returned: %d",
			len(requests), len(requestsPulled))
	}

	for i := range requestsPulled {
		requestsPulled[i].CreatedAt = requests[i].CreatedAt
		requestsPulled[i].UpdatedAt = requests[i].UpdatedAt
		requestsPulled[i].DeletedAt = requests[i].DeletedAt

		if *requestsPulled[i] != *requests[i] {
			t.Errorf("requests at index %d where not equal", i)
		}
	}

	cleanUpDB()
}
