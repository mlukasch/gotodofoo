package handlers

import (
	"net/http"
	"html/template"
	"fmt"
)

var handlerConfig struct {
	templates map[string]*template.Template
}

func init() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/user/register", RegisterHandler)
}

func RegisterTemplate(tmpls map[string]*template.Template) {
	handlerConfig.templates = tmpls
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	data := struct {
	}{}
	handlerConfig.templates["home"].Execute(w, data)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Register")
	data := struct {
	}{}
	handlerConfig.templates["register"].Execute(w, data)
}
