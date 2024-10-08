package services_tests

import (
	"errors"
	"reflect"
	"testing"

	"github.com/WilliamKSilva/book-reservation/internal/domain/user"
	repositories_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/db/repositories/mocks"
	encrypter_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/encrypter/mocks"
	uuid_mocks "github.com/WilliamKSilva/book-reservation/internal/infra/uuid/mocks"
	"github.com/WilliamKSilva/book-reservation/internal/services"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
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
	t.Run("should return ValidationError with Name field missing", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}

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

		var vErr *services_errors.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "Name" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with Email field missing", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}

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

		var vErr *services_errors.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "Email" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with Password field missing", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}
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

		var vErr *services_errors.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "Password" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with CPF field missing", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}
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

		var vErr *services_errors.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "CPF" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return ValidationError with BirthDate field missing", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}
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

		var vErr *services_errors.ValidationError
		if !errors.As(err, &vErr) {
			t.Errorf("Expected error of type *ValidationError, but got %T", err)
		}

		if vErr.Field != "BirthDate" {
			t.Errorf("Expected validation error for 'Name' field, but got %s", vErr.Field)
		}
	})

	t.Run("should return a InternalServerError if UuidService fails", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceFailure(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Email:     req.Email,
			Password:  req.Password,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected err, got nil")
		}

		var internalServerError *services_errors.InternalServerError
		if !errors.As(err, &internalServerError) {
			t.Errorf("Expected error of type *InternalServerError, but got %T", err)
		}
	})

	t.Run("should return a InternalServerError if EncrypterService fails", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceFailure(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceFailure(),
		}
		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Email:     req.Email,
			Password:  req.Password,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected err, got nil")
		}

		var internalServerError *services_errors.InternalServerError
		if !errors.As(err, &internalServerError) {
			t.Errorf("Expected error of type *InternalServerError, but got %T", err)
		}
	})

	t.Run("should return an empty CreateUserResponseDTO struct and an error if UserRepository fails", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositoryFailure(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}

		req := MockCreateUserRequestDTO()
		_, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Email:     req.Email,
			Password:  req.Password,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err == nil {
			t.Error("Expected UserRepository error, got nil")
		}
	})

	t.Run("should return an CreateUserResponseDTO struct and an empty error on success", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}
		req := MockCreateUserRequestDTO()
		u, err := userService.Create(DTOs.CreateUserRequestDTO{
			Name:      req.Name,
			Email:     req.Email,
			Password:  req.Password,
			CPF:       req.CPF,
			BirthDate: req.BirthDate,
		})

		if err != nil {
			t.Errorf("Expected nil, got error: %s", err.Error())
		}

		mocked, _ := user.MockUser()
		expected := DTOs.CreateUserResponseDTO{
			ID:        mocked.ID,
			Name:      mocked.Name,
			Email:     mocked.Email,
			Password:  mocked.Password,
			CPF:       mocked.CPF,
			BirthDate: mocked.BirthDate.String(),
		}

		if !reflect.DeepEqual(u, expected) {
			t.Error("CreateUserResponseDTO struct returned do not match with expected")
		}
	})
}

func TestUserServiceFindByEmail(t *testing.T) {
	t.Run("should return an empty FindUserByEmailResponseDTO struct and an error if UserRepository fails", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositoryFailure(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}

		u, err := userService.FindByEmail("teste@teste.com")

		if err == nil {
			t.Error("Expected err, got nil")
		}

		expected := DTOs.FindUserByEmailResponseDTO{}

		if !reflect.DeepEqual(u, expected) {
			t.Error("Expected empty FindUserByEmailResponseDTO")
		}
	})

	t.Run("should return an empty FindUserByEmailResponseDTO struct and a UserNotFoundError if UserRepository return empty User struct", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccessFindByEmailNotFound(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}

		u, err := userService.FindByEmail("teste@teste.com")

		if err == nil {
			t.Error("Expected a UserNotFoundError, got nil")
		}

		var notFoundErr *services_errors.UserNotFoundError
		if !errors.As(err, &notFoundErr) {
			t.Errorf("Expected error of type *UserNotFoundError, but got %T", err)
		}

		expected := DTOs.FindUserByEmailResponseDTO{}

		if !reflect.DeepEqual(u, expected) {
			t.Error("Expected empty FindUserByEmailResponseDTO")
		}
	})

	t.Run("should return a FindUserByEmailResponseDTO struct and no error on success", func(t *testing.T) {
		userService := services.UserService{
			UserRepository:   repositories_mocks.NewMockedUserRepositorySuccess(),
			UuidService:      uuid_mocks.NewMockedUuidServiceSuccess(),
			EncrypterService: encrypter_mocks.NewMockedEncrypterServiceSuccess(),
		}
		u, err := userService.FindByEmail("teste@teste.com")

		if err != nil {
			t.Error("Expected nil, got err")
		}

		mockedUser, _ := user.MockUser()
		expected := DTOs.FindUserByEmailResponseDTO{
			ID:        mockedUser.ID,
			Name:      mockedUser.Name,
			Email:     mockedUser.Email,
			Password:  mockedUser.Password,
			CPF:       mockedUser.CPF,
			BirthDate: mockedUser.BirthDate.String(),
		}

		if !reflect.DeepEqual(u, expected) {
			t.Error("FindUserByEmailResponseDTO returned does not match with expected")
		}
	})
}
