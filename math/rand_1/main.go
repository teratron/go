package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(Randomly())
}

func Randomly() uint8 {
	return uint8(rand.Int())
}
