// Подобным образом мы можем использовать еще ряд операторов, которые аналогичны стандартным операторам сравнения:
// eq: возвращает true, если два значения равны
// ne: возвращает true, если два значения НЕ равны
// le: возвращает true, если первое значение меньше или равно второму
// gt: возвращает true, если первое значение больше второго
// ge: возвращает true, если первое значение больше или равно второму
//
// Кроме того, есть ряд операторов, которые аналогичны логическим операторам:
// and: возвращает true, если два выражения равны true
// or: возвращает true, если хотя бы одно из двух выражений равно true
// not: возвращает true, если выражение возвращает false
//
// <div>
// 		{{ if or (gt 2 1) (lt 5 7) }}
// 			<p>Первый вариант</p>
// 		{{ else }}
// 			<p>Второй вариант</p>
//		{{ end }}
// </div>
package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Age     uint8
	Money   float64
	Hobbies []string
}

type ViewData struct {
	Title string
	Users []User
}

func homePage(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:    "Bob",
		Age:     23,
		Money:   2100,
		Hobbies: []string{"Skate", "Dance", "Football"},
	}
	tmpl, _ := template.ParseFiles("template/home.html")
	_ = tmpl.Execute(w, &user)
}

func dataPage(w http.ResponseWriter, r *http.Request) {
	data := &ViewData{
		Title: "Users List",
		Users: []User{
			{Name: "Tom", Age: 21},
			{Name: "Kate", Age: 23},
			{Name: "Alice", Age: 25},
		},
	}
	tmpl, _ := template.ParseFiles("template/data.html")
	_ = tmpl.Execute(w, data)
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/data", dataPage)

	_ = http.ListenAndServe(":8181", nil)
}

func main() {
	handleRequest()
}
