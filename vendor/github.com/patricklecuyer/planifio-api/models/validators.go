package models

import (
	validator "gopkg.in/validator.v2"
)

//InitValidators Initialize Validatprs
func InitValidators() {
	validator.SetValidationFunc("basic", basicValidator)
}

//BasicValidator for input validation
func basicValidator(v interface{}, param string) error {
	return nil
}
