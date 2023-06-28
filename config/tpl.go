package config

import "html/template"

var TPL *template.Template

func Init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
