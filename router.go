package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"golang-api/handlers"
)

type Route struct {
	Name 		string
	Method 		string
	Pattern 	string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"GetTodos",
		"GET",
		"/todos",
		handlers.GetTodos,
	},
	Route{
		"GetTodo",
		"GET",
		"/todo/{id}",
		handlers.GetTodo,
	},
	Route{
		"CreateTodo",
		"POST",
		"/todo",
		handlers.CreateTodo,
	},
	Route{
		"UpdateTodo",
		"PUT",
		"/todo/{id}",
		handlers.UpdateTodo,
	},
	Route{
		"DeleteTodo",
		"DELETE",
		"/todo/{id}",
		handlers.DeleteTodo,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}