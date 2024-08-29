package main

import (
	"context"
	"log"

	"github.com/WilliamKSilva/book-reservation/internal/app/user"
	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/WilliamKSilva/book-reservation/internal/infra/uuid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env")
	}

	conn := db.Connect()
	defer conn.Close(context.Background())

	userRepository := db.PostgresUserRepository{Conn: conn}
	uuidGenerator := uuid.GoogleUUIDGenerator{}
	userService := user.NewUserService(userRepository, uuidGenerator)

	userService.Create("william", "teste", "52293872845", "2023-08-12")
}
