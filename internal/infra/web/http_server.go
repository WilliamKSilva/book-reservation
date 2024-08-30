package web

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/WilliamKSilva/book-reservation/docs"
	"github.com/WilliamKSilva/book-reservation/internal/infra/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func registerUserRoutes(r *chi.Mux, userHandler handlers.IUserHandler) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		userHandler.Create(w, r)
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

// @host book-reservation.swagger.io
// @BasePath /v2
func StartListening(port int, dbConn *pgx.Conn) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	userHandler := handlers.NewUserHandler(dbConn)
	registerUserRoutes(r, userHandler)

	log.Printf("Listening at %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Println("Error trying to initialize web server")
		os.Exit(1)
	}
}
