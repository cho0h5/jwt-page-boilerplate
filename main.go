package main

import (
	"net/http"
)

func main() {
	http.Handle("/", permissionCheckMiddleware(indexHandler))
	http.HandleFunc("/askPassword", askPasswordHandler)
	http.HandleFunc("/login", loginHandler)

	

	http.ListenAndServe(":8081", nil)
}