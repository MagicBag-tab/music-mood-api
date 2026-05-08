package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"music-mood-api/internal/handlers"
	database "music-mood-api/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env vars")
	}

	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("connected to database")

	router := handlers.NewRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8009"
	}

	log.Printf("server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
