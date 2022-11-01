package controller

import "jdortega12/Software-Engineering/GoServerApp/model"

// wrappers.go -> data structures and funcs for easier controller data handling

// Wrapper for a User and a UserPersonalInfo struct to make JSON serialization easier.
type userDataWrapper struct {
	User         model.User             `json:"user"`
	PersonalInfo model.UserPersonalInfo `json:"personal_info"`
	TeamName     string                 `json:"team_name"`
}
