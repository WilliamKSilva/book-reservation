package repositories

import (
	"context"

	"github.com/WilliamKSilva/book-reservation/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepository struct {
	Conn *pgxpool.Pool
	Ctx  context.Context
}

const QUERY_ALL_USER_FIELDS = "SELECT id, name, email, cpf, password, birth_date FROM users "

func (userRepository *PostgresUserRepository) Save(userData domain.User) (domain.User, error) {
	_, err := userRepository.Conn.Exec(
		userRepository.Ctx,
		"INSERT INTO users (id, name, email, cpf, password, birth_date) VALUES ($1, $2, $3, $4, $5, $6)",
		userData.ID,
		userData.Name,
		userData.Email,
		userData.CPF,
		userData.Password,
		userData.BirthDate,
	)

	if err != nil {
		return domain.User{}, err
	}

	rows, err := userRepository.Conn.Query(
		userRepository.Ctx,
		QUERY_ALL_USER_FIELDS+"WHERE id = $1",
		userData.ID,
	)

	if err != nil {
		return domain.User{}, err
	}

	defer rows.Close()

	var u domain.User
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CPF, &u.Password, &u.BirthDate)
		if err != nil {
			return domain.User{}, err
		}
	}

	return u, nil
}

func (userRepository *PostgresUserRepository) FindByEmail(email string) (domain.User, error) {
	rows, err := userRepository.Conn.Query(
		context.Background(),
		QUERY_ALL_USER_FIELDS+"WHERE email = $1",
		email,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return domain.User{}, err
		}
	}

	var u domain.User
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CPF, &u.Password, &u.BirthDate)
		if err != nil {
			return domain.User{}, err
		}
	}

	return u, nil
}
