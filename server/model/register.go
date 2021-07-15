package model

// "gorm.io/driver/sqlite"

// username
type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
