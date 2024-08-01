package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

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

	// new router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerError)

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