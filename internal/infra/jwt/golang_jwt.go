package jwt

import (
	"os"

	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
	"github.com/golang-jwt/jwt/v5"
)

type GolangJwt struct{}

func (golangJwt *GolangJwt) New() (DTOs.JwtToken, error) {
	key := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	signed, err := token.SignedString([]byte(key))
	if err != nil {
		return DTOs.JwtToken{}, err
	}

	return DTOs.JwtToken{Raw: token.Raw, Signed: signed}, nil
}
