package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is the Home Page Handler

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// About is the page handler

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
