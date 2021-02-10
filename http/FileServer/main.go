// Содержимое веб-приложения или веб-сайта нередко определяется в виде статических html-страниц.
// Для них не нужен какой-то дополнительный рендеринг на стороне сервера.
// Для прямой отправки статических файлов в пакете http определена функция FileServer,
// которая возващает объект Handler:
package main

import "net/http"

// http://localhost:8181
// http://localhost:8181/about.html
func main() {
	fs := http.FileServer(http.Dir("static"))

	_ = http.ListenAndServe(":8181", fs)
}
