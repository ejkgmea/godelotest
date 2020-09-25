package main

import (
    "io"
    "net/http"
)

func main () {
    http.HandleFunc("/", homePage)

    http.ListenAndServe(":5000", nil)
}

func homePage (w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1> Yo man, What's up? </h1>")
}
