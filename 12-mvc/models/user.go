package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Address  string
	Role     string
}
