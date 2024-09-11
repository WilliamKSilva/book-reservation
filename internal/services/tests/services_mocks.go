package services_tests

import (
	"errors"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
	repositories_tests "github.com/WilliamKSilva/book-reservation/internal/infra/db/repositories/tests"
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

// This is horrible, but we can't override the FindByEmail method of MockedUserRepoistorySuccess
type MockedUserRepositorySuccessFindByEmailNotFound struct{}

func (userRepository *MockedUserRepositorySuccessFindByEmailNotFound) Save(user domain.User) (domain.User, error) {
	u, err := repositories_tests.MockUser()
	return u, err
}

func (userRepository *MockedUserRepositorySuccessFindByEmailNotFound) FindByEmail(email string) (domain.User, error) {
	return domain.User{}, nil
}

func NewMockedUserRepositorySuccessFindByEmailNotFound() *MockedUserRepositorySuccessFindByEmailNotFound {
	return &MockedUserRepositorySuccessFindByEmailNotFound{}
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
