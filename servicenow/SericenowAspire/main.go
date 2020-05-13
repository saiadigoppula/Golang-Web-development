package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	Shortdescription string `json:"short_description"`
}

func main() {

	fmt.Println("Starting the application...")
	client := &http.Client{}
	todo := Todo{"from Golang"}
	jsonReq, err := json.Marshal(todo)
	//resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	req, err := http.NewRequest("POST", "https://dev58140.service-now.com/api/now/table/incident", bytes.NewBuffer(jsonReq))
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
