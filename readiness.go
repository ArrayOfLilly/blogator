package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	jsonResponse := map[string]string{
		"status": "ok",
	}
	
	respondWithJSON(w, 200, jsonResponse)
}
