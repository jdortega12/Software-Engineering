package model

import (
	"time"

	"gorm.io/gorm"
)

// teams.go -> database CRUDing for teams

// Corresponds to teams table in DB.
type Team struct {
	ID           uint
	Name         string `gorm:"unique;not null" json:"team_name"`
	TeamLocation string `json:"team_location"`
	// TeamManager uint (maybe??)

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Creates a team in the DB. Returns error if one occurred.
func CreateTeam(team *Team) error {
	err := DBConn.Create(team).Error
	return err
}

// Soft deletes a team from the DB.
func DeleteTeam(team *Team) error {
	err := DBConn.Where("id = ?", team.ID).Delete(team).Error
	return err
}
