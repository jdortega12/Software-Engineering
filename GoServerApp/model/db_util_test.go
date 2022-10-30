package model

import (
	"os"
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

// Cleans up all DB tables.
func cleanUpDB() {
	DBConn.Exec("DELETE FROM users")
	DBConn.Exec("DELETE FROM user_personal_infos")
	DBConn.Exec("DELETE FROM teams")
	DBConn.Exec("DELETE FROM team_notifications")
	DBConn.Exec("DELETE FROM matches")
}

func TestMain(m *testing.M) {
	initTestDB()
	rc := m.Run()
	os.Exit(rc)
}

/* THESE TWO TESTS DO NOT WORK WITH TEST_MAIN AND CAUSE ERRORS WHEN TRYING TO OPEN DB
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
}*/
