package main

import (
    "encoding/json"
    "net/http"
)

type Book struct {
    // fields in a struct can have tags, e.g. to tell the json package how to handle them
    Title string `json:"title"`
    WrittenBy string `json:"author"`// json key different from field name
    ISBN string `json:",omitempty"` // use field name and omit if empty
    InternalRef string `json:"-"`   // always omit
}

func main() {
    http.HandleFunc("/", ShowBooks)
    http.ListenAndServe(":3000", nil)
}

func ShowBooks(rw http.ResponseWriter, r *http.Request) {
    book := Book{"My Family and Other Animals", "Gerald Durrell", "9780140289022", "X51667"}

    js, err := json.MarshalIndent(book, "", "  ") // Marshal with pretty print
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    rw.Header().Set("Content-Type", "application/json")
    rw.Write(js)
}
