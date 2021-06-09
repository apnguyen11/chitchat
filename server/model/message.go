package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )

// username
type Message struct {
	gorm.Model

	user-id int

	channel-id int
	
	content string `json:"content"`

}

