package repository

import (
	"errors"

	"github.com/masoud-mohajeri/kea-backend/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByMobile(string) (*entity.User, error)
	Save(user *entity.User) error
	Update(user *entity.User) error
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
			return nil, nil
		}
		return nil, errors.New("internal server error")
	}
	return &user, nil
}

func (uc *userConnection) Save(user *entity.User) error {
	tx := uc.connection.Debug().Create(&user)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (uc *userConnection) Update(user *entity.User) error {
	tx := uc.connection.Debug().Model(&entity.User{}).Where("id = ?", user.ID).Updates(user)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
