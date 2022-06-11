package validator

import "gopkg.in/go-playground/validator.v9"

type DataValidator struct {
	validator *validator.Validate
}

func NewValidator() *DataValidator {
	return &DataValidator{
		validator: validator.New(),
	}
}

func (v *DataValidator) Validate(i interface{}) (err error) {
	if err = v.validator.Struct(i); err != nil {
		return
	}

	return
}
