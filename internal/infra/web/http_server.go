package web

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WilliamKSilva/book-reservation/internal/infra/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func registerUserRoutes(r *chi.Mux, userHandler handlers.IUserHandler) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		userHandler.Create(w, r)
	})
}

func StartListening(port int, dbConn *pgx.Conn) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	userHandler := handlers.NewUserHandler(dbConn)

	registerUserRoutes(r, userHandler)

	log.Printf("Listening at %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Println("Error trying to initialize web server")
		os.Exit(1)
	}
}
