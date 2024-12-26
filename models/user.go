package models

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:USER"`
}
