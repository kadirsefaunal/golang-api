package models

type User struct {
	ID 			int 	`json:"id" valid:"-"`
	UserName 	string 	`json:"userName" valid:"required~userName is required"`
	Password 	string 	`json:"password" valid:"required~password is required"`
}