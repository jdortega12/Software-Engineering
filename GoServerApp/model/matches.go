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

type Match struct {
	MatchID  uint
	SeasonID uint

	MatchType matchType `gorm:"not null"`

	// probably should just be the name of the stadium or whatever,
	// could add street num/name, state, zip, etc.
	Location string `gorm:"not null"`

	// date AND time
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time

	HomeTeamID uint `gorm:"not null"`
	AwayTeamID uint `gorm:"not null"`

	HomeTeamScore uint
	AwayTeamScore uint

	Likes    uint
	Dislikes uint

	// player likes and dislikes? Not sure if those
	// are meant to be for a match only or just on the
	// player's profile permanently or both

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
