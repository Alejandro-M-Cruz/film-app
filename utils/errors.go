package utils

type Errors struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

func NewError(message string) Errors {
	return Errors{Message: message}
}

func NewErrors(message string, errors map[string][]string) Errors {
	return Errors{Message: message, Errors: errors}
}
