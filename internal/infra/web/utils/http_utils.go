package utils

import "net/http"

type HttpError struct {
	Message string
	Code    int
}

const ErrorReadingBody = "Error reading request body"
const ErrorDecodingJson = "Error trying to decode body JSON"
const ErrorInternal = "Internal server error"

func HttpResponse[T ~string | ~[]byte](w http.ResponseWriter, response T, code int) {
	w.WriteHeader(code)
	w.Write([]byte(response))
}

func UnprocessableEntityError(message string) HttpError {
	return HttpError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func InternalServerError(message string) HttpError {
	return HttpError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}
