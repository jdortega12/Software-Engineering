package model

import (
	"time"

	"gorm.io/gorm"
)

// matches.go -> database CRUDing for matches

// enum for match type
type matchType uint

const (
	REGULAR matchType = iota
	PLAYOFF
	FINAL
)

// Corresponds to matches table in DB.
type Match struct {
	ID       uint `json:"id"`
	SeasonID uint `json:"season_id"`

	MatchType matchType `gorm:"not null" json:"match_type"`

	// probably should just be the name of the stadium or whatever,
	// could add street num/name, state, zip, etc.
	Location string `gorm:"not null" json:"location"`

	// date AND time
	StartTime time.Time `gorm:"not null" json:"start_time"`
	EndTime   time.Time `json:"end_time"`

	// whether match is currently in progress or archived
	InProgress bool `json:"in_progress"`

	// quarter and how much time is left within the quarter
	Quarter     uint      `json:"quarter"`
	QuarterTime time.Time `json:"quarter_time"`

	HomeTeamID uint `gorm:"not null" json:"home_id"`
	AwayTeamID uint `gorm:"not null" json:"away_id"`

	HomeTeamScore uint `json:"home_team_score"`
	AwayTeamScore uint `json:"away_team_score"`

	Likes    uint `json:"likes"`
	Dislikes uint `json:"dislikes"`


	// player likes and dislikes? Not sure if those
	// are meant to be for a match only or just on the
	// player's profile permanently or both

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Insert a match into the database
func CreateMatch(match *Match) error {
	err := DBConn.Create(match).Error
	return err
}

// Retrieve all matches involving a given team ID
func GetMatchesByTeam(id uint) ([]Match, error) {
	matches := []Match{}
	err := DBConn.Where("home_team_id = ? OR away_team_id = ?", id, id).Find(&matches).Error
	return matches, err
}


// Retrieve all matches
func GetMatchesThisSeason() ([]Match, error) {
	matches := []Match{}
	err := DBConn.Find(&matches).Error
	return matches, err
}

// Get match by id
func GetMatchById(id uint) (Match, error) {
	match := Match{}
	err := DBConn.Where("id = ?", id).First(&match).Error
	return match, err
}

// Marks a match as finished in the DB by updating InProgress and EndTime.
func FinishMatch(id uint) error {
	// make sure match exists
	match := &Match{}
	err := DBConn.Where("id = ?", id).First(match).Error
	if err != nil {
		return err
	}

	err = DBConn.Model(&Match{}).Where("id = ?", id).Updates(map[string]interface{}{
		"in_progress": false,
		"end_time":    time.Now(),
	}).Error

	return err
}

func AddLike(id uint) error{
	match := &Match{}
	err := DBConn.Where("id = ?", id).First(match).Error
	if err != nil {
		return err
	}
	err = DBConn.Model(&Match{}).Where("id = ?", id).Update("likes", match.Likes + 1).Error
	return err
}

func AddDislike(id uint) error{
	match := &Match{}
	err := DBConn.Where("id = ?", id).First(match).Error
	if err != nil {
		return err
	}
	err = DBConn.Model(&Match{}).Where("id = ?", id).Update("dislikes", match.Dislikes + 1).Error
	return err
}
