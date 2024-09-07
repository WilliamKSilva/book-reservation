package services_tests

import (
	"errors"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
	repositories_tests "github.com/WilliamKSilva/book-reservation/internal/infra/db/repositories/tests"
	"github.com/WilliamKSilva/book-reservation/internal/services"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

type MockedJwtServiceSuccess struct{}

func (jwt *MockedJwtServiceSuccess) New() (DTOs.JwtToken, error) {
	raw := `
		eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
		eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.
		SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
	`

	return DTOs.JwtToken{
		Raw:    raw,
		Signed: "tokensigned12",
	}, nil
}

func NewMockedJwtServiceSuccess() *MockedJwtServiceSuccess {
	return &MockedJwtServiceSuccess{}
}

type MockedUserRepositorySuccess struct{}

func (userRepository *MockedUserRepositorySuccess) Save(user domain.User) (domain.User, error) {
	u, err := repositories_tests.MockUser()
	return u, err
}

func (userRepository *MockedUserRepositorySuccess) FindByEmail(email string) (domain.User, error) {
	u, err := repositories_tests.MockUser()
	return u, err
}

func NewMockedUserRepositorySuccess() *MockedUserRepositorySuccess {
	return &MockedUserRepositorySuccess{}
}

type MockedUserRepositoryFailure struct{}

func (userRepository *MockedUserRepositoryFailure) Save(user domain.User) (domain.User, error) {
	return domain.User{}, errors.New("generic MockedUserRepositoryFailure save error")
}

func (userRepository *MockedUserRepositoryFailure) FindByEmail(email string) (domain.User, error) {
	return domain.User{}, errors.New("generic MockedUserRepositoryFailure save error")
}

func NewMockedUserRepositoryFailure() *MockedUserRepositoryFailure {
	return &MockedUserRepositoryFailure{}
}

type MockedUuidService struct{}

func (uuidService *MockedUuidService) Generate() string {
	return `
		904cf2f4-eb41-4512-bce1-a1082cc674f2
	`
}

func NewMockedUuidService() *MockedUuidService {
	return &MockedUuidService{}
}

type MockedUserServiceSuccess struct {
	UserRepository services.UserRepositoryAdapter
	UuidGenerator  services.UuidServiceAdapter
}

func NewMockedUserServiceSuccess() *MockedUserServiceSuccess {
	return &MockedUserServiceSuccess{
		UserRepository: NewMockedUserRepositorySuccess(),
		UuidGenerator:  NewMockedUuidService(),
	}
}

func (userService *MockedUserServiceSuccess) Create(req DTOs.CreateUserRequestDTO) (DTOs.CreateUserResponseDTO, error) {
	u, _ := repositories_tests.MockUser()
	return DTOs.CreateUserResponseDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CPF:       u.CPF,
		BirthDate: u.BirthDate.String(),
	}, nil
}

func (userService *MockedUserServiceSuccess) FindByEmail(email string) (DTOs.FindUserByEmailResponseDTO, error) {
	u, _ := repositories_tests.MockUser()
	return DTOs.FindUserByEmailResponseDTO{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		CPF:       u.CPF,
		BirthDate: u.BirthDate.String(),
	}, nil
}

type MockedUserServiceFailure struct {
	UserRepository services.UserRepositoryAdapter
	UuidGenerator  services.UuidServiceAdapter
}

func NewMockedUserServiceFailure() *MockedUserServiceFailure {
	return &MockedUserServiceFailure{
		UserRepository: NewMockedUserRepositorySuccess(),
		UuidGenerator:  NewMockedUuidService(),
	}
}

func (userService *MockedUserServiceFailure) Create(req DTOs.CreateUserRequestDTO) (DTOs.CreateUserResponseDTO, error) {
	return DTOs.CreateUserResponseDTO{}, errors.New("Generic MockedUserServiceFailure error")
}

func (userService *MockedUserServiceFailure) FindByEmail(email string) (DTOs.FindUserByEmailResponseDTO, error) {
	return DTOs.FindUserByEmailResponseDTO{}, errors.New("Generic MockedUserServiceFailure error")
}
