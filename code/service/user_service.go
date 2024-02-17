package service

import (
	"github.com/masoud-mohajeri/kea-backend/entity"
	"github.com/masoud-mohajeri/kea-backend/repository"
)

type UserService interface {
	GetUserByMobile(string) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository,
	}
}

func (us *userService) GetUserByMobile(mobile string) (*entity.User, error) {
	return us.userRepository.GetUserByMobile(mobile)
}
