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
	// TODO: traverse all directories inside "internal/"" for go files
	arg2 := "internal/infra/web/,internal/infra/web/handlers/,internal/app/user/,internal/app/auth/"
	arg3 := "-g"
	arg4 := "http_server.go"
	cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4)
	stdout, stderr := cmd.CombinedOutput()
	if stderr != nil {
		log.Printf("Command failed with error: %v\nOutput:\n%s", stderr, stdout)
		os.Exit(1)
	}

	log.Printf("Command output:\n%s", stdout)
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

	ctx := context.Background()

	pool := db.Connect(ctx)
	defer pool.Close()

	web.StartListening(ctx, 8080, pool)
}
