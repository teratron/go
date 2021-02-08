package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", saveArticle)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err == nil {
		err = tmpl.ExecuteTemplate(w, "index", nil)
	}
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err == nil {
		err = tmpl.ExecuteTemplate(w, "create", nil)
	}
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	fullText := r.FormValue("full-text")

	_, _ = fmt.Fprintln(w, title, anons, fullText)

	/*tmpl, err := template.ParseFiles("templates/save_article.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}
	err = tmpl.ExecuteTemplate(w, "save_article", nil)
	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}*/
}

func main() {
	handleFunc()
}
