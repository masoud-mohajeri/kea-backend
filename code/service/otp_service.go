package service

import (
	"errors"

	"github.com/masoud-mohajeri/kea-backend/entity"
	"github.com/masoud-mohajeri/kea-backend/repository"
)

type OtpService interface {
	Request(string) (*entity.Otp, error)
}

type otpService struct {
	otpRepository repository.OtpRepository
	smsService    SmsService
}

func NewOtpService(smsService SmsService, otpRepository repository.OtpRepository) OtpService {
	return &otpService{
		otpRepository,
		smsService,
	}
}

func (otpServ *otpService) Request(phone string) (*entity.Otp, error) {
	otp := entity.NewOtp()
	otp.Mobile = phone
	dbOtp, saveError := otpServ.otpRepository.Save(otp)
	if saveError != nil {
		return nil, saveError
	}

	if dbOtp != nil {
		return dbOtp, nil
	}

	err := otpServ.smsService.sendSms(phone, otp.Code)
	if err != nil {
		otpServ.otpRepository.Remove(otp.Mobile)
		return nil, errors.New("error in sending sms")
	}

	return otp, nil
}
