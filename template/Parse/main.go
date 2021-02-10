// В следующих уроках для файлов шаблонов мы будет использовать именовании следующего формата <название>.<роль>.tmpl,
// где <роль> будет одно из трех page, partial или layout.
// Возможность определить роль шаблона из самого названия файла поможет нам, когда придет время создания кеша шаблонов.
package main

import (
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.New("data") // определяется имя шаблона
	tmpl, _ = tmpl.Parse("<a href='/cup'>{{ . }}</a>")
	_ = tmpl.Execute(w, "Go Template")
}

func cup(w http.ResponseWriter, r *http.Request) {
	data := &ViewData{
		Title:   "World Cup",
		Message: "FIFA will never regret it",
	}
	// Так как в данном случае используются сложные данные, то их надо обернуть в функцию template.Must()
	tmpl := template.Must(template.New("data").Parse(
		`<div>
            <h1>{{ .Title }}</h1>
            <p>{{ .Message }}</p>
        </div>`))
	_ = tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cup", cup)

	_ = http.ListenAndServe(":8181", nil)
}
