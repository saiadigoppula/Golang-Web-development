package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	f := req.FormValue("first")
	l := req.FormValue("last")

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}