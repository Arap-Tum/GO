package main

import (
	"contactManagerApp/internal/router"
	"encoding/json"
	"fmt"
	"net/http"
)

// ENTRY POINT
/*
	Entry Point
	starts HTTP server
	loads dependancies (store, router , etc)
*/

type HealthResponse struct {
	Status string `json:"status"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// SET header to JSON
	w.Header().Set("Content-Type", "aplication/json")
	w.Header().Set("Content-Disposition", "inline")

	// CREATE RESPONSE
	response := HealthResponse{
		Status: "ok",
	}

	// encode to json and send
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to ecode response", http.StatusInternalServerError)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// set header tojson
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "inline")

	// create a response
	response := HelloResponse{
		Message: "Hello welcome to Cotact manager app ",
	}

	//Encode  to json and send
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Faailed to encode response ", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/hello", helloHandler)
	router.SetupRoutes()

	fmt.Println("Starting server on :8080.....")

	err := http.ListenAndServe(":8080", nil) // start listening
	if err != nil {
		fmt.Println("Eror Starting server: ", err)
	}
}
