package domain

import (
	"github.com/Binit-Dhakal/Foody/internal/validator"
)

type RegisterUserRequest struct {
	Name            string               `json:"name"`
	Email           string               `json:"email"`
	Password        string               `json:"password"`
	ConfirmPassword string               `json:"confirmPassword"`
	Validator       *validator.Validator `json:"-"`
}

func (r *RegisterUserRequest) Validate() {
	r.Validator = &validator.Validator{}
	v := r.Validator

	validateName(v, r.Name, "name")
	validateEmail(v, r.Email)
	validatePassword(v, r.Password)

	v.CheckField(r.Password == r.ConfirmPassword, "confirmPassword", "passwords do not match")
}

type RegisterResturantRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	License         string `json:"license"`
	ResturantName   string `json:"resturantName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`

	Validator *validator.Validator `json:"-"`
}

func (r *RegisterResturantRequest) Validate() {
	r.Validator = &validator.Validator{}
	v := r.Validator

	validateName(v, r.Name, "name")
	validateEmail(v, r.Email)
	validatePassword(v, r.Password)
	validateName(v, r.ResturantName, "resturantName")
}

type LoginUserRequest struct {
	Email     string               `json:"email"`
	Password  string               `json:"password"`
	Validator *validator.Validator `json:"-"`
}

func (r *LoginUserRequest) Validate() {
	r.Validator = &validator.Validator{}
	v := r.Validator
	validateEmail(v, r.Email)
	validatePassword(v, r.Password)
}

type SessionDataResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
