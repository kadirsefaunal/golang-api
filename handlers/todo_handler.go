package handlers

import (
		"fmt"
		"net/http"
		"encoding/json"
	"golang-api/repositories"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := repositories.GetAll()
	json.NewEncoder(w).Encode(todos)
}
