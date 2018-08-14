package handlers

import (
	"net/http"
	"golang-api/models"
	"encoding/json"
	"golang-api/repositories"
	)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	id := repositories.InsertUser(user)
	json.NewEncoder(w).Encode(id)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	token := repositories.GetUser(user)
	json.NewEncoder(w).Encode(token)
}