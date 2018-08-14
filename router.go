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
	IsProtected	bool
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
		false,
	},
	Route{
		"GetTodos",
		"GET",
		"/todos",
		handlers.GetTodos,
		false,
	},
	Route{
		"GetTodo",
		"GET",
		"/todo/{id}",
		handlers.GetTodo,
		true,
	},
	Route{
		"CreateTodo",
		"POST",
		"/todo",
		handlers.CreateTodo,
		true,
	},
	Route{
		"UpdateTodo",
		"PUT",
		"/todo/{id}",
		handlers.UpdateTodo,
		true,
	},
	Route{
		"DeleteTodo",
		"DELETE",
		"/todo/{id}",
		handlers.DeleteTodo,
		true,
	},
	Route{
		"CreateUser",
		"POST",
		"/user",
		handlers.CreateUser,
		false,
	},
	Route{
		"Login",
		"POST",
		"/login",
		handlers.Login,
		false,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		if route.IsProtected {
			handler = AuthMiddleware(route.HandlerFunc)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}