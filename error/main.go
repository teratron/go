package main

import (
	"errors"
	"fmt"
	"os"
)

type errorOne struct {
	text string
	Err  error
}

func (e errorOne) Error() string {
	return "err1:"
}

func main() {
	e1 := errorOne{text: "This is error", Err: errors.New("error")}
	e2 := fmt.Errorf("err2: %w", e1)
	e3 := fmt.Errorf("err3: %w", e2)

	fmt.Println(e1, " / ", errors.Unwrap(e1))
	fmt.Println(e2, " / ", errors.Unwrap(e2))
	fmt.Println(e3, " / ", errors.Unwrap(e3), " / ", errors.Unwrap(errors.Unwrap(e3)))
	fmt.Println()

	// func Is(err, target error) bool
	var err1 errorOne
	var err2 error = errorOne{}
	var err3 = fmt.Errorf("e3: %w", errorOne{})

	fmt.Printf("%v(%T) - %v(%T)\n", err1, err1, err2, err2)
	if err1 == err2 {
		fmt.Println("Equality Operator: Both errors are equal")
	}
	if errors.Is(err1, err2) {
		fmt.Println("Is function: Both errors are equal")
	}

	fmt.Printf("\n%v(%T) - %v(%T)\n", err1, err1, err3, err3)
	if err1 == err3 {
		fmt.Println("Equality Operator: Both errors are equal")
	} else {
		fmt.Println("Equality Operator: Both errors are not equal")
	}
	if errors.Is(err3, err1) {
		fmt.Println("Is function: Both errors are equal")
	}
	fmt.Println()

	// func As(err error, target interface{}) bool
	err := openFile("non-existing.txt")

	if e, ok := err.(*os.PathError); ok {
		//fmt.Printf("Using Assert: Error e is of type path error. Path: %v\n", e.Path)
		fmt.Printf("Using Assert: Error e is of type path error. Error: %v\n", e)
	} else {
		fmt.Println("Using Assert: Error not of type path error")
	}
	fmt.Println()

	var pathError *os.PathError
	fmt.Println(err, " / ", pathError)
	if errors.As(err, &pathError) {
		fmt.Println(err, " / ", pathError)
		//fmt.Printf("Using As function: Error e is of type path error. Path: %v\n", pathError.Path)
		fmt.Printf("Using As function: Error e is of type path error. Error: %v\n", pathError)
	}
}

func openFile(fileName string) error {
	_, err := os.Open(fileName)
	return err
}
