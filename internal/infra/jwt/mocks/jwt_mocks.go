package jwt_mocks

import (
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
)

type MockedJwtServiceSuccess struct{}

const MockedRawJwtToken string = `
		eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
		eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.
		SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
	`
const MockedSignedJwtToken string = "tokensigned12"

func (jwt *MockedJwtServiceSuccess) New() (DTOs.JwtToken, error) {
	return DTOs.JwtToken{
		Raw:    MockedRawJwtToken,
		Signed: MockedSignedJwtToken,
	}, nil
}

func NewMockedJwtServiceSuccess() *MockedJwtServiceSuccess {
	return &MockedJwtServiceSuccess{}
}

type MockedJwtServiceFailure struct{}

func (jwt *MockedJwtServiceFailure) New() (DTOs.JwtToken, error) {
	return DTOs.JwtToken{}, &services_errors.InternalServerError{}
}

func NewMockedJwtServiceFailure() *MockedJwtServiceFailure {
	return &MockedJwtServiceFailure{}
}
