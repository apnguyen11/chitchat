package model

import (
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
  )

// username
type Message struct {
	gorm.Model

	Username string `json:"username"`

	Channel string `json:"channel"`
	
	Content string `json:"content"`

}

