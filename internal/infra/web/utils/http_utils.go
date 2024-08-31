package utils

import "net/http"

type HttpError struct {
	Message string
	Code    int
}

func HttpResponse(w http.ResponseWriter, response string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(response))
}
