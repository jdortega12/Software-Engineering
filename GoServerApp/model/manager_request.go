package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type ManagerRequest struct {
	ID uint

	PlayerID uint `gorm:"not null" json:"player_id"`
	AdminID  uint `gorm:"not null" json: "admin_id"`

	PlayerUsername string `gorm:"not null" json:"player_username"`
	AdminUsername  string `gorm:"not null" json:"admin_username"`

	// Message string might be needed in request

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func CreateManagerRequest(managerRequest *ManagerRequest) error {
	player, err := GetUserByUsername(managerRequest.PlayerUsername)
	if err != nil {
		return err
	}

	admin, err := GetUserByUsername(managerRequest.AdminUsername)
	if err != nil {
		return err
	}

	// If request is not being sent from a player account
	if player.Role != PLAYER {
		return errors.New("Manager request: Request for manager only permitted by PLAYER accounts.")
	}

	// If request is not being sent to manager account
	if admin.Role != ADMIN {
		return errors.New("Manager request: Request for manager must be sent to administrator.")
	}

	err = DBConn.Create(managerRequest).Error
	return err
}
