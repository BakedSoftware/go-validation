package validation

type ValidationError struct {
	Key     string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Key + " " + e.Message
}
