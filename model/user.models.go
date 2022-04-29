package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"uniqueIndex;not null"`
	Role     string `json:"role" gorm:"uniqueIndex;not null;default:'Member'"`
}

type RequestLogin struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
