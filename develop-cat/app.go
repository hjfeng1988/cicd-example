package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/handlers"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "This is app cat v0.0.111: Hello, you've requested: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":80", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}

