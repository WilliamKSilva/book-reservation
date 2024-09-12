package services

import (
	"github.com/WilliamKSilva/book-reservation/internal/domain/user"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

type JwtServiceInterface interface {
	New() (DTOs.JwtToken, error)
}

type UserRepositoryInterface interface {
	Save(user user.User) (user.User, error)
	FindByEmail(email string) (user.User, error)
}

type UuidServiceInterface interface {
	Generate() (string, error)
}

type EncrypterServiceInterface interface {
	Hash(password string) (string, error)
}
