package constants

import (
	"errors"
	"strings"
)

type UserRole int

const (
	ADMIN UserRole = iota
	USER
)

func GetRole(role string) (UserRole, error) {
	if strings.ToUpper(role) == "ADMIN" {
		return ADMIN, nil
	} else if strings.ToUpper(role) == "USER" {
		return USER, nil
	} else {
		return -1, errors.New("invalid input")
	}
}
