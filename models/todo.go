package models

type Todo struct {
	ID          int       `json:"id" valid:"-"`
	Content     string    `json:"content" valid:"required~content is required"`
	Status      bool      `json:"status" valid:"optional"`
}
