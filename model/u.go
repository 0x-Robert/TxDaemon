package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"go-web-boilerplate/config"
)

var count int

func PutUser(r *http.Request) (User, error) {
	u := User{}
	fmt.Println("Putuser Insert ~~~~~~", r)
	u.UserName = r.FormValue("username")
	// u.Password = getCode(r.FormValue("password"))
	u.Password = r.FormValue("password")
	u.First = r.FormValue("first")
	u.Last = r.FormValue("last")
	u.Role = r.FormValue("role")
	fmt.Println("u Println ~~~ ", u)

	id, err := config.DB.Query("select count(name) from account")
	// LastInsertId, err := id.LastInsertId()
	// RowsAffected, err := id.RowsAffected()
	// fmt.Println("LastInsertId RowsAffected", LastInsertId, RowsAffected)
	// fmt.Println("id count : ")
	// if err != nil {
	// 	return u, errors.New("500. Internal Server Error." + err.Error())
	// }
	fmt.Println("id", id)
	for id.Next() {

		if err := id.Scan(&count); err != nil {
			log.Fatal(err)
		}
		fmt.Println("count", count)
	}
	count = count + 1
	fmt.Println("count2", count)
	sql, err := config.DB.Exec("INSERT INTO account ( id, name, password, first, last, role) VALUES ($1, $2, $3, $4, $5, $6)", count, u.UserName, u.Password, u.First, u.Last, u.Role)
	// sql, err := config.DB.Exec("INSERT INTO account ( name, password, first, last, role) VALUES ($1, $2, $3, $4, $5 )", u.UserName, u.Password, u.First, u.Last, u.Role)
	if err != nil {
		return u, errors.New("500. Internal Server Error." + err.Error())
	}
	fmt.Println("sql", sql)

	return u, nil
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 로그인시 가입한 계정과 비밀번호가 맞는지 체크
func CheckUserDb(r *http.Request) (User, error) {
	u := User{}
	u.UserName = r.FormValue("username")
	u.Password = getCode(r.FormValue("password"))

	row := config.DB.QueryRow("SELECT * FROM accout WHERE name = $1 and password = $2", u.UserName, u.Password)

	err := row.Scan(&u.UserName, &u.Password)
	if err != nil {
		return u, err
	}
	fmt.Println("row", row)

	return u, nil
}

// 이미 존재하는 유저의 아이디와 비밀번호가 맞는지 체크
func CheckExistUser(r *http.Request) (User, error) {
	u := User{}
	u.UserName = r.FormValue("username")

	row := config.DB.QueryRow("SELECT * FROM accout WHERE name = $1", u.UserName)

	err := row.Scan(&u.UserName, &u.Password)
	if err != nil {
		return u, err
	}
	fmt.Println("row", row)

	return u, nil
}
