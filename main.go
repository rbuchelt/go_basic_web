package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// Home is the Home Page Handler

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home Page!")
}

// About is the page handler

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is the about page and 2 + 2 is: %d", sum)
	// Usando o Fprintf a string será impressa na página do navegador, enquanto que usando o Println será impresso no console
}

// addValues add two integers and return the sum

func addValues(x, y int) int {
	return x + y
}

// main is the main application function

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
