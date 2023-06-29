package model

import (
	"time"
)

const SessionLength int = 30

var (
	DBUsers           = map[string]User{}    // user ID, user
	DBSessions        = map[string]Session{} // session ID, user ID
	DBSessionsCleaned time.Time
)

type User struct {
	UserName string
	Password string
	First    string
	Last     string
	Role     string
}

type Session struct {
	Un           string
	LastActivity time.Time
}

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}
