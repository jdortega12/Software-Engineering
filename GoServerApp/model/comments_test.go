package model

import (
	"testing"
	//"time"
)

// see if we can insert a match into the DB error free
func TestCreateComment(t *testing.T) {
	comment := Comment{
		Message: "Hello",
		MatchID: "12",
		UserID: 1,
	}

	err := CreateComment(&comment)
	if err != nil {
		t.Error(err)
	}
	comment2 := Comment{
		Message: "World",
		MatchID: "12",
		UserID: 2,
	}
	
	err = CreateComment(&comment2)
	if err != nil {
		t.Error(err)
	}

	cleanUpDB()
}

func TestFindMatchComments(t *testing.T) {
	comment := Comment{
		Message: "Hello",
		MatchID: "12",
		UserID: 1,
	}

	err := CreateComment(&comment)
	if err != nil {
		t.Error(err)
	}
	comment2 := Comment{
		Message: "World",
		MatchID: "12",
		UserID: 2,
	}
	
	err = CreateComment(&comment2)
	if err != nil {
		t.Error(err)
	}

	comments, err := GetCommentsById(uint(12))
	if err != nil {
		t.Error(err)
	}
	if (comments[0].Message != "Hello" || comments[1].Message != "World"){
		t.Error("wrong message")
	}
}
