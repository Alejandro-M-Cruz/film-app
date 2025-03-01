package utils

type Errors struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors,omitempty"`
}

func NewErrors(message string, errors map[string][]string) Errors {
	return Errors{Message: message, Errors: errors}
}

func (e Errors) Error() string {
	return e.Message
}
