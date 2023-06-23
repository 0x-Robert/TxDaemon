package controller

import (
	"fmt"
	"net/http"
	"time"

	"go-web-boilerplate/model"

	uuid "github.com/satori/go.uuid"
)

var u model.User

func GetUser(w http.ResponseWriter, req *http.Request) model.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = model.SessionLength
	http.SetCookie(w, c)

	// if the user exists already, get user

	if s, ok := model.DBSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		model.DBSessions[c.Value] = s
		u = model.DBUsers[s.Un]
	}
	return u
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := model.DBSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		model.DBSessions[c.Value] = s
	}
	_, ok = model.DBUsers[s.Un]
	// refresh session
	c.MaxAge = model.SessionLength
	http.SetCookie(w, c)
	return ok
}

func CleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	ShowSessions()              // for demonstration purposes
	for k, v := range model.DBSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(model.DBSessions, k)
		}
	}
	model.DBSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	ShowSessions()             // for demonstration purposes
}

// for demonstration purposes
func ShowSessions() {
	fmt.Println("********")
	for k, v := range model.DBSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
