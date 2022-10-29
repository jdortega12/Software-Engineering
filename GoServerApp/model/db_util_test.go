package model

import (
	"testing"
)

const (
	TEST_DB_PATH = "file::memory:?cache=shared"
)

// Just a wrapper for InitDB() and error handling to save
// space in tests. Also declared in handler.
func initTestDB() {
	err := InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}
}

// Checks that InitDB() does not return any errors when given a valid path.
func Test_InitDB(t *testing.T) {
	err := InitDB(TEST_DB_PATH)
	if err != nil {
		t.Error(err)
	}
}

// Checks that InitDB() does return an error when the path given is bad.
func Test_InitDB_BadPath(t *testing.T) {
	err := InitDB(" ")
	if err == nil {
		t.Error(err)
	}
}
