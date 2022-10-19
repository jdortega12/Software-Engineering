package model

import (
	"time"

	"gorm.io/gorm"
)

// notifications.go -> functionality for CRUDing different types of
// notification in DB

// Struct for table with invites to teams and requests to join teams
type TeamNotification struct {
	TeamNotificationID uint

	// if manager, team invite, if player, team request
	SenderID   uint
	ReceiverID uint

	Message string

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
