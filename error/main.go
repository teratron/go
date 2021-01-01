package main

import (
	"errors"
	"fmt"
)

type errorOne struct {
	text string
}

func (e errorOne) Error() string {
	return "Error One happened"
}

func main() {
	e1 := errorOne{"This is error"}
	e2 := fmt.Errorf("e2: %w", e1)
	e3 := fmt.Errorf("e3: %w", e2)

	fmt.Println(e1, " / ", errors.Unwrap(e1))
	fmt.Println(e2, " / ", errors.Unwrap(e2))
	fmt.Println(e3, " / ", errors.Unwrap(e3))
}
