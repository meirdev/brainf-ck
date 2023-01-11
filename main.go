package main

import (
	"fmt"
	"os"
)

const (
	memorySize = 2048
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: bf filename")
		os.Exit(1)
	}

	program, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	Run(string(program), memorySize)
}
