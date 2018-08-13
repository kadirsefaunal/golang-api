package models

type Todo struct {
	ID          int       `json:"id"`
	Content     string    `json:"content"`
	Status      bool      `json:"status"`
}
