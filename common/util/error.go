package util

import (
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

type Error struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned object
func NewValidatorError(err error) Error {
	res := Error{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		if v.Param() != "" {
			res.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			res.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}

	}
	return res
}

// Warp the error info in a object
func NewError(key string, err error) Error {
	res := Error{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
