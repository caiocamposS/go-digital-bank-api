package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	UserID  uint   `gorm:"index;not null"`
	Number  string `gorm:"uniqueIndex;not null"`
	Balance int64  `gorm:"default:0"`
	Type    string `gorm:"not null"` // (checking, savings)
}
