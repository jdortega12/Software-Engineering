package model

import (
	"testing"

	"gorm.io/gorm"
)

const (
	TEST_DB_PATH = "file::memory:?cache=shared"
)

// Initializes a DB for testing purposes. Just a wrapper
// for InitDB() and error handling to save space in tests.
// Also declared in handler because go tests are strange.
func initTestDB() *gorm.DB {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	return DBConn
}

// Checks that InitDB() does not return any errors when
// given a valid path.
func Test_InitDB(t *testing.T) {
	_, err := InitDB(TEST_DB_PATH)
	if err != nil {
		t.Error(err)
	}
}

// Checks that InitDB() does return an error when the
// path given is bad.
func Test_InitDB_BadPath(t *testing.T) {
	_, err := InitDB(" ")
	if err == nil {
		t.Error(err)
	}
}
