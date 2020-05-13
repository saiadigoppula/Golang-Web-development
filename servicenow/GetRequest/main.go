package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)



type body []struct {
	Number string
}

type res struct {
	Body body `json:"result"`
}

func main() {
	fmt.Println("Starting the application...")
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://dev58140.service-now.com/api/now/table/incident?sysparm_limit=1", nil)
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

	var todoStruct res
	json.Unmarshal(data, &todoStruct)
	fmt.Println("")
	fmt.Println(todoStruct)

}

/*func (r *Request) SetBasicAuth(username, password string) {
	r.Header.Set("Authorization", "Basic "+basicAuth(username, password))
} */
