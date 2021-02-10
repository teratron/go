package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "<b>%s</b>", "Bold text")
		_, _ = fmt.Fprintf(w, `<b>%s</b>
							<h1>Header</h1>`, "Bold text")
	})

	_ = http.ListenAndServe(":8181", nil)
}
