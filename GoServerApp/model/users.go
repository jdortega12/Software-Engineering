package model

import "image"

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
type user struct {
	userID uint
	teamID uint

	username string `gorm:"unique"`
	password string // not sure if this should be in separate table
	email    string

	firstname string
	lastname  string

	role userRole

	position playerPosition

	// biographical info, could add more
	height uint
	weight uint

	profPic image.Image
}
