package jwt_mocks

import "github.com/WilliamKSilva/book-reservation/internal/services/DTOs"

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
