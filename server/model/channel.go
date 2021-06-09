package model

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )

type Channel struct {
	gorm.Model
	
	Name string `json:"name"`

}

