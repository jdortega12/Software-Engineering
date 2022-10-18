package model

import (
	"time"

	"gorm.io/gorm"
)

// teams.go -> database CRUDing for teams

type Team struct {
	TeamID uint

	Name string `gorm:"unique;not null"`

	// other data maybe

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
