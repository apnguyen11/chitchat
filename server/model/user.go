package model

import (
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite"
  )
// username
type Username struct {
	gorm.Model

	Username string `json:"username"`

	Password string `json:"password"`
}
