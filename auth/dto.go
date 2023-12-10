package auth

import "github.com/go-playground/validator/v10"

type LoginDto struct {
	Username string `json:"username" validate:"required,min=5,max=40"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

func (dto LoginDto) Validate() error {
	validate := validator.New()
	return validate.Struct(dto)
}
