package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Sai adigoppula"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>`)

	NewFile, err := os.Create("index.html")

	if err != nil {
		log.Fatal("error creationg file", err)
	}

	defer NewFile.Close()

	io.Copy(NewFile, strings.NewReader(str))
	
}
