package dto

type RequestOtpDto struct {
	Mobile string `json:"mobile"`
}

func (otp RequestOtpDto) Validate() error {
	// validate if it is really a mobile number

	return nil
}

type OTPDetails struct {
	ExpiredAt int64 `json:"expired_at"`
	IsNew     bool  `json:"is_new"`
}
