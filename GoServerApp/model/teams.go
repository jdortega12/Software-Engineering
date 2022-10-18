package model

import (
	"time"

	"gorm.io/gorm"
)

// teams.go -> database CRUDing for teams

type team struct {
	teamID uint

	name string `gorm:"unique;not null"`

	// other data maybe

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
