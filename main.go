package main

import (
	"net/http"

	// "github.com/golang-jwt/jwt/v4"
)

func main() {
	http.HandleFunc("/", indexHandler)

	http.Handle("/login", permissionCheckMiddleware(loginHandler))

	http.ListenAndServe(":8081", nil)
}