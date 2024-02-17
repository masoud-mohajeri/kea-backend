package repository

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/masoud-mohajeri/kea-backend/entity"
)

type OtpRepository interface {
	Save(*entity.Otp) (*entity.Otp, error)
	FindOne(string) (*entity.Otp, error)
	Attempt(string) error
	Remove(string) error
}

type otpRepository struct {
	redisRepository RedisRepository
}

func NewOtpRepository(redisRepository RedisRepository) OtpRepository {
	return &otpRepository{
		redisRepository,
	}
}

func (repo *otpRepository) Save(o *entity.Otp) (*entity.Otp, error) {
	res, _ := repo.redisRepository.Get(o.Mobile)

	if res != "" {
		oldOtp := new(entity.Otp)
		if err := json.Unmarshal([]byte(res), oldOtp); err != nil {
			return nil, errors.New("marshal entity error")
		}

		return oldOtp, nil
	}

	bytes, _ := json.Marshal(o)
	if err := repo.redisRepository.Set(o.Mobile, bytes, time.Duration((o.ExpireAt-time.Now().UTC().Unix())*int64(time.Second))); err != nil {
		return nil, err
	}

	return nil, nil
}

func (repo *otpRepository) Remove(mobile string) error {
	if err := repo.redisRepository.Remove(mobile); err != nil {
		return err
	}

	return nil
}

func (repo *otpRepository) FindOne(phone string) (*entity.Otp, error) {
	res, err := repo.redisRepository.Get(phone)

	if res == "" {
		return nil, errors.New("otp not found")
	}

	if err != nil {
		return nil, err
	}

	otp := new(entity.Otp)
	if err := json.Unmarshal([]byte(res), otp); err != nil {
		return nil, errors.New("marshal entity error")
	}

	return otp, nil

}

func (repo *otpRepository) Attempt(mobile string) error {
	res, err := repo.redisRepository.Get(mobile)
	if err != nil {
		return err
	}
	otp := new(entity.Otp)

	if err := json.Unmarshal([]byte(res), otp); err != nil {
		return errors.New("marshal entity error")
	}

	if otp.Attempt == 0 {
		if err := repo.redisRepository.Remove(otp.Code); err != nil {
			return err
		}

		return errors.New("expired otp")
	}

	otp.Attempt = otp.Attempt - 1
	bytes, _ := json.Marshal(otp)
	if err := repo.redisRepository.Set(otp.Code, bytes, time.Duration((otp.ExpireAt-time.Now().UTC().Unix())*int64(time.Second))); err != nil {
		return err
	}

	return nil
}
