package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
	
}

func main() {
	

	err := tpl.Execute(os.Stdout, "Sai_Adigoppula")

	if err != nil {
		log.Fatalln(err)
	}

}
