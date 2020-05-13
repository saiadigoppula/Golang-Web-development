package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	Name string `json:"name"`
	Table string `json:"table"`
	Type string `json:"type"`
	Description string `json:"description"`
	Script string `json:"script"`
}

func main() {

	fmt.Println("Starting the application...")
	client := &http.Client{}


	todo := Todo{Name:"Info_message_from_Golang",Table:"incident",Type:"onLoad"}
	todo.Description = "demo created from golang"
	todo.Script = `function onLoad() {
		//Type appropriate comment here, and begin script below
		 g_form.addInfoMessage('this is a info message form golang');
		
	 }`

	jsonReq, err := json.Marshal(todo)
	
	req, err := http.NewRequest("POST", "https://dev58140.service-now.com/api/now/table/sys_script_client", bytes.NewBuffer(jsonReq))
	req.SetBasicAuth("admin", "Belikebro@123")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))

}
