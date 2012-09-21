package main

import (
    "net/http"
    "flag"
)

var (
    addr string
)

func init() {
    flag.StringVar(&addr, "http", ":8080", "Address of webserver")
}

func main() {

    flag.Parse()

    http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("dist/js"))))
    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("dist/css"))))
    http.Handle("/", http.FileServer(http.Dir("dist")))
    http.ListenAndServe(addr, nil)
}
