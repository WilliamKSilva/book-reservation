package services

import (
	"github.com/WilliamKSilva/book-reservation/internal/domain/user"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

type JwtServiceAdapter interface {
	New() (DTOs.JwtToken, error)
}

type UserRepositoryAdapter interface {
	Save(user user.User) (user.User, error)
	FindByEmail(email string) (user.User, error)
}

type UuidServiceAdapter interface {
	Generate() string
}
