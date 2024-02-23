package dto

type PasswordLoginDto struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type OtpLoginDto struct {
	Mobile string `json:"mobile"`
	Otp    string `json:"otp"`
}
