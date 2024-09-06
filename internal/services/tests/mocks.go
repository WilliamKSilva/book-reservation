package tests

import (
	"errors"
	"time"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
)

type MockedJwtService struct{}

func (jwt *MockedJwtService) New() string {
	return `
		eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
		eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.
		SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
	`
}

func NewMockedJwtService() *MockedJwtService {
	return &MockedJwtService{}
}

type MockedUserRepositorySuccess struct{}

func mockUser() (domain.User, error) {
	birthDateTime, err := time.Parse("2006-01-02", "2024-08-13")
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:        "ff508158-8da7-4840-b891-38c240f9aee1",
		Name:      "johndoe",
		Email:     "johndoe@teste.com",
		CPF:       "12312312323",
		Password:  "teste1234",
		BirthDate: birthDateTime,
	}, err
}

func (userRepository *MockedUserRepositorySuccess) Save(user domain.User) (domain.User, error) {
	u, err := mockUser()
	return u, err
}

func (userRepository *MockedUserRepositorySuccess) FindByEmail(email string) (domain.User, error) {
	u, err := mockUser()
	return u, err
}

func NewMockedUserRepositorySuccess() *MockedUserRepositorySuccess {
	return &MockedUserRepositorySuccess{}
}

type MockedUserRepositoryFailure struct{}

func (userRepository *MockedUserRepositoryFailure) Save(user domain.User) (domain.User, error) {
	return domain.User{}, errors.New("Could not save User in the Database")
}

func (userRepository *MockedUserRepositoryFailure) FindByEmail(email string) (domain.User, error) {
	u, err := mockUser()
	return u, err
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
