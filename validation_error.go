package validation

import "fmt"

type ValidationError struct {
	Key     string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Key + " " + e.Message
}

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return "Empty validation.ValidationErrors"
	}
	err := e[0].Error()
	if len(e) > 1 {
		err += fmt.Sprintf(" and %d other errors.", len(e))
	}
	return err
}
