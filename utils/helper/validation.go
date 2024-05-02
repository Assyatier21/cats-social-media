package helper

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/backend-magang/cats-social-media/models"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func UseCustomValidatorHandler(c *echo.Echo) {
	c.Validator = &CustomValidator{validator: validator.New()}
	c.HTTPErrorHandler = func(err error, c echo.Context) {
		var MessageValidation []string
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required", "required_with", "required_without", "required_unless":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is required",
						err.Field()))
				case "email":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is not valid email",
						err.Field()))
				case "gte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param()))
				case "lte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param()))
				case "numeric|eq=*":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be numeric or *",
						err.Field()))
				case "oneof":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value not one of %s",
						err.Field(), strings.ReplaceAll(err.Param(), " ", "/")))
				}
			}
			_ = WriteResponse(c, models.StandardResponseReq{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
				Data:    MessageValidation,
				Error:   nil,
			})
		}
	}

}

func IsValidSlug(s string) bool {
	regex, _ := regexp.Compile(`^[a-z0-9-]+$`)
	return regex.MatchString(s)
}

func BuildSortByQuery(sort_by string) string {
	if sort_by == "title" {
		sort_by = "title.keyword"
	} else if sort_by == "slug" {
		sort_by = "slug.keyword"
	} else if sort_by == "html_content" {
		sort_by = "html_content.keyword"
	} else {
		sort_by = "updated_at"
	}

	return sort_by
}

func BuildOrderByQuery(order_by string) bool {
	var order_by_bool bool

	order_by_bool = false
	if order_by == "asc" {
		order_by_bool = true
	} else if order_by == "desc" {
		order_by_bool = false
	}

	return order_by_bool
}

func IsNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
