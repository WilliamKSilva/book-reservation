package user

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	CPF       string
	Password  string
	BirthDate time.Time
}
