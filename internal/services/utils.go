package services

import (
	"log"
	"reflect"

	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
)

func LogUnexpectedError(service string, method string, message string) {
	log.Printf("ERROR: unexpeted error at service %s on method %s\n%s", service, method, message)
}

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
