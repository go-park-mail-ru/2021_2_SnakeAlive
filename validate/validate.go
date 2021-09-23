package validate

import (
	"snakealive/m/auth"
)

func Validate(user auth.User) bool {
	if user.Email == "" {
		return false
	}
	if user.Password == "" {
		return false
	}
	return true
}
