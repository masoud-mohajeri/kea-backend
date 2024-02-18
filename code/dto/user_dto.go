package dto

type UserInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type PasswordLoginDto struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
