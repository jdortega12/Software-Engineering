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
// then automigrates all tables. Returns pointer to db connection and error
// if any operations fail.
func InitDB(path string) (*gorm.DB, error) {
	DBConn, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		return nil, err
	}

	err = DBConn.AutoMigrate(
		&User{},
		&UserPersonalInfo{},
		&Team{},
		&Match{},
		&TeamNotification{},
	)

	return DBConn, err
}
