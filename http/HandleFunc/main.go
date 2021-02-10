package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "hello.html")
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Contact Page")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Index Page")
	})

	_ = http.ListenAndServe("localhost:8181", nil)
}
