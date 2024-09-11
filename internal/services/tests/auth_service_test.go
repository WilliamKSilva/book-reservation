package services_tests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/WilliamKSilva/book-reservation/internal/services"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

func TestAuthServiceLogin(t *testing.T) {
	t.Run("should return an empty LoginResponseDTO struct and an error if UserService FindByEmail fails", func(t *testing.T) {
		userRepository := NewMockedUserRepositoryFailure()
		uuidGenerator := NewMockedUuidService()
		userService := services.NewUserService(userRepository, uuidGenerator)

		authService := services.AuthService{
			JwtService:  NewMockedJwtServiceSuccess(),
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
		userRepository := NewMockedUserRepositorySuccessFindByEmailNotFound()
		uuidGenerator := NewMockedUuidService()
		userService := services.NewUserService(userRepository, uuidGenerator)

		authService := services.AuthService{
			JwtService:  NewMockedJwtServiceSuccess(),
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
