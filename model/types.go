package model

import (
	"text/template"
	"time"
)

const SessionLength int = 30

var (
	Tpl               *template.Template
	DBUsers           = map[string]User{}    // user ID, user
	DBSessions        = map[string]Session{} // session ID, user ID
	DBSessionsCleaned time.Time
)

type Block struct {
	BlockNumber string
}

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
