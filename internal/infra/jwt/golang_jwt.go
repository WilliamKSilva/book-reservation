package jwt

import (
	"os"

	"github.com/WilliamKSilva/book-reservation/internal/app/auth"
	"github.com/golang-jwt/jwt/v5"
)

type GolangJwt struct{}

func (golangJwt *GolangJwt) New() (auth.JwtToken, error) {
	key := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	signed, err := token.SignedString([]byte(key))
	if err != nil {
		return auth.JwtToken{}, err
	}

	return auth.JwtToken{Raw: token.Raw, Signed: signed}, nil
}
