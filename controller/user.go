package controller

import (
	"fmt"
	"net/http"
	"time"

	"go-web-boilerplate/config"
	"go-web-boilerplate/model"

	uuid "github.com/satori/go.uuid"
)

func Bar(w http.ResponseWriter, req *http.Request) {
	u := GetUser(w, req)
	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	ShowSessions()

	config.TPL.ExecuteTemplate(w, "bar.gohtml", u)
}

func Signup(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("w", w)
	fmt.Println("req", req)

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		fmt.Println(" req.Method == http.MethodPost")
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		fmt.Println("post2")
		// username taken?
		//.Println("model.DBUsers[un]", model.DBUsers[un])
		if _, ok := model.DBUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		fmt.Println("2---------------------------------------------")

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = model.SessionLength
		http.SetCookie(w, c)
		fmt.Println("3---------------------------------------------")
		model.DBSessions[c.Value] = model.Session{un, time.Now()}
		// store user in dbUsers

		fmt.Println("4---------------------------------------------")
		u = model.User{un, p, f, l, r}
		fmt.Println("u", u)
		model.DBUsers[un] = u
		fmt.Println("model.DBUsers[un] ", model.DBUsers[un])

		existCheck, err := model.CheckExistUser(req)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("existCheck", existCheck)
		if existCheck == 0 {
			model.PutUser(req)
		} else {
			fmt.Println("existCheck else")
			http.Redirect(w, req, "/login", http.StatusConflict)
		}

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("222")
	// ShowSessions() // for demonstration purposes
	// tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	// tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	fmt.Println("333")
	config.TPL.ExecuteTemplate(w, "signup.gohtml", u)
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
	config.TPL.ExecuteTemplate(w, "login.gohtml", u)
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
