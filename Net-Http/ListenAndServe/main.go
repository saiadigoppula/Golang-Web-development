package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this function")
}

func main() {
	var b hotdog

	http.ListenAndServe(":8080", b)
}
