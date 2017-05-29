package validation

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type maxLengthValidation struct {
	Validation
	length int
}

type minLengthValidation struct {
	Validation
	length int
}

type formatValidation struct {
	Validation
	pattern     *regexp.Regexp
	patternName string
}

func newMaxLengthValidation(options string, kind reflect.Kind) (Interface, error) {
	length, err := strconv.ParseInt(options, 10, 0)
	if err != nil {
		return nil, err
	}

	return &maxLengthValidation{
		length: int(length),
	}, nil
}

func (v *maxLengthValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
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

func newMinLengthValidation(options string, kind reflect.Kind) (Interface, error) {
	length, err := strconv.ParseInt(options, 10, 0)
	if err != nil {
		return nil, err
	}

	return &minLengthValidation{
		length: int(length),
	}, nil
}

func (v *minLengthValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
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

var emailRexep = regexp.MustCompile(`(?i)^[a-z0-9\._%+\-]+@[a-z0-9\.\-]+\.[a-z]{2,}$`)

func newFormatValidation(options string, kind reflect.Kind) (Interface, error) {
	if strings.ToLower(options) == "email" {
		return &formatValidation{
			pattern:     emailRexep,
			patternName: "email",
		}, nil
	} else if strings.Contains(options, "regexp:") {
		patternStr := options[strings.Index(options, ":")+1:]
		pattern, err := regexp.Compile(patternStr)

		if err != nil {
			return nil, &ValidationError{Key: "regexp:", Message: err.Error()}
		}

		return &formatValidation{
			pattern:     pattern,
			patternName: "regexp",
		}, nil
	}

	return nil, &ValidationError{Key: "format", Message: "Has no pattern " + options}
}

func (v *formatValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	strValue, ok := value.(string)
	if !ok {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "is not of type string. FormatValidation only accepts strings",
		}
	}

	if !v.pattern.MatchString(strValue) {
		return &ValidationError{
			Key:     v.FieldName(),
			Message: "does not match " + v.patternName + " format",
		}
	}

	return nil
}

func init() {
	AddValidation("max_length", newMaxLengthValidation)
	AddValidation("min_length", newMinLengthValidation)
	AddValidation("format", newFormatValidation)
}
