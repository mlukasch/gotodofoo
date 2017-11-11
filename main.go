package main

import (
	"html/template"
	"net/http"
	"todoappplus/handlers"
)

func main() {
	handlers.RegisterTemplate(loadTemplates())
	http.ListenAndServe(":8888", nil)
}

func loadTemplates() map[string]*template.Template {
	pages := []string{
		"home",
	}
	result := make(map[string]*template.Template)
	for _, page := range pages {
		result[page] = template.Must(
			template.ParseFiles("templates/layout.html", "templates/pages/"+page+".html"))
	}
	return result
}
