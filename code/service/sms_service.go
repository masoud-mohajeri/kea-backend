package service

import "fmt"

type SmsService interface {
	sendSms(phone string, text string) error
}

type smsService struct{}

func NewSmsService() SmsService {
	return &smsService{}
}

func (ss *smsService) sendSms(phone string, text string) error {
	fmt.Printf("send a sms to %s: %s \n", phone, text)
	return nil
}
