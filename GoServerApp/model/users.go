package model

import (
	"time"
	"fmt"
	"gorm.io/gorm"
)

// users.go -> database CRUDing for users

// Enum definition for user roles
type userRole uint

const (
	PLAYER userRole = iota
	MANAGER
	ADMIN
)

// enum definition for player positions
type playerPosition uint

const (
	// offense
	QUARTERBACK playerPosition = iota
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

// corresponds to users table in DB
type User struct {
	UserID uint
	TeamID uint

	Username string `gorm:"unique;not null"`
	// not sure if this should be in separate table
	Password string `gorm:"not null"`
	Email    string

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
	UserPersonalInfoID uint

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	Height uint `json"height,string"`  // inches
	Weight uint `json:"weight,string"` // lbs

	// metadata
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Updates the personal info of a user in the DB.
func UpdateUserPersonalInfo(userPersInfo *UserPersonalInfo) error {
	err := DBConn.Where("user_personal_info_id = ?", userPersInfo.UserPersonalInfoID).
		Updates(&userPersInfo).Error

	return err
}

// Takes a User struct
// into the DB as a User
// returns 0 on successful insertion, 1 otherwise
func CreateUser(user *User) error {
	err := DBConn.Create(user).Error
	if err != nil {
		return err
	}

	personalInfo := &UserPersonalInfo{
		UserPersonalInfoID: user.UserID,
	}

	err = DBConn.Create(personalInfo).Error
	if err != nil {
		DBConn.Delete(user)
		return err
	}
	return err
}


// Takes photo as base64 String, username, password 
// updates user photo to the new phot 
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


func GetUserId(username string) (uint, error) {
	user := User{} 
	result := DBConn.Where("username = ?", username).First(&user)
	fmt.Println(user) 

	if result.Error != nil {
		return 5, result.Error
	}

	return user.UserID, result.Error
}
