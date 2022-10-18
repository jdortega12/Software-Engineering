package model

// teams.go -> database CRUDing for teams

type team struct {
	teamID uint

	name string `gorm:"unique"`

	// other data maybe
}
