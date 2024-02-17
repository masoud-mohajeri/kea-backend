package repository

import (
	"errors"

	"github.com/masoud-mohajeri/kea-backend/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByMobile(string) (*entity.User, error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (uc *userConnection) GetUserByMobile(mobile string) (*entity.User, error) {

	var user entity.User
	tx := uc.connection.Debug().Model(&entity.User{}).Where("mobile = ?", mobile).Last(&user)
	if err := tx.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("internal server error")
	}
	return &user, nil
}
