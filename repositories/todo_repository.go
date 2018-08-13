package repositories

import (
	"golang-api/db"
	"golang-api/models"
)

func GetAll() []models.Todo {
	db := db.Connect()

	rows, err := db.Query("SELECT * FROM todo")
	models.CheckError(err)

	todo := models.Todo{}
	todos := []models.Todo{}

	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Content, &todo.Status)
		models.CheckError(err)

		todos = append(todos, todo)
	}

	return todos
}