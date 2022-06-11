package main

import (
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	password, err := getPassword()
	if err != nil {
		// redirect to ask password
	}
	// ask password

	w.Write([]byte("Hello World"))
	// redirect to index
}

func askPasswordHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("what your password?"))
	// return template, form
}

func setPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// set password
	
	// redirect loginHandler
}

func permissionCheckMiddleware(next http.HandlerFunc) http.Handler {
	// check is exist jwt token
	// if not, redirect login handler

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}