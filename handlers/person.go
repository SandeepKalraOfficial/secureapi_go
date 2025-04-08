package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"SecureAPI/models"

	"github.com/gorilla/mux"
)

// Simulated in-memory DB
var people []models.Person

// API handlers
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, p := range people {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.NotFound(w, r)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p models.Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	people = append(people, p)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, p := range people {
		if p.ID == params["id"] {
			var updated models.Person
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			updated.ID = p.ID
			people[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, p := range people {
		if p.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			fmt.Fprintf(w, "Deleted person with ID %s", p.ID)
			return
		}
	}
	http.NotFound(w, r)
}
