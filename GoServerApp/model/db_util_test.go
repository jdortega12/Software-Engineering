package model

import (
	"testing"
)

const (
	TEST_DB_PATH = "../test.db"
)

// Checks that InitDB() does not return any errors.
func TestInitDB(t *testing.T) {
	var err error
	_, err = InitDB(TEST_DB_PATH)
	if err != nil {
		t.Error(err)
	}
}
