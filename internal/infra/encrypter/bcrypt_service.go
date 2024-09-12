package encrypter

import (
	"github.com/WilliamKSilva/book-reservation/internal/services"
	services_errors "github.com/WilliamKSilva/book-reservation/internal/services/errors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct{}

func (bcryptService *BcryptService) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		services.LogUnexpectedError("BcryptService", "Hash", err.Error())
		return "", &services_errors.InternalServerError{Message: err.Error()}
	}
	return string(hash), nil
}
