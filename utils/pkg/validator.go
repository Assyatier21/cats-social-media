package pkg

import (
	"fmt"
	"strings"

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
	v := validator.New()

	v.RegisterValidation("validateRaces", customRaceEnum)

	return v
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

func customRaceEnum(fl validator.FieldLevel) bool {
	allowedValues := []string{
		"Persian",
		"Maine Coon",
		"Siamese",
		"Ragdoll",
		"Bengal",
		"Sphynx",
		"British Shorthair",
		"Abyssinian",
		"Scottish Fold",
		"Birman",
	}

	value := fl.Field().String()
	for _, v := range allowedValues {
		if strings.EqualFold(value, v) {
			return true
		}
	}
	return false
}
