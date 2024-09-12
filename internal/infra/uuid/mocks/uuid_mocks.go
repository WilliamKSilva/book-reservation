package uuid_mocks

import services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"

type MockedUuidServiceSuccess struct{}

func (uuidService *MockedUuidServiceSuccess) Generate() (string, error) {
	return `
		904cf2f4-eb41-4512-bce1-a1082cc674f2
	`, nil
}

func NewMockedUuidServiceSuccess() *MockedUuidServiceSuccess {
	return &MockedUuidServiceSuccess{}
}

type MockedUuidServiceFailure struct{}

func (uuidService *MockedUuidServiceFailure) Generate() (string, error) {
	return "", &services_errors.InternalServerError{}
}

func NewMockedUuidServiceFailure() *MockedUuidServiceFailure {
	return &MockedUuidServiceFailure{}
}
