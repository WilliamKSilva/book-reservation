package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func httpResponse(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(response))
}

func validMethod(requestMethod string, expectedMethod string) bool {
	return requestMethod == expectedMethod
}

func registerRoutes() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		validMethod := validMethod(r.Method, "POST")
		if !validMethod {
			response := fmt.Sprintf("Method %s not allowed for /register", r.Method)
			httpResponse(w, response)
			return
		}
	})
}

func StartListening(port int) {
	registerRoutes()

	log.Printf("Listening at %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println("Error trying to initialize web server")
		os.Exit(1)
	}
}
