package service

import (
	"errors"

	"github.com/masoud-mohajeri/kea-backend/dto"
	"github.com/masoud-mohajeri/kea-backend/entity"
	"github.com/masoud-mohajeri/kea-backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserByMobile(string) (*entity.User, error)
	SaveUser(string, *dto.UserInfo) error
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

func (us *userService) SaveUser(mobile string, info *dto.UserInfo) error {
	// TODO: create a service or sth for this!
	password, err := bcrypt.GenerateFromPassword([]byte(info.Password), 6)
	if err != nil {
		return errors.New("error in encrypting password")
	}

	user := &entity.User{
		Mobile:    mobile,
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Password:  string(password),
		// TODO: learn enum for db
		Role: "USER",
	}

	saveErr := us.userRepository.Save(user)

	return saveErr
}
