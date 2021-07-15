package model

// "gorm.io/driver/sqlite"

// username
type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
