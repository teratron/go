package main

import (
	"fmt"
	"net/http"
)

type msg string

func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprint(resp, "--- ", m)
}

func main() {
	_ = http.ListenAndServe("localhost:8181", msg("Hello from Web Server in Go"))
}
