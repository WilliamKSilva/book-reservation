package services_errors

type WrongPasswordError struct {
	Message string
}

func (e *WrongPasswordError) Error() string {
	return e.Message
}
