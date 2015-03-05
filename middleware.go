package main

import (
    "log"
    "net/http"

    "github.com/codegangsta/negroni"
)

func main() {
    n := negroni.New(
        negroni.NewRecovery(),
        negroni.HandlerFunc(MyMiddleware),
        negroni.NewLogger(),
        negroni.NewStatic(http.Dir("public")),
    )

    n.Run(":3000")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("Logging from my custom middleware.")

    if r.URL.Query().Get("password") == "secret123" {
        next(rw, r)
    } else {
        http.Error(rw, "Unauthorized", 401)
    }

    log.Println("Custom middleware executed")
}

