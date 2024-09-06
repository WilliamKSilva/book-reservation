package tests

import (
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

type MockedUserRepository struct{}

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

func (userRepository *MockedUserRepository) Save(user domain.User) (domain.User, error) {
	u, err := mockUser()
	return u, err
}

func (userRepository *MockedUserRepository) FindByEmail(email string) (domain.User, error) {
	u, err := mockUser()
	return u, err
}

func NewMockedUserRepository() *MockedUserRepository {
	return &MockedUserRepository{}
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
