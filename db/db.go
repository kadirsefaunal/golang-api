package db

import (
	"database/sql"
	"golang-api/models"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234qwe@/todo")
	models.CheckError(err)

	return db
}
