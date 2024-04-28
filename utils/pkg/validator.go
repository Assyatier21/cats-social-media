package pkg

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func SetupValidator() *validator.Validate {
	return validator.New()
}

func BindValidate(c echo.Context, req interface{}) (err error) {
	if err = c.Bind(req); err != nil {
		err = fmt.Errorf("[Utils][Pkg][Validator] failed to bind request, err: %s", err.Error())
		return
	}

	if err = c.Validate(req); err != nil {
		err = fmt.Errorf("[Utils][Pkg][Validator] failed to validate request, err: %s", err.Error())
		return
	}

	return
}
