package repositories

import (
	"golang-api/db"
	"golang-api/models"
)

func GetAll() []models.Todo {
	db := db.Connect()
	defer db.Close()

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

func Get(id int) models.Todo {
	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM todo WHERE id=?", id)
	models.CheckError(err)

	todo := models.Todo{}

	for rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Content, &todo.Status)
		models.CheckError(err)
	}

	return todo
}

func Insert(todo models.Todo) int {
	db := db.Connect()
	defer db.Close()

	row, err := db.Prepare("INSERT INTO todo (content, status) VALUES (?, ?)")
	models.CheckError(err)

	res, err := row.Exec(todo.Content, todo.Status)
	models.CheckError(err)

	id, err := res.LastInsertId()
	models.CheckError(err)

	return int(id)
}