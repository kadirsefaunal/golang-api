package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"golang-api/repositories"
	"github.com/gorilla/mux"
	"strconv"
	"golang-api/models"
	"github.com/asaskevich/govalidator"
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

	result, err := govalidator.ValidateStruct(todo)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	if result {
		id := repositories.Insert(todo)
		json.NewEncoder(w).Encode(id)
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	models.CheckError(err)

	todo := repositories.Get(id)
	json.NewDecoder(r.Body).Decode(&todo)

	result, err := govalidator.ValidateStruct(todo)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}

	if result {
		affected := repositories.Update(todo)
		json.NewEncoder(w).Encode(affected)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	models.CheckError(err)

	affected := repositories.Delete(id)
	json.NewEncoder(w).Encode(affected)
}