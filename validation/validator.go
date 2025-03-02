package validation

import (
	"errors"
	"film-app/film"
	"film-app/utils"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var customValidation = map[string]validator.Func{
	"genre": func(fl validator.FieldLevel) bool {
		_, ok := film.Genres[fl.Field().String()]
		return ok
	},
	"starts_with_alpha": func(fl validator.FieldLevel) bool {
		value := fl.Field().String()

		if len(value) == 0 {
			return false
		}

		return value[0] >= 'A' && value[0] <= 'Z' || value[0] >= 'a' && value[0] <= 'z'
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

		return NewErrors(messageForValidationErrors(validationErrors), errorMessages)
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
		return messageForFieldError(ve[0]) + " (and 1 more error)"
	default:
		return messageForFieldError(ve[0]) + fmt.Sprintf(" (and %d more errors)", len(ve)-1)
	}
}

func messageForFieldError(e validator.FieldError) string {
	var msg string

	switch e.Tag() {
	case "required":
		msg = fmt.Sprintf("%s is required", e.Field())
	case "min":
		msg = messageForMinError(e)
	case "max":
		msg = messageForMaxError(e)
	case "genre":
		msg = fmt.Sprintf("%s is not valid", e.Field())
	case "alphanum":
		msg = fmt.Sprintf("%s must be alphanumeric", e.Field())
	case "starts_with_alpha":
		msg = fmt.Sprintf("%s must start with a letter", e.Field())
	default:
		msg = fmt.Sprintf("%s is not valid", e.Field())
	}

	return utils.Capitalize(msg)
}

func messageForMinError(e validator.FieldError) string {
	return messageForMinOrMaxError(e, "least")
}

func messageForMaxError(e validator.FieldError) string {
	return messageForMinOrMaxError(e, "most")
}

func messageForMinOrMaxError(e validator.FieldError, superlative string) string {
	switch e.Kind() {
	case reflect.String:
		characters := pluralize("character", e.Param())

		return fmt.Sprintf("%s must be at %s %s %s long", e.Field(), superlative, e.Param(), characters)
	case reflect.Slice, reflect.Array, reflect.Map:
		items := pluralize("item", e.Param())

		return fmt.Sprintf("%s must have at %s %s %s", e.Field(), superlative, e.Param(), items)
	default:
		return fmt.Sprintf("%s must be at %s %s", e.Field(), superlative, e.Param())
	}
}

func pluralize(s string, count string) string {
	if count == "1" {
		return s
	}

	return s + "s"
}
