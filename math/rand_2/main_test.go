package main

import (
	"fmt"
	"testing"
)

// In a _test.go file (assume this is seeded in main.go)
func init() {
	randInt = func() int {
		return 42
	}
}

func Test_main(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fmt.Println("Rand:", randInt())
	})
}
