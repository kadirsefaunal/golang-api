package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"golang-api/repositories"
	"github.com/gorilla/mux"
	"strconv"
	"golang-api/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := repositories.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	models.CheckError(err)

	todo := repositories.Get(id)
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	id := repositories.Insert(todo)
	json.NewEncoder(w).Encode(id)
}
