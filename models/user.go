package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string `json:"user"`
	Email        string `json:"mail"`
	PasswordHash []byte
	PasswordSalt []byte
}

// --- DTO Model struct
/// Login User

/// Register User

/// User Exists
