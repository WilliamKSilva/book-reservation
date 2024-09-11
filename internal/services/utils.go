package services

import (
	"reflect"

	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
)

func ValidateStructData[T any](s T) *services_errors.ValidationError {
	v := reflect.ValueOf(s)

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		// Get the field value
		fieldVal := v.Field(i)

		// Check if the field is zero (empty)
		if reflect.DeepEqual(fieldVal.Interface(), reflect.Zero(fieldVal.Type()).Interface()) {
			return &services_errors.ValidationError{
				Field: t.Field(i).Name,
			}
		}
	}

	return nil
}
