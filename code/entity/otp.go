package entity

import (
	"time"
)

type Otp struct {
	Code     string
	ExpireAt int64
	Attempt  int
	Mobile   string
}

func NewOtp() *Otp {
	return &Otp{
		Code:     "1111",
		ExpireAt: time.Now().UTC().Add(1 * time.Minute).Unix(),
		Attempt:  3,
	}
}

func (otp Otp) IsExpired() bool {
	if otp.Attempt <= 0 {
		return true
	}

	return time.Unix(otp.ExpireAt, 0).Before(time.Now().UTC())
}
