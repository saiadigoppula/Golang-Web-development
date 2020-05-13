package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := (template.ParseGlob("*.gohtml"))
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)

	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
}
