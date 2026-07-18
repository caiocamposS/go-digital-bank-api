package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	AccoutID uint   `gorm:"index;not null"`
	Type     string `gorm:"not null"`
	Amount   int64  `gorm:"not null"`
}
