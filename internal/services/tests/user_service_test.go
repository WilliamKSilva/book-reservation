package tests

import (
	"errors"
	"testing"

	"github.com/WilliamKSilva/book-reservation/internal/services"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

func MockCreateUserRequestDTO() DTOs.CreateUserRequestDTO {
	return DTOs.CreateUserRequestDTO{
		Name:      "john doe",
		Email:     "john@gmail.com",
		Password:  "teste12345",
		CPF:       "23332212332",
		BirthDate: "2024-08-13",
	}
}

func TestUserServiceCreate(t *testing.T) {
	userService := services.UserService{
		UserRepository: NewMockedUserRepositorySuccess(),
		UuidGenerator:  NewMockedUuidService(),
	}

	t.Run("should return ValidationError with Name field missing", func(t *testing.T) {
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Email:     req.Email,
			Password:  req.Password,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected a validate data error, got nil")
		}

		var vErr *services.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "Name" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with Email field missing", func(t *testing.T) {
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Password:  req.Password,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected a validate data error, got nil")
		}

		var vErr *services.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "Email" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with Password field missing", func(t *testing.T) {
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Email:     req.Email,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected a validate data error, got nil")
		}

		var vErr *services.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "Password" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with CPF field missing", func(t *testing.T) {
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Email:     req.Email,
			Password:  req.Password,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected a validate data error, got nil")
		}

		var vErr *services.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "CPF" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with BirthDate field missing", func(t *testing.T) {
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			CPF:      req.CPF,
		})

		if err == nil {
			t.Error("Expected a validate data error, got nil")
		}

		var vErr *services.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "BirthDate" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with BirthDate field missing", func(t *testing.T) {
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			CPF:      req.CPF,
		})

		if err == nil {
			t.Error("Expected a validate data error, got nil")
		}

		var vErr *services.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "BirthDate" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return an empty user struct and an error if UserRepository fails", func(t *testing.T) {
		userService.UserRepository = NewMockedUserRepositoryFailure()
		req := MockCreateUserRequestDTO()
		u, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Email:     req.Email,
			Password:  req.Password,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected UserRepository error, got nil")
		}

		if u.ID != "" {
			t.Error("User returned by UserRepository should be empty")
		}
	})
}
