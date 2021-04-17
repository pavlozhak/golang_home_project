package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go serve is runing")
}

func about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about", about_page)
	http.ListenAndServe(":4443", nil)
}

func main() {
	handleRequest()
}
