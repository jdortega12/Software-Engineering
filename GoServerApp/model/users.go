package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// users.go -> database CRUDing for users

// Enum definition for user roles. Defaults in DB to PLAYER.
type userRole uint

const (
	PLAYER userRole = iota
	MANAGER
	ADMIN
)

// Enum definition for player positions. Defaults to NULL in DB.
type playerPosition uint

const (
	// NULL for non-players or players not on team
	NULL playerPosition = iota

	// offense
	QUARTERBACK
	RUNNING_BACK
	FULLBACK
	WIDE_REC
	TIGHT_END
	OFF_TACKLE
	OFF_GUARD
	CENTER

	// defense
	NOSE_TACKLE
	DEF_TACKLE
	DEF_END
	MID_LINEBACK
	OUT_LINEBACK
	CORNERBACK
	FREE_SAFETY
	STRONG_SAFETY

	// special teams
	KICKER
	PUNTER
	LONG_SNAP
	HOLDER
	KICK_RETURN
	PUNT_RETURN
)

// Corresponds to users table in DB.
type User struct {
	ID     uint
	TeamID uint

	Username string `gorm:"unique;not null" json:"username"`
	Email    string `json:"email"`
	Password string `gorm:"not null" json:"password"`

	Role userRole `gorm:"not null"`

	Position playerPosition

	Photo string

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Personal info of a user that they can change at any time
// without any permission or extra complication.
type UserPersonalInfo struct {
	// must be same ID as user whom it belongs to
	ID uint

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	Height uint `json:"height,string"` // inches
	Weight uint `json:"weight,string"` // lbs

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Updates the personal info of a user in the DB. Returns error if one ocurred.
func UpdateUserPersonalInfo(userPersInfo *UserPersonalInfo) error {
	err := DBConn.Where("id = ?", userPersInfo.ID).
		Updates(&userPersInfo).Error

	return err
}

// Creates a User in the DB. ALso creates a corresponding UserPersonalInfo
// with the same ID as the User. Returns error if one ocurred.
func CreateUser(user *User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("users: username and password cannot be empty")
	}

	err := DBConn.Create(user).Error
	if err != nil {
		return err
	}

	personalInfo := &UserPersonalInfo{
		ID: user.ID,
	}

	err = DBConn.Create(personalInfo).Error
	if err != nil {
		DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)
	}

	return err
}

// Takes photo as base64 string, username, password, and updates user's photo
func UpdateUserPhoto(photo string, username string, password string) error {
	user := User{}
	err := DBConn.Where("username = ?", username, &user).Error

	if err != nil || user.Password != password {
		return err
	}

	user.Photo = photo
	DBConn.Save(&user)
	return err
}

// Assigns a user to a given TeamID.
func AssignUserToTeam(user *User, teamID uint) error {
	err := DBConn.Model(&User{}).Where("id = ?", user.ID).Update("team_id", teamID).Error
	return err
}

// Pulls User out of DB by username.
func getUserByUsername(username string) (*User, error) {
	user := &User{}
	err := DBConn.Where("username = ?", username).First(user).Error

	return user, err
}
