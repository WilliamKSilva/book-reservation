package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/WilliamKSilva/book-reservation/docs"
	"github.com/WilliamKSilva/book-reservation/internal/infra/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func registerUserRoutes(r *chi.Mux, authHandler handlers.IAuthHandler) {
	r.Post("/auth/register", func(w http.ResponseWriter, r *http.Request) {
		authHandler.Register(w, r)
	})
}

// @title Book Reservation API
// @version 1.0
// @description This is the API for the online Book Reservation service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
func StartListening(ctx context.Context, port int, dbPool *pgxpool.Pool) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(corsMiddleware)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	authHandler := handlers.NewAuthHandler(ctx, dbPool)
	registerUserRoutes(r, authHandler)

	log.Printf("Listening at %d", port)
	log.Printf("API documentation at %s", "http://localhost:8080/swagger/")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Println("Error trying to initialize web server")
		os.Exit(1)
	}
}
