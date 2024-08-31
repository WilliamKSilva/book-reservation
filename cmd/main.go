package main

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/WilliamKSilva/book-reservation/internal/infra/db"
	"github.com/WilliamKSilva/book-reservation/internal/infra/web"
	"github.com/joho/godotenv"
)

func generateSwaggerDocs() {
	app := "swag"
	arg0 := "init"
	// Defines the directories to search for docs
	arg1 := "-d"
	arg2 := "internal/infra/web/,internal/infra/web/handlers/,internal/app/user/"
	arg3 := "-g"
	arg4 := "http_server.go"
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	log.Println(string(stdout[:]))
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		log.Println(arg)
		if arg == "--docs" {
			generateSwaggerDocs()
		}
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env")
	}

	conn := db.Connect()
	defer conn.Close(context.Background())
	web.StartListening(8080, conn)
}
