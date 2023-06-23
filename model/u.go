package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go-web-boilerplate/config"
)

func PutUser(r *http.Request) (User, error) {
	u := User{}
	fmt.Println("Putuser ", r)
	u.UserName = r.FormValue("username")
	u.Password = getCode(r.FormValue("password"))
	u.First = r.FormValue("first")
	u.Last = r.FormValue("last")
	u.Role = r.FormValue("role")

	_, err := config.DB.Exec("INSERT INTO account (name, password, first, last, role) VALUES ($1, $2, $3, $4 $5)", u.UserName, u.Password, u.First, u.Last, u.Role)
	if err != nil {
		return u, errors.New("500. Internal Server Error." + err.Error())
	}

	return u, nil
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
