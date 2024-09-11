package services_errors

import "fmt"

type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Missing field: %s", e.Field)
}
