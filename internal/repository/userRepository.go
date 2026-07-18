package repository

import (
	"digital-bank-api/internal/dto/request"
	"digital-bank-api/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Login(req request.LoginRequest) (*models.User, error) {
	var user models.User

	return &user, r.db.First(&user, "email = ?", req.Email).Error
}

func (r *UserRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	
	return &user, r.db.First(&user, "id = ?", id).Error
}