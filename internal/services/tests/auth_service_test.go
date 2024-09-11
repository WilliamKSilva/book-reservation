package services_tests

import (
	"errors"
	"reflect"
	"testing"

	repositories_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/db/repositories/mocks"
	jwt_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/jwt/mocks"
	uuid_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/uuid/mocks"
	"github.com/WilliamKSilva/book-reservation/internal/services"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

func TestAuthServiceLogin(t *testing.T) {
	t.Run("should return an empty LoginResponseDTO struct and an error if UserService FindByEmail fails", func(t *testing.T) {
		userRepository := repositories_mocks.NewMockedUserRepositoryFailure()
		uuidGenerator := uuid_mocks.NewMockedUuidService()
		userService := services.NewUserService(userRepository, uuidGenerator)

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
		uuidGenerator := uuid_mocks.NewMockedUuidService()
		userService := services.NewUserService(userRepository, uuidGenerator)

		authService := services.AuthService{
			JwtService:  jwt_mocks.NewMockedJwtServiceSuccess(),
			UserService: userService,
		}

		res, err := authService.Login("teste", "teste12")

		if err == nil {
			t.Error("Expected err, got nil")
		}

		var notFoundErr *services.UserNotFoundError
		if !errors.As(err, &notFoundErr) {
			t.Errorf("Expected error of type UserNotFoundError, but got %s", err.Error())
		}

		expected := DTOs.LoginResponseDTO{}

		if !reflect.DeepEqual(expected, res) {
			t.Error("Expected empty LoginResponseDTO struct, got populated struct")
		}
	})
}
