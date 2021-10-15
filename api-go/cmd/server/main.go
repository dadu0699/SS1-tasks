package main

import (
	"api-go/pkg/app"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func run() (err error) {
	app := app.New()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("[API] Server started")
	handler := cors.Default().Handler(app.Router)

	err = http.ListenAndServe(":5000", handler)
	return
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s: %s", "Error loading .env file", err)
	}

	if err := run(); err != nil {
		log.Fatalf("This is the startup error: %s", err)
		os.Exit(1)
	}
}
