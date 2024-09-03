package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/WilliamKSilva/book-reservation/internal/app/user"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/WilliamKSilva/book-reservation/internal/infra/uuid"
	"github.com/WilliamKSilva/book-reservation/internal/infra/web/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewUserHandler(dbConn *pgxpool.Pool) *UserHandler {
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

type IUserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	UserService user.IUserService
}

// CreateUser godoc
// @Summary      Create an User
// @Description
// @Tags         users
// @Accept       json
// @Produce      json
// @Param user body user.CreateUserRequestDTO true "User details"
// @Success      200  {object}  user.CreateUserResponseDTO
// @Failure      500  {object}  utils.HttpError
// @Router       /users [post]
func (userHandler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var httpError utils.HttpError
	b, err := io.ReadAll(r.Body)
	if err != nil {
		httpError.Code = http.StatusUnprocessableEntity
		httpError.Message = "Error trying to read body"

		utils.HttpResponse(w, httpError.Message, httpError.Code)
		return
	}

	var userRequest user.CreateUserRequestDTO
	err = json.Unmarshal(b, &userRequest)
	if err != nil {
		httpError.Code = http.StatusUnprocessableEntity
		httpError.Message = "Error trying to decode request body"
		utils.HttpResponse(w, httpError.Message, httpError.Code)
		return
	}

	log.Println(userRequest)
}
