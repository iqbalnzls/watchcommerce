package validator

import "gopkg.in/go-playground/validator.v9"

type DataValidator struct {
	ValidatorData *validator.Validate
}

func NewValidator() *DataValidator {
	v := &DataValidator{ValidatorData: validator.New()}

	return v
}

func (v *DataValidator) Validate(i interface{}) (err error) {
	if err = v.ValidatorData.Struct(i); err != nil {
		return
	}

	return
}
