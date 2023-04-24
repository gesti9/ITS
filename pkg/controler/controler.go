package controler

import (
	"log"
	"net/http"
	"text/template"
)

type AuthNaP struct {
	Login string
	Pass  string
}

type Task struct {
	Name string
	Task string
}

func Auth(w http.ResponseWriter, r *http.Request) {
	user := AuthNaP{}

	tmpl, err := template.ParseFiles("../html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, user)

}
