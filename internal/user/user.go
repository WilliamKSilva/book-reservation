package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CPF       string
	BirthDate time.Time
}
