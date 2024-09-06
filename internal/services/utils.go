package services

import (
	"fmt"
	"reflect"
)

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Missing field: %s", e.Field)
}

func ValidateStructData[T any](s T) *ValidationError {
	v := reflect.ValueOf(s)

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		// Get the field value
		fieldVal := v.Field(i)

		// Check if the field is zero (empty)
		if reflect.DeepEqual(fieldVal.Interface(), reflect.Zero(fieldVal.Type()).Interface()) {
			return &ValidationError{
				Field: t.Field(i).Name,
			}
		}
	}

	return nil
}
