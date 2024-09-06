package services

import (
	"github.com/WilliamKSilva/book-reservation/internal/domain"
	"github.com/WilliamKSilva/book-reservation/internal/services/DTOs"
)

type JwtServiceAdapter interface {
	New() (DTOs.JwtToken, error)
}

type UserRepositoryAdapter interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}

type UuidServiceAdapter interface {
	Generate() string
}
