package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type RemovePlayer struct {
	ID uint

	PlayerID  uint `gorm:"not null" json:"player_id"`
	ManagerID uint `gorm:"not null" json: "manager_id"`

	PlayerUsername  string `gorm:"not null" json:"playername"`
	ManagerUsername string `gorm:"not null" json:"managername"`

	//Metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func ManagerRemovePlayer(removePlayer *RemovePlayer) error {
	player, err := GetUserByUsername(removePlayer.PlayerUsername)
	if err != nil {
		return err
	}
	manager, err := GetUserByUsername(removePlayer.ManagerUsername)
	if err != nil {
		return err
	}

	if player.Role != PLAYER {
		return errors.
			New("Player removal: User to be removed must be PLAYER")
	}

	if manager.Role != MANAGER {
		return errors.
			New("Player removal: User must be removed by MANAGER")
	}

	if player.TeamID != manager.TeamID {
		return errors.
			New("Player removal: Player must be on same team as MANAGER")
	}
	err = DBConn.Create(removePlayer).Error
	return err
}
