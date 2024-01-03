package handlers

import (
	"net/http"

	"github.com/rbuchelt/basic_web/pkg/render"
)

// Home is the Home Page Handler

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// About is the page handler

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
