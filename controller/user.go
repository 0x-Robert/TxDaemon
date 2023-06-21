package controller

import (
	"fmt"
	"net/http"
	"time"

	"block-go-web/model"

	uuid "github.com/satori/go.uuid"
)

// func Bar(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("bar", bar)
// 	u := GetUser(w, req)
// 	if !AlreadyLoggedIn(req) {
// 		http.Redirect(w, req, "/", http.StatusSeeOther)
// 		return
// 	}
// 	if u.Role != "007" {
// 		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
// 		return
// 	}
// 	model.Tpl.ExecuteTemplate(w, "bar.gohtml", u)
// }

func Signup(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// var u model.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := model.DBUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = model.SessionLength
		http.SetCookie(w, c)
		// dbSessions[c.Value] = session{un, time.Now()}
		model.DBSessions[c.Value] = model.Session{un, time.Now()}
		// store user in dbUsers
		// bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		// if err != nil {
		// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
		// 	return
		// }
		u = model.User{un, p, f, l, r}
		model.DBUsers[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	ShowSessions() // for demonstration purposes
	model.Tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func Login(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// var u model.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := model.DBUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		// err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		// if err != nil {
		// 	http.Error(w, "Username and/or password do not match", http.StatusForbidden)
		// 	return
		// }
		// create session
		// p2 := getCode(p)
		fmt.Println(u, p)

		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = model.SessionLength
		http.SetCookie(w, c)
		model.DBSessions[c.Value] = model.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	ShowSessions() // for demonstration purposes
	model.Tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func Logout(w http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(model.DBSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(model.DBSessionsCleaned) > (time.Second * 30) {
		go CleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
