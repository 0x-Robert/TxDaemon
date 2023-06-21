package main

import (
	"net/http"

	cont "block-go-web/controller"
)

func main() {
	http.HandleFunc("/", cont.Index)
	http.HandleFunc("/signup", cont.Signup)
	http.HandleFunc("/login", cont.Login)
	http.HandleFunc("/logout", cont.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
