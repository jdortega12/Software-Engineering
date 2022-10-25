package model

import (
	"testing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	TEST_DB_PATH = "../test.db"
)

// Checks that InitDB() does not return any errors.
func TestInitDB(t *testing.T) {
	err := InitDB(TEST_DB_PATH)
	if err != nil {
		t.Error(err)
	}
}

// ensure InsertTeamNotification can insert a valid notification
func TestInsertGood(t *testing.T) {
	TDBConn, err := gorm.Open(sqlite.Open("../test.db"))

	if err != nil {
		panic(err)
	}
	
	InitDBTest(TDBConn)
	
	m := make(map[string]string)
	m["Message"] = "Pls Work"
	m["SenderID"] = "1"
	m["ReceiverID"] = "2"
	if InsertTeamNotification(m) != 0 {
		t.FailNow()
	}
}


// ensure InsertTeamNotification can insert an invalid notification
func TestInsertBad(t *testing.T) {
	TDBConn, err := gorm.Open(sqlite.Open("../test.db"))

	if err != nil {
		panic(err)
	}
	
	InitDBTest(TDBConn)
	
	m := make(map[string]string)
	m["Message"] = "Pls Work"
	m["SenderID"] = "1"
	if InsertTeamNotification(m) != 1 {
		t.FailNow()
	}
}
