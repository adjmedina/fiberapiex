package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique" json:"user"`
	Email        string `json:"mail"`
	PasswordHash []byte
	PasswordSalt []byte
}

/// Login User

/// Register User

/// User Exists
