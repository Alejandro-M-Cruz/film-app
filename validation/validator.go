package validation

import (
	"errors"
	"film-app/film"
	"film-app/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
)

var customValidation = map[string]validator.Func{
	"genre": func(fl validator.FieldLevel) bool {
		_, ok := film.Genres[fl.Field().String()]
		return ok
	},
}

type StructValidator struct {
	validate *validator.Validate
}

func NewStructValidator(validate *validator.Validate) *StructValidator {
	validate.RegisterTagNameFunc(utils.ExtractJsonFieldName)
	for key, val := range customValidation {
		err := validate.RegisterValidation(key, val)
		if err != nil {
			panic(err)
		}
	}

	return &StructValidator{validate}
}

func (v *StructValidator) Validate(i any) error {
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}

	var invalidValidationError *validator.InvalidValidationError
	if errors.As(err, &invalidValidationError) {
		return errors.New("validation failed")
	}

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		errorMessages := make(map[string][]string)

		for _, e := range validationErrors {
			errorMessages[e.Field()] = append(errorMessages[e.Field()], messageForFieldError(e))
		}

		return utils.NewErrors(messageForValidationErrors(validationErrors), errorMessages)
	}

	return err
}

func messageForValidationErrors(ve validator.ValidationErrors) string {
	switch len(ve) {
	case 0:
		return "Validation failed"
	case 1:
		return messageForFieldError(ve[0])
	case 2:
		return messageForFieldError(ve[0]) + " (and 1 other error)"
	default:
		return messageForFieldError(ve[0]) + fmt.Sprintf(" (and %d other errors)", len(ve)-1)
	}
}

func messageForFieldError(e validator.FieldError) string {
	var msg string

	switch e.Tag() {
	case "required":
		msg = fmt.Sprintf("%s is required", e.Field())
	case "min":
		msg = fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
	case "max":
		msg = fmt.Sprintf("%s must be at most %s characters long", e.Field(), e.Param())
	case "genre":
		msg = fmt.Sprintf("%s is not valid", e.Field())
	default:
		msg = fmt.Sprintf("%s is not valid", e.Field())
	}

	return utils.Capitalize(msg)
}
