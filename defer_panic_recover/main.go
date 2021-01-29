package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

// Функция f откладывает функцию, которая вызывает recover и печатает восстановленное значение (если оно не равно нулю).
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

// Функция g принимает int i и паникует, если i больше 3, иначе она вызывает себя с аргументом i + 1.
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
