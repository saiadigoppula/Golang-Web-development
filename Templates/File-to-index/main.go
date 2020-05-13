package main

import(
	"log"
	"os"
	"text/template"
)

func main(){
	tpl,err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	
	err = tpl.Execute(os.Stdout,nil)
	 // OR 
	/* for msking in to index.html
	NewFile, err := os.Create("index.html")

	if err != nil {
		log.Fatal("error creationg file", err)
	}

	defer NewFile.Close()

	err = tpl.Execute(NewFile,nil)
	*/
	if err != nil {
		log.Fatalln(err)
	}
}