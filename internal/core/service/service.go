package service

import (
	"zuri-stage-one/internal/core/models"
	"zuri-stage-one/internal/ports"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) ports.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) Get() (*models.User, error) {
	return u.userRepository.Get()
}
