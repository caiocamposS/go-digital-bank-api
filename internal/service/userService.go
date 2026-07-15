package service

import (
	"digital-bank-api/internal/auth"
	"digital-bank-api/internal/dto/request"
	"digital-bank-api/internal/dto/response"
	"digital-bank-api/internal/models"
	"digital-bank-api/internal/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(req request.UserRequest) (*response.UserResponse, error) {
	// hash the pass
	hash, err := auth.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Name: req.Name,
		Email: req.Email,
		Password: string(hash),
	}

	err = s.userRepository.Create(&user)

	if err != nil {
		return nil, err
	}

	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
	}

	return &userResponse, nil
}