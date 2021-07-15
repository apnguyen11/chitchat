package model

import (
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
)

// username
type Message struct {
	gorm.Model
	Channel string `json:"channel"`
	Content string `json:"content"`
	UserID  int
	User    User
}
