package model

// "gorm.io/driver/sqlite"

// username
type LoginResponse struct {
	Success bool `json:"success"`
}

type WhoAmIResponse struct {
	Username string `json:"username"`
}
