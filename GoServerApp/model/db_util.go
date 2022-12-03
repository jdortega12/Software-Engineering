package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// db_util.go -> contains functionality for initializing db,
// automigrating tables, etc.

var (
	// connection to the database (can be local or remote)
	DBConn *gorm.DB
)

// Creates/opens the database specified by the path parameter. If successful,
// then automigrates all tables. Returns error if any operations fail.
func InitDB(path string) error {
	var err error
	DBConn, err = gorm.Open(sqlite.Open(path))
	if err != nil {
		return err
	}

	err = DBConn.AutoMigrate(
		&User{},
		&UserPersonalInfo{},
		&Team{},
		&Match{},
		&TeamNotification{},
		&PromotionToManagerRequest{},
		&Comment{},
	)

	return err
}
