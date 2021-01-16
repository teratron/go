package main

import (
	"fmt"
	"math/rand"
)

var in = rand.Int()
var randInt = rand.Int

func main() {
	rand.Seed(1)
	fmt.Println("Rand:", randInt(), in)
}
