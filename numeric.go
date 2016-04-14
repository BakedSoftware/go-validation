package validation

import (
	"reflect"
	"strconv"
)

type intValueValidation struct {
	Validation
	value int64
	less  bool
}

func (m *intValueValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	var compareValue int64
	switch value := value.(type) {
	case int:
		compareValue = int64(value)
	case int8:
		compareValue = int64(value)
	case int16:
		compareValue = int64(value)
	case int32:
		compareValue = int64(value)
	case int64:
		compareValue = int64(value)
	default:
		return &ValidationError{
			Key:     m.FieldName(),
			Message: "is not convertible to type int64",
		}
	}

	if m.less {
		if compareValue < m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be greater than or equal to " + strconv.FormatInt(m.value, 10),
			}
		}
	} else {
		if compareValue > m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be less than or equal to " + strconv.FormatInt(m.value, 10),
			}
		}
	}

	return nil
}

type uintValueValidation struct {
	Validation
	value uint64
	less  bool
}

func (m *uintValueValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	var compareValue uint64
	switch value := value.(type) {
	case uint:
		compareValue = uint64(value)
	case uint8:
		compareValue = uint64(value)
	case uint16:
		compareValue = uint64(value)
	case uint32:
		compareValue = uint64(value)
	case uint64:
		compareValue = uint64(value)
	default:
		return &ValidationError{
			Key:     m.FieldName(),
			Message: "is not convertible to type uint64",
		}
	}

	if m.less {
		if compareValue < m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be greater than or equal to " + strconv.FormatUint(m.value, 10),
			}
		}
	} else {
		if compareValue > m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be less than or equal to " + strconv.FormatUint(m.value, 10),
			}
		}
	}

	return nil
}

type floatValueValidation struct {
	Validation
	value float64
	less  bool
}

func (m *floatValueValidation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	var compareValue float64
	switch value := value.(type) {
	case float32:
		compareValue = float64(value)
	case float64:
		compareValue = float64(value)
	default:
		return &ValidationError{
			Key:     m.FieldName(),
			Message: "is not convertible to type float64",
		}
	}

	if m.less {
		if compareValue < m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be greater than or equal to " + strconv.FormatFloat(m.value, 'E', -1, 64),
			}
		}
	} else {
		if compareValue > m.value {
			return &ValidationError{
				Key:     m.FieldName(),
				Message: "must be less than or equal to " + strconv.FormatFloat(m.value, 'E', -1, 64),
			}
		}
	}

	return nil
}

func newMinValueValidation(options string, kind reflect.Kind) (Interface, error) {
	switch kind {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value, err := strconv.ParseInt(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &intValueValidation{
			value: value,
			less:  true,
		}, nil
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value, err := strconv.ParseUint(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &uintValueValidation{
			value: value,
			less:  true,
		}, nil
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value, err := strconv.ParseFloat(options, 64)
		if err != nil {
			return nil, err
		}
		return &floatValueValidation{
			value: value,
			less:  true,
		}, nil
	default:
		return nil, &ValidationError{
			Key:     "Invalid Validation",
			Message: "Field is not of numeric type. Min validation only accepts numeric types",
		}
	}
}

func newMaxValueValidation(options string, kind reflect.Kind) (Interface, error) {
	switch kind {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		value, err := strconv.ParseInt(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &intValueValidation{
			value: value,
			less:  false,
		}, nil
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		value, err := strconv.ParseUint(options, 10, 0)
		if err != nil {
			return nil, err
		}
		return &uintValueValidation{
			value: value,
			less:  false,
		}, nil
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		value, err := strconv.ParseFloat(options, 64)
		if err != nil {
			return nil, err
		}
		return &floatValueValidation{
			value: value,
			less:  false,
		}, nil
	default:
		return nil, &ValidationError{
			Key:     "Invalid Validation",
			Message: "Field is not of numeric type. Max validation only accepts numeric types",
		}
	}
}

func init() {
	AddValidation("min", newMinValueValidation)
	AddValidation("max", newMaxValueValidation)
}
