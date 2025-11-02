package domain

import "github.com/Binit-Dhakal/Foody/internal/validator"

func validateName(v *validator.Validator, value string, name string) {
	v.CheckField(len(value) != 0, name, name+" cannot be empty")
	v.CheckField(len(value) < 500, name, name+" cannot be more than 500 bytes")
}

func validateUserName(v *validator.Validator, username string) {
	v.CheckField(len(username) != 0, "username", "username cannot be empty")
	v.CheckField(len(username) < 500, "username", "username cannot be more than 500 bytes")
}

func validatePassword(v *validator.Validator, password string) {
	v.CheckField(len(password) >= 8, "password", "password must be at least 8 characters")
	v.CheckField(len(password) <= 72, "password", "password cannot be greater than 72 characters")
}

func validateEmail(v *validator.Validator, email string) {
	v.CheckField(len(email) != 0, "email", "email cannot be empty")
	v.CheckField(validator.IsEmail(email), "email", "email is not valid")
}
