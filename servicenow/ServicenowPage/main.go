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
	Url      string
	Id       string
	Password string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/apply", apply)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func apply(w http.ResponseWriter, req *http.Request) {

	u := req.FormValue("url")
	i := req.FormValue("id")
	//p := req.FormValue("password")

	err := tpl.ExecuteTemplate(w, "servicenowlink.gohtml", person{u, i, "sindhu"})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
