package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog dog")
}

func (d hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat cat")
}

func main() {

	var d hotdog
	var c hotcat

	MUX := http.NewServeMux()
	MUX.Handle("/dog/", d)
	MUX.Handle("/cat", c)

	http.ListenAndServe(":8080", MUX)
}
