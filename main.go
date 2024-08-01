package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/ArrayOfLilly/blogator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main(){
	filepathRoot := "."

	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
  }

  	port := os.Getenv("PORT")
  	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("CONN")
	if dbURL == "" {
		log.Fatal("CONN environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("unsucessful database connection")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	// new router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	mux.HandleFunc("POST /v1/users", apiCfg.handleUsersCreate)

	// new server (port and router)
	srv := &http.Server{
  		Addr:                         ":" + port,
  		Handler:                      mux,
  	}

	// starting server
	fmt.Printf("Server is started. Serving files on %s and listening on %s port\n", filepathRoot, port)
  	if err = srv.ListenAndServe(); err != nil {
		log.Fatal("Error starting server")
  	}

}