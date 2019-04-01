package main

import (
    "fmt"
    "log"
    "net/http"
)

const port = ":8090"

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    fmt.Println("starting server on" + port)
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(port, nil))
}
