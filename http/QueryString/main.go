// Строка запроса представляет набор параметров, которые помещаются в адресе после вопросительного знака.
// При этом каждый параметр определяет название и значение. Например, в адресе:
// localhost:8181/user?name=Sam&age=21
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		_, _ = fmt.Fprintf(w, "Имя: %s Возраст: %s", name, age)
	})

	_ = http.ListenAndServe(":8181", nil)
}
