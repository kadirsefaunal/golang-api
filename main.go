package main

import (
	"golang-api/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Index).Methods("GET")
	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", handlers.GetTodo).Methods("GET")
	router.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", handlers.UpdateTodo).Methods("PUT")
	/*router.HandleFunc("/todo/{id}", index).Methods("DELETE")*/

	log.Fatal(http.ListenAndServe(":8089", router))
}
