package model

import (
	"time"
	"strconv"
	"gorm.io/gorm"
)


// Corresponds to matches table in DB.
type Comment struct {
	ID       uint `json:"id"`
	Message string `json:"message"`
	MatchID string  `json:"match_id"`
	UserID	uint  `json:"user_id"`
	Username string `json:"username"`

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Insert a comment into the database
func CreateComment(comment *Comment) error {
	err := DBConn.Create(comment).Error
	return err
}

// Retrieve all comments for a given match 
func GetCommentsById(id uint) ([]Comment, error) {
	id_str := strconv.FormatUint(uint64(id), 10)
	comments := []Comment{}
	err := DBConn.Where("match_id = ?", id_str).Find(&comments).Error
	return comments, err
}