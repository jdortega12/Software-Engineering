package model

import (
	"testing"
)

// ensure InsertTeamNotification can insert a valid notification
func TestInsertGood(t *testing.T) {
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

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
	var err error
	DBConn, err = InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	m := make(map[string]string)
	m["Message"] = "Pls Work"
	m["SenderID"] = "1"
	if InsertTeamNotification(m) != 1 {
		t.FailNow()
	}
}
