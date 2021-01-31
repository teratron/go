package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)
	logger.Print("Hello, log file!")
	fmt.Print(&buf)

	// Output
	var (
		buf2    bytes.Buffer
		logger2 = log.New(&buf2, "INFO: ", log.Lshortfile)
	)
	info := func(info string) {
		_ = logger2.Output(2, info)
	}
	info("Hello world")
	fmt.Print(&buf2)
}
