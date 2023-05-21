package controllers

import (
	"html/template"
	"net/http"
)

type WelcomePageData struct {
	Title         string
	FirstHtmlFile string
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/welcome.html"))
	tmpl.Execute(w, WelcomePageData{
		Title:         "Go MVC",
		FirstHtmlFile: "views/welcome.html",
	})
}
