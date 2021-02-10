// Также для отправки файлов можно использовать функцию http.ServeFile().
// Она отправляет единичный файл по определенному пути.
// Например, используем ранее определенные файлы index.html и about.html:
package main

import "net/http"

func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/about.html")
}

func main() {
	http.HandleFunc("/about", about)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	_ = http.ListenAndServe(":8181", nil)
}
