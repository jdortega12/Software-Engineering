package model

import (
	"time"
	"gorm.io/gorm"
)

// matches.go -> database CRUDing for matches

// enum for match type
type matchType uint

const (
	REGULAR matchType = iota
	PLAYOFF
	FINAL
)

// Corresponds to matches table in DB.
type Match struct {
	ID       uint `json:"-"`
	SeasonID uint `json:"-"`

	MatchType matchType `gorm:"not null" json:"match_type"`

	// probably should just be the name of the stadium or whatever,
	// could add street num/name, state, zip, etc.
	Location string `gorm:"not null" json:"location"`

	// date AND time
	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time

	HomeTeamID uint `gorm:"not null" json:"home_id"`
	AwayTeamID uint `gorm:"not null" json:"away_id"`

	HomeTeamScore uint
	AwayTeamScore uint

	Likes    uint
	Dislikes uint

	// player likes and dislikes? Not sure if those
	// are meant to be for a match only or just on the
	// player's profile permanently or both

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Insert a match into the database
func CreateMatch(match *Match) error {
	err := DBConn.Create(match).Error
	return err
}

// Retrieve all matches involving a given team ID
func GetMatchesByTeam(id uint) ([]Match, error) {
	matches := []Match{}
	err := DBConn.Where("home_team_id = ? OR away_team_id = ?", id, id).Find(&matches).Error
	return matches, err
}

// Retrieve all matches
func GetMatchesThisSeason() ([]Match, error) {
	matches := []Match{}
	err := DBConn.Find(&matches).Error
	return matches, err
}




