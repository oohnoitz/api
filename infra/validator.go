package infra

import (
	"github.com/asaskevich/govalidator"
	"github.com/tadoku/api/usecases"
)

// NewValidator validates structs and possibly other stuff in the future
func NewValidator() usecases.Validator {
	return &validator{}
}

type validator struct{}

func (v *validator) ValidateStruct(s interface{}) (bool, error) {
	return govalidator.ValidateStruct(s)
}