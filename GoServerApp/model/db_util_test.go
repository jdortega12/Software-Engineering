package model

import (
	"testing"
)

const (
	TEST_DB_PATH = "file::memory:?cache=shared"
)

// Checks that InitDB() does not return any errors when
// given a valid path.
func Test_InitDB(t *testing.T) {
	var err error
	_, err = InitDB(TEST_DB_PATH)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

// Checks that InitDB() does return an error when the
// path given is bad.
func Test_InitDB_BadPath(t *testing.T) {
	var err error
	_, err = InitDB(" ")
	if err == nil {
		t.Error(err)
		t.FailNow()
	}
}
