package main

import (
	"math/rand"
	"testing"
	"testing/quick"
)

func TestRandomly(t *testing.T) {
	r := rand.New(rand.NewSource(0))
	config := &quick.Config{Rand: r}
	//fmt.Println(rand.NewSource(0), r, config)

	assertion := func(num uint8) bool {
		// fail test when argument is 254
		return num != 254
	}

	if err := quick.Check(assertion, config); err != nil {
		t.Error("failed checks", err)
	}
}
