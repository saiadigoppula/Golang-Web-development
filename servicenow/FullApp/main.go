package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	Name        string `json:"name"`
	Table       string `json:"table"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Script      string `json:"script"`
}

type person struct {
	URL      string
	
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/apply", apply)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

	//fmt.Println("Starting the application...")

}

func apply(w http.ResponseWriter, request *http.Request) {

	u := request.FormValue("url")
	i := request.FormValue("id")
	p := request.FormValue("password")

	client := &http.Client{}

	todo := Todo{Name: "Info_message_from_Golang", Table: "incident", Type: "onLoad"}
	todo.Description = "demo created from golang"
	todo.Script = `function onLoad() {
		//Type appropriate comment here, and begin script below
		 g_form.addInfoMessage('this is a info message form golang');
		
	 }`

	jsonReq, err := json.Marshal(todo)
	urlsaver := `https://dev`+u+`.service-now.com/api/now/table/sys_script_client`	
	req, err := http.NewRequest("POST", urlsaver, bytes.NewBuffer(jsonReq))
	//req, err := http.NewRequest("POST", "https://dev58140.service-now.com/api/now/table/sys_script_client", bytes.NewBuffer(jsonReq))
	req.SetBasicAuth(i,p)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
	Link := `https://dev`+u+`.service-now.com/nav_to.do?uri=%2Fsys_script_client_list.do%3Fsysparm_query%3DnameSTARTSWITHInfo_message_from_Golang%26sysparm_first_row%3D1%26sysparm_view%3D`
	err  = tpl.ExecuteTemplate(w, "servicenowlink.gohtml", person{Link})
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
