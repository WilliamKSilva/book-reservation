package db

import (
	"context"
	"log"

	"github.com/WilliamKSilva/book-reservation/internal/app/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	Conn *pgxpool.Pool
}

const QUERY_ALL_USER_FIELDS = "SELECT id, name, email, cpf, password, birth_date FROM users "

func (userRepository *PostgresUserRepository) Save(userData user.User) (user.User, error) {
	log.Println(userData)
	rows, err := userRepository.Conn.Query(
		context.Background(),
		"INSERT INTO users (id, name, email, cpf, password, birth_date) VALUES ($1, $2, $3, $4, $5, $6)",
		userData.ID,
		userData.Name,
		userData.Email,
		userData.CPF,
		userData.Password,
		userData.BirthDate,
	)

	if err != nil {
		if err != pgx.ErrNoRows {
			return user.User{}, err
		}
	}

	defer rows.Close()

	rows, err = userRepository.Conn.Query(
		context.Background(),
		QUERY_ALL_USER_FIELDS+"WHERE id = $1",
		userData.ID,
	)

	if err != nil {
		return user.User{}, err
	}

	defer rows.Close()

	var u user.User
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CPF, &u.Password, &u.BirthDate)
		if err != nil {
			return user.User{}, err
		}
	}

	return u, nil
}

func (userRepository *PostgresUserRepository) FindByEmail(email string) (user.User, error) {
	rows, err := userRepository.Conn.Query(
		context.Background(),
		QUERY_ALL_USER_FIELDS+"WHERE email = $1",
		email,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return user.User{}, err
		}
	}

	var u user.User
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CPF, &u.Password, &u.BirthDate)
		if err != nil {
			return user.User{}, err
		}
	}

	return u, nil
}
