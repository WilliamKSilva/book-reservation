package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/WilliamKSilva/book-reservation/internal/app/auth"
	"github.com/WilliamKSilva/book-reservation/internal/app/user"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/WilliamKSilva/book-reservation/internal/infra/jwt"
	"github.com/WilliamKSilva/book-reservation/internal/infra/uuid"
	"github.com/WilliamKSilva/book-reservation/internal/infra/web/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewAuthHandler(ctx context.Context, dbPool *pgxpool.Pool) *AuthHandler {
	userRepository := db.PostgresUserRepository{Conn: dbPool, Ctx: ctx}
	uuidGenerator := uuid.GoogleUUIDGenerator{}
	authService := auth.AuthService{
		JwtService:  &jwt.GolangJwt{},
		UserService: user.NewUserService(&userRepository, &uuidGenerator),
	}
	return &AuthHandler{AuthService: &authService}
}

type IAuthHandler interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	AuthService auth.IAuthService
}

// CreateUser godoc
// @Summary      Register a new User
// @Description
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param user body auth.RegisterRequestDTO true "User details"
// @Success      200  {object}  auth.RegisterResponseDTO
// @Failure      500  {object}  utils.HttpError
// @Router       /auth/register [post]
func (authHandler *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		httpError := utils.UnprocessableEntityError(utils.ErrorReadingBody)
		utils.HttpResponse(w, httpError.Message, httpError.Code)
		return
	}

	var registerRequestDTO auth.RegisterRequestDTO
	err = json.Unmarshal(b, &registerRequestDTO)
	if err != nil {
		httpError := utils.UnprocessableEntityError(utils.ErrorDecodingJson)
		utils.HttpResponse(w, httpError.Message, httpError.Code)
		return
	}

	registerResponseDTO, err := authHandler.AuthService.Register(registerRequestDTO)
	if err != nil {
		log.Println(err)
		httpError := utils.InternalServerError(utils.ErrorInternal)
		utils.HttpResponse(w, httpError.Message, httpError.Code)
		return
	}

	b, err = json.Marshal(registerResponseDTO)
	if err != nil {
		httpError := utils.InternalServerError(utils.ErrorInternal)
		utils.HttpResponse(w, httpError.Message, httpError.Code)
		return
	}

	utils.HttpResponse(w, b, http.StatusCreated)
}
