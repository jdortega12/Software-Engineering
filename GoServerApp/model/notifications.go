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
	ID uint `json:"-"`

	SenderID   uint `gorm:"not null" json:"sender_id"`
	ReceiverID uint `gorm:"not null" json:"receiver_id"`

	SenderUsername   string `gorm:"not null" json:"sender_username"`
	ReceiverUsername string `gorm:"not null" json:"receiver_username"`

	Message string `json:"message"`

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Corresponds to promotion_to_manager_requests table in DB.
type PromotionToManagerRequest struct {
	ID uint `json:"-"`

	SenderID       uint   `gorm:"not null" json:"-"`
	SenderUsername string `gorm:"not null" json:"sender_username"`

	Message string `json:"message"`

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Creates a TeamNotification in the DB. Makes sure that the receiver exists,
// neither sender nor receiver are admins, and that that they are not both players
// or managers. Returns error if one occurred.
func CreateTeamNotification(teamNotification *TeamNotification) error {
	sender, err := GetUserByUsername(teamNotification.SenderUsername)
	if err != nil {
		return err
	}

	receiver, err := GetUserByUsername(teamNotification.ReceiverUsername)
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

// Creates a PromotionToManagerRequest in the DB.
func CreatePromotionToManagerRequest(request *PromotionToManagerRequest) error {
	err := DBConn.Create(request).Error
	return err
}

// Finds a PromotionToManagerRequest by the sender's username.
func GetPromoToManReqBySendUsername(username string) (*PromotionToManagerRequest, error) {
	request := &PromotionToManagerRequest{}
	err := DBConn.Where("sender_username = ?", username).First(request).Error

	return request, err
}

// Returns slice of all PromotionToManagerRequest's in the DB.
func GetAllPromotionToManagerRequests() ([]*PromotionToManagerRequest, error) {
	requests := []*PromotionToManagerRequest{}
	err := DBConn.Find(&requests).Error

	return requests, err
}
