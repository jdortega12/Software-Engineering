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

type match struct {
	matchID  uint
	seasonID uint

	matchType matchType `gorm:"not null"`

	// probably should just be the name of the stadium or whatever,
	// could add street num/name, state, zip, etc.
	location string `gorm:"not null"`

	// date AND time
	startTime time.Time `gorm:"not null"`
	endTime   time.Time

	homeTeamID uint `gorm:"not null"`
	awayTeamID uint `gorm:"not null"`

	homeTeamScore uint
	awayTeamScore uint

	likes    uint
	dislikes uint

	// player likes and dislikes? Not sure if those
	// are meant to be for a match only or just on the
	// player's profile permanently or both

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
