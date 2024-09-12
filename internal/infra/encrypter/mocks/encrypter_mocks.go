package encrypter_mocks

import services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"

type MockEncrypterServiceSuccess struct{}

func (encrypterService *MockEncrypterServiceSuccess) Hash(password string) (string, error) {
	return "$2a$12$nBx8Bs7knp0QXJUBPKwgeOXjSX.dc3RwLcy1PfJRjfY8pw1yGYB7m", nil
}

func NewMockedEncrypterServiceSuccess() *MockEncrypterServiceSuccess {
	return &MockEncrypterServiceSuccess{}
}

type MockEncrypterServiceFailure struct{}

func (encrypterService *MockEncrypterServiceFailure) Hash(password string) (string, error) {
	return "", &services_errors.InternalServerError{}
}

func NewMockedEncrypterServiceFailure() *MockEncrypterServiceFailure {
	return &MockEncrypterServiceFailure{}
}
