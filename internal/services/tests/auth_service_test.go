package services_tests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/WilliamKSilva/book-reservation/internal/domain/user"
	repositories_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/db/repositories/mocks"
	jwt_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/jwt/mocks"
	uuid_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/uuid/mocks"
	"github.com/WilliamKSilva/book-reservation/internal/services"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
)

func TestAuthServiceLogin(t *testing.T) {
	t.Run("should return an empty LoginResponseDTO struct and an error if UserService FindByEmail fails", func(t *testing.T) {
		userRepository := repositories_mocks.NewMockedUserRepositoryFailure()
		uuidService := uuid_mocks.NewMockedUuidServiceSuccess()
		userService := services.NewUserService(userRepository, uuidService)

		authService := services.AuthService{
			JwtService:  jwt_mocks.NewMockedJwtServiceSuccess(),
			UserService: userService,
		}

		res, err := authService.Login("teste", "teste12")

		if err == nil {
			t.Error("Expected err, got nil")
		}

		expected := DTOs.LoginResponseDTO{}

		if !reflect.DeepEqual(expected, res) {
			t.Error("Expected empty LoginResponseDTO struct, got populated struct")
		}
	})

	t.Run("should return an empty LoginResponseDTO struct and an user not found error if UserService FindByEmail don't find an user", func(t *testing.T) {
		userRepository := repositories_mocks.NewMockedUserRepositorySuccessFindByEmailNotFound()
		uuidGenerator := uuid_mocks.NewMockedUuidServiceSuccess()
		userService := services.NewUserService(userRepository, uuidGenerator)

		authService := services.AuthService{
			JwtService:  jwt_mocks.NewMockedJwtServiceSuccess(),
			UserService: userService,
		}

		res, err := authService.Login("teste", "teste12")

		if err == nil {
			t.Error("Expected err, got nil")
		}

		var notFoundErr *services_errors.UserNotFoundError
		if !errors.As(err, &notFoundErr) {
			t.Errorf("Expected error of type UserNotFoundError, but got %s", err.Error())
		}

		expected := DTOs.LoginResponseDTO{}

		if !reflect.DeepEqual(expected, res) {
			t.Error("Expected empty LoginResponseDTO struct, got populated struct")
		}
	})

	t.Run("should return a WrongPasswordError if it is provided a valid Email with wrong Password", func(t *testing.T) {
		userRepository := repositories_mocks.NewMockedUserRepositorySuccess()
		uuidGenerator := uuid_mocks.NewMockedUuidServiceSuccess()
		userService := services.NewUserService(userRepository, uuidGenerator)

		authService := services.AuthService{
			JwtService:  jwt_mocks.NewMockedJwtServiceSuccess(),
			UserService: userService,
		}

		u, _ := user.MockUser()

		res, err := authService.Login(u.Email, "wrong password")

		if err == nil {
			t.Error("Expected err, got nil")
		}

		var wrongPassError *services_errors.WrongPasswordError
		if !errors.As(err, &wrongPassError) {
			t.Errorf("Expected error of type WrongPasswordError, but got %s", err.Error())
		}

		expected := DTOs.LoginResponseDTO{}

		if !reflect.DeepEqual(expected, res) {
			t.Error("Expected empty LoginResponseDTO struct, got populated struct")
		}
	})

	t.Run("should return a InternalServerError if JwtService fails", func(t *testing.T) {
		userRepository := repositories_mocks.NewMockedUserRepositorySuccess()
		uuidGenerator := uuid_mocks.NewMockedUuidServiceSuccess()
		userService := services.NewUserService(userRepository, uuidGenerator)

		authService := services.AuthService{
			JwtService:  jwt_mocks.NewMockedJwtServiceFailure(),
			UserService: userService,
		}

		u, _ := user.MockUser()
		res, err := authService.Login(u.Email, u.Password)

		if err == nil {
			t.Error("Expected err, got nil")
		}

		var internalServerError *services_errors.InternalServerError
		if !errors.As(err, &internalServerError) {
			t.Errorf("Expected error of type InternalServerError, but got %s", err.Error())
		}

		expected := DTOs.LoginResponseDTO{}

		if !reflect.DeepEqual(expected, res) {
			t.Error("Expected empty LoginResponseDTO struct, got populated struct")
		}
	})

	t.Run("should return a LoginResponseDTO and no error on success", func(t *testing.T) {
		userRepository := repositories_mocks.NewMockedUserRepositorySuccess()
		uuidGenerator := uuid_mocks.NewMockedUuidServiceSuccess()
		userService := services.NewUserService(userRepository, uuidGenerator)

		authService := services.AuthService{
			JwtService:  jwt_mocks.NewMockedJwtServiceSuccess(),
			UserService: userService,
		}

		u, _ := user.MockUser()
		res, err := authService.Login(u.Email, u.Password)

		if err != nil {
			t.Errorf("Expected nil, got error: %s", err.Error())
		}

		expected := DTOs.LoginResponseDTO{
			User: DTOs.LoginUser{
				ID:        u.ID,
				Name:      u.Name,
				Email:     u.Email,
				Password:  u.Password,
				CPF:       u.CPF,
				BirthDate: u.BirthDate.String(),
			},
			AccessToken: DTOs.JwtToken{
				Raw:    jwt_mocks.MockedRawJwtToken,
				Signed: jwt_mocks.MockedSignedJwtToken,
			},
		}

		if !reflect.DeepEqual(expected, res) {
			t.Errorf("Expected LoginResponseDTO %T, got %T", expected, res)
		}
	})
}
