package dto

type PasswordLoginDto struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type OtpConfirmationDto struct {
	Mobile string `json:"mobile"`
	Otp    string `json:"otp"`
}

type ChangeMobileDto struct {
	Mobile    string `json:"mobile"`
	NewMobile string `json:"new_mobile"`
	Otp       string `json:"otp"`
}

type RefreshTokenDto struct {
	RefreshToken string `json:"refresh_token"`
}
