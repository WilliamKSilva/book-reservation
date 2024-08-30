package db

import (
	"context"
	"log"

	"github.com/WilliamKSilva/book-reservation/internal/app/user"
	"github.com/jackc/pgx/v5"
)

type PostgresUserRepository struct {
	Conn *pgx.Conn
}

func (userRepository *PostgresUserRepository) Save(userData user.User) (user.User, error) {
	var result string
	err := userRepository.Conn.QueryRow(
		context.Background(),
		"INSERT INTO users (id, name, email, cpf, birth_date) VALUES ($1, $2, $3, $4, $5)",
		userData.ID,
		userData.Name,
		userData.Email,
		userData.CPF,
		userData.BirthDate,
	).Scan(&result)

	log.Println(err)

	var user user.User
	return user, nil
}

func (userRepository *PostgresUserRepository) FindByEmail(email string) (user.User, error) {
	var result string
	err := userRepository.Conn.QueryRow(
		context.Background(),
		"SELECT * FROM users WHERE email = $1",
		email,
	).Scan(&result)

	log.Println(err)

	var user user.User
	return user, nil
}
