package model

import (
	"testing"
	"time"
)

// see if we can insert a match into the DB error free
func TestCreateMatch(t *testing.T) {
	match := Match {
		MatchType: REGULAR,
		Location: "Twitter HQ",
		StartTime: time.Now(),
		HomeTeamID: 1, 
		AwayTeamID: 2,
		HomeTeamScore: 20,
		AwayTeamScore: 10,
	}

	err := CreateMatch(&match)

	if err != nil {
		t.Error(err)
	}
}

// insert a match and see if we can retrieve it by the ID of the home team 
func TestGetMatchByHomeTeam(t *testing.T) {
	match := Match {
		MatchType: REGULAR,
		Location: "Twitter HQ",
		StartTime: time.Now(),
		HomeTeamID: 1, 
		AwayTeamID: 2,
		HomeTeamScore: 20,
		AwayTeamScore: 10,
	}

	DBConn.Create(&match)

	matches, err := GetMatchesByTeam(1)

	if err != nil {
		t.Error(err)
	}

	if matches[0].HomeTeamID != 1 {
		t.Error("something went wrong")
	}

	cleanUpDB()
}


// insert a match and see if we can retrieve it by the ID of the away team
func TestGetMatchByAwayTeam(t *testing.T) {
	match := Match {
		MatchType: REGULAR,
		Location: "Twitter HQ",
		StartTime: time.Now(),
		HomeTeamID: 1, 
		AwayTeamID: 2,
		HomeTeamScore: 20,
		AwayTeamScore: 10,
	}

	DBConn.Create(&match)

	_, err := GetMatchesByTeam(2)

	if err != nil {
		t.Error(err)
	}

	cleanUpDB()
}

// test if we can retrieve all inserted matches from the DB
func TestGetMatchesThisSeason(t *testing.T) { 
	match := Match {
		MatchType: REGULAR,
		Location: "Twitter HQ",
		StartTime: time.Now(),
		HomeTeamID: 1, 
		AwayTeamID: 2,
		HomeTeamScore: 20,
		AwayTeamScore: 10,
	}

	match2 := Match {
		MatchType: REGULAR,
		Location: "Metlife Stadium",
		StartTime: time.Now(),
		HomeTeamID: 3, 
		AwayTeamID: 4,
		HomeTeamScore: 5,
		AwayTeamScore: 28,
	}

	DBConn.Create(&match)
	DBConn.Create(&match2)

	matches, err := GetMatchesThisSeason()

	if err != nil {
		t.Error(err)
	}

	if len(matches) != 2 {
		t.Error("We didn't catch em all :(")
	}
}