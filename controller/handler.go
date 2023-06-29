package controller

import (
	"net/http"

	"go-web-boilerplate/config"
	"go-web-boilerplate/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bks, err := model.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "books.gohtml", bks)
}
