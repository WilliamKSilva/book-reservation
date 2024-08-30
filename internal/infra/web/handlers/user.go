package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/WilliamKSilva/book-reservation/internal/app/user"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/WilliamKSilva/book-reservation/internal/infra/uuid"
	"github.com/jackc/pgx/v5"
)

func NewUserHandler(dbConn *pgx.Conn) *UserHandler {
	userPostgresRepository := db.PostgresUserRepository{Conn: dbConn}
	googleUuidGenerator := uuid.GoogleUUIDGenerator{}
	userService := user.UserService{
		UserRepository: &userPostgresRepository,
		UuidGenerator:  &googleUuidGenerator,
	}

	return &UserHandler{
		UserService: &userService,
	}
}

func HttpResponse(w http.ResponseWriter, response string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(response))
}

type IUserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	UserService user.IUserService
}

func (userHandler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		HttpResponse(w, "Error trying to read body")
		return
	}

	var userRequest user.CreateRequestDTO
	err = json.Unmarshal(b, &userRequest)
	if err != nil {
		log.Printf("ERROR: %s", err)
		HttpResponse(w, "Error trying to decoded request body")
		return
	}

	log.Println(userRequest)
}
