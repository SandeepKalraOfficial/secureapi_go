// File: main.go
package main

import (
	"SecureAPI/config"
	"SecureAPI/handlers"
	"SecureAPI/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load allowed origins from config (e.g., config/config.json)
	config.LoadConfig("config/config.json")

	router := mux.NewRouter()
	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
	router.HandleFunc("/people", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")
	router.HandleFunc("/people/{id}", handlers.UpdatePerson).Methods("PUT")

	// Apply middleware
	securedRouter := middleware.StrictOriginEnforcer(router)
	handler := middleware.GetCORSHandler().Handler(securedRouter)

	log.Println("🚀 Server running on http://localhost:8000")
	http.ListenAndServe(":8000", handler)
}
