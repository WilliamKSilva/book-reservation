package repositories_tests

import (
	"time"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
)

func MockUser() (domain.User, error) {
	birthDateTime, err := time.Parse("2006-01-02", "2024-08-13")
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:        "ff508158-8da7-4840-b891-38c240f9aee1",
		Name:      "johndoe",
		Email:     "johndoe@teste.com",
		CPF:       "12312312323",
		Password:  "teste1234",
		BirthDate: birthDateTime,
	}, err
}
