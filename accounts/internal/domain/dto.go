package domain

import "github.com/Binit-Dhakal/Foody/internal/validator"

type RegisterUserRequest struct {
	Name            string              `json:"name"`
	Username        string              `json:"username"`
	Email           string              `json:"email"`
	Password        string              `json:"password"`
	ConfirmPassword string              `json:"confirmPassword"`
	Validator       validator.Validator `json:"-"`
}

func (r *RegisterUserRequest) Validate() {
	r.Validator.CheckField(len(r.Name) != 0, "name", "name cannot be empty")
	r.Validator.CheckField(len(r.Name) < 500, "name", "name cannot be more than 500 bytes")

	r.Validator.CheckField(len(r.Username) != 0, "username", "username cannot be empty")
	r.Validator.CheckField(len(r.Username) < 500, "username", "username cannot be more than 500 bytes")

	r.Validator.CheckField(len(r.Password) >= 8, "password", "password must be at least 8 characters")
	r.Validator.CheckField(len(r.Password) <= 72, "password", "password cannot be greater than 72 characters")

	r.Validator.CheckField(r.Password == r.ConfirmPassword, "confirmPassword", "passwords do not match")

	r.Validator.CheckField(len(r.Email) != 0, "email", "email cannot be empty")
	r.Validator.CheckField(validator.IsEmail(r.Email), "email", "email is not valid")
}

type RegisterResturantRequest struct {
	Name            string              `json:"name"`
	Username        string              `json:"username"`
	Email           string              `json:"email"`
	License         string              `json:"license"`
	ResturantName   string              `json:"resturantName"`
	Password        string              `json:"password"`
	ConfirmPassword string              `json:"confirmPassword"`
	Validator       validator.Validator `json:"-"`
}

func (r *RegisterResturantRequest) Validate() {
	r.Validator.CheckField(len(r.Name) != 0, "name", "name cannot be empty")
	r.Validator.CheckField(len(r.Name) < 500, "name", "name cannot be more than 500 bytes")

	r.Validator.CheckField(len(r.Username) != 0, "username", "username cannot be empty")
	r.Validator.CheckField(len(r.Username) < 500, "username", "username cannot be more than 500 bytes")

	r.Validator.CheckField(len(r.Password) >= 8, "password", "password must be at least 8 characters")
	r.Validator.CheckField(len(r.Password) <= 72, "password", "password cannot be greater than 72 characters")
	r.Validator.CheckField(r.Password == r.ConfirmPassword, "confirmPassword", "passwords do not match")

	r.Validator.CheckField(len(r.Email) != 0, "email", "email cannot be empty")
	r.Validator.CheckField(validator.IsEmail(r.Email), "email", "email is not valid")

	r.Validator.CheckField(len(r.ResturantName) != 0, "resturantName", "resturant name cannot be empty")
	r.Validator.CheckField(len(r.ResturantName) < 500, "resturantName", "resturant name cannot be more than 500 bytes")
}
