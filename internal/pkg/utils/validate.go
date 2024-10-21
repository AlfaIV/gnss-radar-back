package utils

import (
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateSignup(input *model.SignupInput) error {
	return validation.ValidateStruct(input,
		validation.Field(&input.Login, validation.Required),
		validation.Field(&input.Password, validation.Required),
	)
}

func ValidateSignin(input *model.SigninInput) error {
	return validation.ValidateStruct(input,
		validation.Field(&input.Login, validation.Required),
		validation.Field(&input.Password, validation.Required),
	)
}
