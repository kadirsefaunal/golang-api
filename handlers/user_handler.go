package handlers

import (
	"net/http"
	"golang-api/models"
	"encoding/json"
	"golang-api/repositories"
	"github.com/asaskevich/govalidator"
	)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result, err := govalidator.ValidateStruct(user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	if result {
		id := repositories.InsertUser(user)
		json.NewEncoder(w).Encode(id)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	result, err := govalidator.ValidateStruct(user)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	if result {
		token := repositories.GetUser(user)
		json.NewEncoder(w).Encode(token)
	}
}