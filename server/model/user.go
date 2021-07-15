package model

import (
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
)

// username
type User struct {
	gorm.Model

	Username string `json:"username"`

	Password string `json:"password"`
}
