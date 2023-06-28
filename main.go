package main

import (
	"net/http"
	"time"

	"go-web-boilerplate/config"
	cont "go-web-boilerplate/controller"
	"go-web-boilerplate/model"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", cont.Bar)
	http.HandleFunc("/signup", cont.Signup)
	http.HandleFunc("/login", cont.Login)
	http.HandleFunc("/logout", cont.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func init() {
	config.Init()
	model.DBSessionsCleaned = time.Now()
}

func index(w http.ResponseWriter, req *http.Request) {
	u := cont.GetUser(w, req)
	cont.ShowSessions() // for demonstration purposes
	config.TPL.ExecuteTemplate(w, "index.gohtml", u)
}
