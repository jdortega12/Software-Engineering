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

// Creates/opens the database specified by the path argument.
// If successful, then automigrates all tables. Returns error
// if either of these operations fail. Currently not set up to
// work with anything other than local sqlite db.
func InitDB(path string) (*gorm.DB, error) {
	// open db and panic if cannot
	DBConn, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		return nil, err
	}

	// auto-migrate all tables. If DB already existed,
	// AutoMigrate() will update the tables with any changed
	// fields if the corresponding structs have changed.
	err = DBConn.AutoMigrate(
		&User{},
		&Team{},
		&Match{},
		&TeamNotification{},
	)

	return DBConn, err
}
