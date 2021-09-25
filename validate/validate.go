package validate

import (
	ent "snakealive/m/entities"
)

func Validate(user ent.User) bool {
	if user.Email == "" {
		return false
	}
	if user.Password == "" {
		return false
	}
	return true
}
