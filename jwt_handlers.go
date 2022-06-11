package main

import (
	"log"
	_ "embed"
	"html/template"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

//go:embed templates/askPasswordTemplate.html
var askPasswordTemplate string

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	if r.FormValue("password") != getPassword() {
		w.Write([]byte("Password is not valid."))
	}

	// publish jwt
	claims := &jwt.StandardClaims {
		ExpiresAt: 15000,
		Issuer: "me",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(signKey))

	http.SetCookie(w, &http.Cookie {
		Name: "jwt",
		Value: signedToken,
		Path: "/",
	})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func askPasswordHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("askPassword")
	t, _ = t.Parse(askPasswordTemplate)
	t.Execute(w, nil)
}

func permissionCheckMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Redirect(w, r, "/askPassword", http.StatusTemporaryRedirect)
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(signKey), nil
		})
		if err != nil {
			http.Redirect(w, r, "/askPassword", http.StatusTemporaryRedirect)
		}
		log.Println(token)

		next.ServeHTTP(w, r)
	})
}