package main

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}

		token, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { // TODO: Mantık düzenlenecek.
			fmt.Println(claims["id"])
		} else {
			panic(err.Error())
		}

		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	signingKey := []byte("secret")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token, err
}