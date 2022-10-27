package model

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// notifications.go -> functionality for CRUDing different types of
// notification in DB

// Struct for table with invites to teams and requests to join teams
type TeamNotification struct {
	ID uint

	// if manager, team invite, if player, team request
	SenderID   uint
	ReceiverID uint

	Message string

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Takes a map containing a message, sender, and reciever ID and insert it
// into the DB as a TeamNotification
// returns 0 on successful insertion, 1 otherwise
func InsertTeamNotification(data map[string]string) int {
	message, checkMessage := data["Message"]
	senderID, checkSender := data["SenderID"]
	recieverID, checkReciever := data["ReceiverID"]

	if !checkMessage || !checkSender || !checkReciever {
		return 1
	}

	senderID_int, err := strconv.ParseUint(senderID, 10, 64)

	if err != nil {
		panic(err)
	}

	receiverID_int, err := strconv.ParseUint(recieverID, 10, 64)

	if err != nil {
		panic(err)
	}

	newNotification := TeamNotification{
		Message:    message,
		SenderID:   uint(senderID_int),
		ReceiverID: uint(receiverID_int),
	}

	fmt.Println(DBConn)

	result := DBConn.Create(&newNotification)

	if result.Error != nil {
		panic(result.Error)
	}

	return 0
}
