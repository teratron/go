// Метод FormValue() извлекает данные по ключу из запроса POST и PUT, а также из строки запроса.
// При этом он всегда возвращает строку.
// При этом также FormValue позволяет получить данные из строки запроса,
// то есть мы, например, можем передать значения для name и age через строку запроса:
// http://localhost:8181/postform?name=Sam&age=21
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "user.html")
	})

	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		age := r.FormValue("age")
		_, _ = fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
	})

	_ = http.ListenAndServe(":8181", nil)
}
