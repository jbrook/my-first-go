package main

import (
    "html/template"
    "net/http"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

type Book struct {
    Title string
    Author string
}

func main() {
    http.HandleFunc("/", ShowBooks)
    http.ListenAndServe(":3000", nil)
}

func ShowBooks(rw http.ResponseWriter, r *http.Request) {
    book := Book{"My Family and Other Animals", "Gerald Durrell"}

    // Access the template with the define name in the template
    if err := templates.ExecuteTemplate(rw, "indexPage", book); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }
}
