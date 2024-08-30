package main

import (
	"context"
	"log"

	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/WilliamKSilva/book-reservation/internal/infra/web"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env")
	}

	conn := db.Connect()
	defer conn.Close(context.Background())
	web.StartListening(8080, conn)
}
