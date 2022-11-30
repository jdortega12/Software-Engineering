package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// users.go -> database CRUDing for users

// Enum definition for user roles. Defaults in DB to PLAYER.
type userRole string

const (
	PLAYER  userRole = "player"
	MANAGER userRole = "panager"
	ADMIN   userRole = "admin"
)

// Enum definition for player positions. Defaults to NULL in DB.
type playerPosition string

const (
	// NULL for non-players or players not on team
	NO_POSITION playerPosition = "no position"

	// offense
	QUARTERBACK  playerPosition = "quarterback"
	RUNNING_BACK playerPosition = "running back"
	FULLBACK     playerPosition = "fullback"
	WIDE_REC     playerPosition = "wide receiver"
	TIGHT_END    playerPosition = "tight end"
	OFF_TACKLE   playerPosition = "offensive tackle"
	OFF_GUARD    playerPosition = "offensive guard"
	CENTER       playerPosition = "center"

	// defense
	NOSE_TACKLE   playerPosition = "nose tackle"
	DEF_TACKLE    playerPosition = "defensive tackle"
	DEF_END       playerPosition = "defensive end"
	MID_LINEBACK  playerPosition = "middle linebacker"
	OUT_LINEBACK  playerPosition = "outside linebacker"
	CORNERBACK    playerPosition = "cornerback"
	FREE_SAFETY   playerPosition = "free safety"
	STRONG_SAFETY playerPosition = "strong safety"

	// special teams
	KICKER      playerPosition = "kicker"
	PUNTER      playerPosition = "punter"
	LONG_SNAP   playerPosition = "long snapper"
	HOLDER      playerPosition = "holder"
	KICK_RETURN playerPosition = "kick returner"
	PUNT_RETURN playerPosition = "punt returner"
)

// Corresponds to users table in DB.
type User struct {
	ID     uint `json:"-"`
	TeamID uint `json:"-"`

	Username string `gorm:"unique;not null" json:"username"`
	Email    string `json:"email"`
	Password string `gorm:"not null" json:"password"`

	Role userRole `gorm:"not null;default:player" json:"role"`

	Position playerPosition `gorm:"default:no position" json:"position"`

	Photo string `json:"photo"`

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Personal info of a user that they can change at any time
// without any permission or extra complication.
type UserPersonalInfo struct {
	// must be same ID as user whom it belongs to
	ID uint `json:"-"`

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`

	Height uint `json:"height,string"` // inches
	Weight uint `json:"weight,string"` // lbs

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// Data for Accepting Player
type AcceptData struct {
	PlayerName  string `json:"playername"`
	ManagerName string `json:"managername"`

	// metadata
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// struct for data to send to front end for player display
type UserTeamData struct {
	ID       uint   `json:"-"`
	Teamname string `json:"teamname"`

	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// struct for return data from change roster
type UserTeamReturnData struct {
	UserId   uint   `json:"userid"`
	Teamname string `json:"teamname"`
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
	result := DBConn.First(&user, "username = ?", username)
	err := result.Error

	if err != nil || user.Password != password {
		return err
	}

	user.Photo = photo
	DBConn.Save(&user)

	return err
}

// Updates a user's TeamID.
func UpdateUserTeam(user *User, teamID uint) error {
	err := DBConn.Model(&User{}).Where("id = ?", user.ID).Update("team_id", teamID).Error
	return err
}

// Gathers user and team data based on user
func GatherUserTeamData(user *User) UserTeamData {
	personalInfo, _ := GetUserPersonalInfoByID(user.ID)
	team, _ := GetTeamByID(user.TeamID)
	userTeamData := UserTeamData{
		ID:        user.ID,
		Teamname:  team.Name,
		Firstname: personalInfo.Firstname,
		Lastname:  personalInfo.Lastname,
	}
	return userTeamData
}

// gets all users
func GetUsers() ([]User, error) {
	users := []User{}
	err := DBConn.Find(&users).Error
	return users, err
}

// Pulls User ouf of DB by ID
func GetUserbyID(id uint) (*User, error) {
	user := &User{}
	err := DBConn.Where("id = ?", id).First(user).Error
	return user, err
}

// Pulls User out of DB by username.
func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := DBConn.Where("username = ?", username).First(user).Error
	return user, err
}

// Pulls a user's personal info by user's ID
func GetUserPersonalInfoByID(id uint) (*UserPersonalInfo, error) {
	info := &UserPersonalInfo{}
	err := DBConn.Where("id = ?", id).First(info).Error
	return info, err
}

// Pulls a Manager given a teamid
func GetManagerByTeamID(teamid uint) (User, error) {
	user := User{}
	err := DBConn.Where("team_id = ? AND role = ?", teamid, MANAGER).First(&user).Error

	return user, err
}

// Pulls Players given a teamid
func GetPlayersByTeamID(teamid uint) ([]User, error) {
	users := []User{}
	err := DBConn.Select("Username, Position").Where("team_id = ? AND role = ?", teamid, PLAYER).Find(&users).Error

	return users, err
}
