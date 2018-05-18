// Package validation provides validations for struct fields based on a
// validation tag
package validation

import (
	"log"
	"reflect"
	"strings"
	"sync"
)

// Interface specifies the necessary methods a validation must
// implement to be compatible with this package
type Interface interface {
	SetFieldIndex(index int)
	FieldIndex() int
	SetFieldName(name string)
	FieldName() string
	Validate(value interface{}, obj reflect.Value) *ValidationError
}

// Validation is an implementation of a Interface and can be used to
// provide basic functionality to a new validation type through an anonymous
// field
type Validation struct {
	Name       string
	fieldIndex int
	fieldName  string
	options    string
}

// SetFieldIndex stores the index of the field the validation was applied to
func (v *Validation) SetFieldIndex(index int) {
	v.fieldIndex = index
}

// FieldIndex retrieves the index of the field the validation was applied to
func (v *Validation) FieldIndex() int {
	return v.fieldIndex
}

// SetFieldName stores the name of the field the validation was applied to
func (v *Validation) SetFieldName(name string) {
	v.fieldName = name
}

// FieldName retrieves the name of the field the validation was applied to
func (v *Validation) FieldName() string {
	return v.fieldName
}

// Validate determines if the value is valid. Nil is returned if it is valid
func (v *Validation) Validate(value interface{}, obj reflect.Value) *ValidationError {
	return &ValidationError{
		Key:     v.fieldName,
		Message: "Validation not implemented",
	}
}

// DefaultValidationMap is the default validation
// map used to tell if a struct is valid.
var DefaultValidationMap = Map{}

// Map is an atomic validation map
// when two Set happen at the same time,
// latest that started wins.
type Map struct {
	validator               sync.Map // map[reflect.Type][]Interface
	validationNameToBuilder sync.Map // map[string]func(string, reflect.Kind) (Interface, error)
}

func (vm *Map) get(k reflect.Type) []Interface {
	v, ok := vm.validator.Load(k)
	if !ok {
		return []Interface{}
	}
	return v.([]Interface)
}
func (vm *Map) set(k reflect.Type, v []Interface) {
	vm.validator.Store(k, v)
}

// AddValidation registers the validation specified by key to the known
// validations. If more than one validation registers with the same key, the
// last one will become the validation for that key
// using DefaultValidationMap.
func AddValidation(key string, fn func(string, reflect.Kind) (Interface, error)) {
	DefaultValidationMap.AddValidation(key, fn)
}

// AddValidation registers the validation specified by key to the known
// validations. If more than one validation registers with the same key, the
// last one will become the validation for that key.
func (vm *Map) AddValidation(key string, fn func(string, reflect.Kind) (Interface, error)) {
	vm.validationNameToBuilder.Store(key, fn)
}

// IsValid determines if an object is valid based on its validation tags
// using DefaultValidationMap.
func IsValid(object interface{}) (bool, []ValidationError) {
	return DefaultValidationMap.IsValid(object)
}

// IsValid determines if an object is valid based on its validation tags.
func (vm *Map) IsValid(object interface{}) (bool, []ValidationError) {
	objectValue := reflect.ValueOf(object)
	objectType := reflect.TypeOf(object)
	validations := vm.get(objectType)
	if objectValue.Kind() == reflect.Ptr && !objectValue.IsNil() {
		return IsValid(objectValue.Elem().Interface())
	}
	if len(validations) == 0 {
		var err error
		for i := objectType.NumField() - 1; i >= 0; i-- {
			field := objectType.Field(i)
			validationTag := field.Tag.Get("validation")
			if len(validationTag) > 0 {
				validationComps := strings.Split(validationTag, " ")
				for _, v := range validationComps {
					comps := strings.Split(v, "=")
					if len(comps) != 2 {
						log.Fatalln("Invalid Validation Specification:", objectType.Name(), field.Name, v)
					}
					var validation Interface
					if builder, ok := vm.validationNameToBuilder.Load(comps[0]); ok && builder != nil {
						fn := builder.(func(string, reflect.Kind) (Interface, error))
						validation, err = fn(comps[1], field.Type.Kind())
					} else {
						log.Fatalln("Unknown validation named", comps[0])
					}
					if err != nil {
						log.Fatalln("Error Creating Validation", objectType.Name(), field.Name, v, err)
					}
					validation.SetFieldName(field.Name)
					validation.SetFieldIndex(i)
					validations = append(validations, validation)
				}
			}
		}
		vm.set(objectType, validations)
	}

	var errors []ValidationError
	for _, validation := range validations {
		field := objectValue.Field(validation.FieldIndex())
		value := field.Interface()
		if err := validation.Validate(value, objectValue); err != nil {
			errors = append(errors, *err)
		}
	}

	return len(errors) == 0, errors
}
