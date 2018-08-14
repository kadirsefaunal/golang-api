package repositories

import (
	"golang-api/models"
	"golang-api/db"

	"golang.org/x/crypto/bcrypt"
			"github.com/dgrijalva/jwt-go"
	)

func InsertUser(user models.User) int {
	db := db.Connect()
	defer db.Close()

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	models.CheckError(err)

	user.Password = string(hashedPwd)

	row, err := db.Prepare("INSERT INTO user (username, password) VALUES (?, ?)")
	models.CheckError(err)

	res, err := row.Exec(user.UserName, user.Password)
	models.CheckError(err)

	id, err := res.LastInsertId()
	models.CheckError(err)

	return int(id)
}

func GetUser(user models.User) string {
	db := db.Connect()
	defer db.Close()

	row, err := db.Query("SELECT * FROM user WHERE username=?", user.UserName)
	models.CheckError(err)

	var userFromDb models.User

	for row.Next() {
		err := row.Scan(&userFromDb.ID, &userFromDb.UserName, &userFromDb.Password)
		models.CheckError(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	models.CheckError(err)	// TODO: Parola hatalı mesajı dön.

	tokenString := getToken(userFromDb.ID)
	return string(tokenString)
}

func getToken(userId int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})

	tokenString, err := token.SignedString([]byte("secret")) // TODO: Key değiştir.
	models.CheckError(err)

	return tokenString
}