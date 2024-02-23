package service

import (
	"errors"

	"github.com/masoud-mohajeri/kea-backend/entity"
	"github.com/masoud-mohajeri/kea-backend/repository"
)

type OtpService interface {
	Request(mobile string) (*entity.Otp, error)
	Validate(code string, mobile string) error
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

func (otpServ *otpService) Request(mobile string) (*entity.Otp, error) {
	otp := entity.NewOtp()
	otp.Mobile = mobile
	dbOtp, saveError := otpServ.otpRepository.Save(otp)
	if saveError != nil {
		return nil, saveError
	}

	if dbOtp != nil {
		return dbOtp, nil
	}

	err := otpServ.smsService.sendSms(mobile, otp.Code)
	if err != nil {
		otpServ.otpRepository.Remove(otp.Mobile)
		return nil, errors.New("error in sending sms")
	}

	return otp, nil
}

func (ots *otpService) Validate(code string, mobile string) error {

	otp, err := ots.otpRepository.FindOne(mobile)

	if err != nil {
		return err
	}

	if otp == nil {
		return errors.New("go get otp")
	}

	if otp.Code != code {
		ots.otpRepository.Attempt(mobile)
		return errors.New("otp did not match")
	}
	ots.otpRepository.Remove(otp.Mobile)

	return nil
}
