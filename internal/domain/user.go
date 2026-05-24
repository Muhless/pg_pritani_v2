package domain

import "gorm.io/gorm"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-" gorm:"unique"`
	Role     Role   `json:"role" gorm:"type:varchar(20);default:'user'"`
}
