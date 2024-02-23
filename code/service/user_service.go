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
	PasswordLogin(dto.PasswordLoginDto) (*entity.User, error)
	SaveUser(string, *dto.UserInfo) (*entity.User, error)
	UpdateMobile(mobile, newMobile string) (*entity.User, error)
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

func (us *userService) SaveUser(mobile string, info *dto.UserInfo) (*entity.User, error) {
	// TODO: create a service or sth for this!
	password, err := bcrypt.GenerateFromPassword([]byte(info.Password), 6)
	if err != nil {
		return nil, errors.New("error in encrypting password")
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

	return user, saveErr
}

func (us *userService) PasswordLogin(userInfo dto.PasswordLoginDto) (*entity.User, error) {
	user, err := us.GetUserByMobile(userInfo.Mobile)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password))

	if passErr != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}
func (us *userService) UpdateMobile(mobile, newMobile string) (*entity.User, error) {
	user, err := us.GetUserByMobile(mobile)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	user.Mobile = newMobile
	err = us.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
