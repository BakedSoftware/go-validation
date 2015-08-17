// Package validation provides validations for struct fields based on a
// validation tag
package validation

import (
	"log"
	"reflect"
	"strings"
)

// Interface specifies the necessary methods a validation must
// implement to be compatible with this package
type Interface interface {
	SetFieldIndex(index int)
	FieldIndex() int
	SetFieldName(name string)
	FieldName() string
	Validate(value interface{}, obj interface{}) *ValidationError
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
func (v *Validation) Validate(value interface{}, obj interface{}) *ValidationError {
	return &ValidationError{
		Key:     v.fieldName,
		Message: "Validation not implemented",
	}
}

var validationMap map[reflect.Type][]Interface
var validationNameToBuilder map[string]func(string) (Interface, error)

func prepareMap() {
	if validationNameToBuilder == nil {
		validationNameToBuilder = make(map[string]func(string) (Interface, error))
	}
}

func init() {
	validationMap = make(map[reflect.Type][]Interface, 10)
}

// AddValidation registers the validation specified by key to the known
// validations. If more than one validation registers with the same key, the
// last one will become the validation for that key
func AddValidation(key string, fn func(string) (Interface, error)) {
	prepareMap()
	validationNameToBuilder[key] = fn
}

// IsValid determines if an object is valid based on its validation tags
func IsValid(object interface{}) (bool, []ValidationError) {
	objectValue := reflect.ValueOf(object)
	objectType := reflect.TypeOf(object)
	validations := validationMap[objectType]
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
					if builder := validationNameToBuilder[comps[0]]; builder != nil {
						validation, err = builder(comps[1])
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
		validationMap[objectType] = validations
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
