package main

import (
    "net/http"

    "gopkg.in/unrolled/render.v1"
)

func main() {
    r := render.New(render.Options{
        Layout: "layout",
    })
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
        rw.Write([]byte("Welcome! Visit sub-pages now."))
    })

    mux.HandleFunc("/data", func(rw http.ResponseWriter, req *http.Request) {
        r.Data(rw, http.StatusOK, []byte("Some binary data here."))
    })

    mux.HandleFunc("/json", func(rw http.ResponseWriter, req *http.Request) {
        r.JSON(rw, http.StatusOK, map[string]string{"hello": "json"})
    })

    mux.HandleFunc("/html", func(rw http.ResponseWriter, req *http.Request) {
        book := Book{"My Family and Other Animals", "Gerald Durrell"}
        r.HTML(rw, http.StatusOK, "example2", book)
    })

    http.ListenAndServe(":3000", mux)
}

type Book struct {
    Title string
    Author string
}