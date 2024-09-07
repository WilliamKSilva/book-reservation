package services_tests

import (
	"reflect"
	"testing"

	"github.com/WilliamKSilva/book-reservation/internal/services"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

func TestAuthServiceLogin(t *testing.T) {
	authService := services.AuthService{
		JwtService:  NewMockedJwtServiceSuccess(),
		UserService: NewMockedUserServiceSuccess(),
	}

	t.Run("should return an empty LoginResponseDTO struct and an error if UserService FindByEmail fails", func(t *testing.T) {
		authService.UserService = NewMockedUserServiceFailure()

		res, err := authService.Login("teste", "teste12")

		if err == nil {
			t.Error("Expected err, got nil")
		}

		expected := DTOs.LoginResponseDTO{}

		if !reflect.DeepEqual(expected, res) {
			t.Error("Expected empty LoginResponseDTO struct, got populated struct")
		}
	})
}
