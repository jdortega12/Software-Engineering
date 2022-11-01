package model

import (
	"time"

	"gorm.io/gorm"
)

// teams.go -> database CRUDing for teams

// Corresponds to teams table in DB.
type Team struct {
	ID           uint   `json:"-"`
	Name         string `gorm:"unique;not null" json:"team_name"`
	TeamLocation string `json:"team_location"`
	// TeamManager uint (maybe??)

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
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

func GetTeamByID(id uint) (*Team, error) {
	team := &Team{}
	err := DBConn.Where("id = ?", id).First(team).Error

	return team, err
}
