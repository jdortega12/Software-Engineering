package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// notifications.go -> functionality for CRUDing different types of
// notification in DB

// Corresponds to team_notifications table in DB. If SenderID is a manager,
// it is a team invite and the ReceiverID must belong to a player. If the
// roles are reversed, it is a team request. Neither can belong to users with
// the same role, and neither can belong to an Admin. SenderUsername and
// ReceiverUsername are included to make converting to/from JSON easier.
type TeamNotification struct {
	ID uint

	SenderID   uint `gorm:"not null" json:"sender_id"`
	ReceiverID uint `gorm:"not null" json:"receiver_id"`

	SenderUsername   string `gorm:"not null" json:"sender_username"`
	ReceiverUsername string `gorm:"not null" json:"receiver_username"`

	Message string

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Creates a TeamNotification in the DB. Makes sure that the receiver exists,
// neither sender nor receiver are admins, and that that they are not both players
// or managers. Returns error if one occurred.
func CreateTeamNotification(teamNotification *TeamNotification) error {
	sender, err := getUserByUsername(teamNotification.SenderUsername)
	if err != nil {
		return err
	}

	receiver, err := getUserByUsername(teamNotification.ReceiverUsername)
	if err != nil {
		return err
	}

	if sender.Role == ADMIN || receiver.Role == ADMIN {
		return errors.
			New("team notifications: neither sender nor receiver can be ADMIN")
	}

	if sender.Role == receiver.Role {
		return errors.
			New("team notifications: sender and receiver cannot both be player or manager")
	}

	err = DBConn.Create(teamNotification).Error
	return err
}
