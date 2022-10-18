package model

import (
	"image"
	"time"

	"gorm.io/gorm"
)

// users.go -> database CRUDing for users

// Enum definition for user roles
type userRole uint

const (
	PLAYER userRole = iota
	MANAGER
	ADMIN
)

// enum definition for player positions
type playerPosition uint

const (
	// offense
	QUARTERBACK playerPosition = iota
	RUNNING_BACK
	FULLBACK
	WIDE_REC
	TIGHT_END
	OFF_TACKLE
	OFF_GUARD
	CENTER

	// defense
	NOSE_TACKLE
	DEF_TACKLE
	DEF_END
	MID_LINEBACK
	OUT_LINEBACK
	CORNERBACK
	FREE_SAFETY
	STRONG_SAFETY

	// special teams
	KICKER
	PUNTER
	LONG_SNAP
	HOLDER
	KICK_RETURN
	PUNT_RETURN
)

// corresponds to users table in DB
type User struct {
	UserID uint
	TeamID uint

	Username string `gorm:"unique;not null"`
	// not sure if this should be in separate table
	Password string `gorm:"not null"`
	Email    string

	Firstname string
	Lastname  string

	Role userRole `gorm:"not null"`

	Position playerPosition

	// biographical info, could add more
	Height uint
	Weight uint

	ProfPic image.Image

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
