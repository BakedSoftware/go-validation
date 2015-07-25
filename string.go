package validation

import (
	"strconv"
)

type maxLengthValidation struct {
	Validation
	length int
}

type minLengthValidation struct {
	Validation
	length int
}

func newMaxLengthValidation(options string) (Interface, error) {
	length, err := strconv.ParseInt(options, 10, 0)
	if err != nil {
		return nil, err
	}

	return &maxLengthValidation{
		length: int(length),
	}, nil
}

func (v *maxLengthValidation) Validate(value interface{}) *ValidationError {
	strValue, ok := value.(string)
	if !ok {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "is not of type string. MaxLengthValidation only accepts strings",
		}
	}

	if len(strValue) > v.length {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "must be no more than " + strconv.Itoa(v.length) + " characters",
		}
	}

	return nil
}

func newMinLengthValidation(options string) (Interface, error) {
	length, err := strconv.ParseInt(options, 10, 0)
	if err != nil {
		return nil, err
	}

	return &minLengthValidation{
		length: int(length),
	}, nil
}

func (v *minLengthValidation) Validate(value interface{}) *ValidationError {
	strValue, ok := value.(string)
	if !ok {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "is not of type string. MinLengthValidation only accepts strings",
		}
	}

	if len(strValue) < v.length {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "must be at least " + strconv.Itoa(v.length) + " characters",
		}
	}

	return nil
}

func init() {
	AddValidation("max_length", newMaxLengthValidation)
	AddValidation("min_length", newMinLengthValidation)
}
