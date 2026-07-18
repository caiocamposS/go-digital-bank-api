package repository

import (
	"digital-bank-api/internal/models"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Create(account *models.Account) error {
	return r.db.Create(account).Error
}