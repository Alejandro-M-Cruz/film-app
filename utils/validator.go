package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func extractJsonFieldName(field reflect.StructField) string {
	name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

type StructValidator struct {
	validate *validator.Validate
}

func NewStructValidator(validate *validator.Validate) *StructValidator {
	validate.RegisterTagNameFunc(extractJsonFieldName)

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
			errorMessages[e.Field()] = append(errorMessages[e.Field()], v.messageForFieldError(e))
		}

		return NewErrors(v.messageForValidationErrors(validationErrors), errorMessages)
	}

	return err
}

func (v *StructValidator) messageForValidationErrors(ve validator.ValidationErrors) string {
	switch len(ve) {
	case 0:
		return "Validation failed"
	case 1:
		return v.messageForFieldError(ve[0])
	case 2:
		return v.messageForFieldError(ve[0]) + " (and 1 other error)"
	default:
		return v.messageForFieldError(ve[0]) + fmt.Sprintf(" (and %d other errors)", len(ve)-1)
	}
}

func (v *StructValidator) messageForFieldError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s is not valid", e.Field())
	}
}
